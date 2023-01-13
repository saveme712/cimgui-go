package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// returnTypeType represents an arbitrary type of return value of the function.
// for example Known reffers to returnTypeWrappersMap (see below)
type returnTypeType byte

const (
	// default value - will cause the function to be skipped and an error will be printed to stdout
	returnTypeUnknown returnTypeType = iota
	// return type is void (in go - the function returns nothing)
	returnTypeVoid
	// METHOD returns nothing, but it has receiver (called self)
	returnTypeStructSetter
	// Known - reffers to getReturnTypeWrapperFunc
	returnTypeKnown
	// the return value is an enum type (autogenerated by gengo_enums.go)
	returnTypeEnum
	// return type is a pointer to ImGui struct
	returnTypeStructPtr
	// returns ImGui struct
	returnTypeStruct
	// the method is a constructor
	returnTypeConstructor
)

// generateGoFuncs generates given list of functions and writes them to file
func generateGoFuncs(prefix string, validFuncs []FuncDef, enumNames []string, structNames []string) error {
	generator := &goFuncsGenerator{
		prefix:      prefix,
		structNames: make(map[string]bool),
		enumNames:   make(map[string]bool),
	}

	for _, v := range structNames {
		generator.structNames[v] = true
	}

	for _, v := range enumNames {
		generator.enumNames[v] = true
	}

	generator.writeFuncsFileHeader()

	for _, f := range validFuncs {
		// check whether the function shouldn't be skipped
		if skippedFuncs[f.FuncName] {
			continue
		}

		args, argWrappers := generator.generateFuncArgs(f)

		if len(f.ArgsT) == 0 {
			generator.shouldGenerate = true
		}

		// stop, when the function should not be generated
		if !generator.shouldGenerate {
			fmt.Printf("not generated: %s%s\n", f.FuncName, f.Args)
			continue
		} else {
			fmt.Printf("generated: %s%s\n", f.FuncName, f.Args)
		}

		if noErrors := generator.GenerateFunction(f, args, argWrappers); !noErrors {
			continue
		}
	}

	fmt.Printf("Convert progress: %d/%d\n", generator.convertedFuncCount, len(validFuncs))

	goFile, err := os.Create(fmt.Sprintf("%s_funcs.go", prefix))
	if err != nil {
		panic(err.Error())
	}

	defer goFile.Close()

	_, err = goFile.WriteString(generator.sb.String())
	if err != nil {
		return fmt.Errorf("failed to write content of GO file: %w", err)
	}

	return nil
}

// goFuncsGenerator is an internal state of GO funcs' generator
type goFuncsGenerator struct {
	prefix                 string
	structNames, enumNames map[string]bool

	sb                 strings.Builder
	convertedFuncCount int

	shouldGenerate bool
}

// writeFuncsFileHeader writes a header of the generated file
func (g *goFuncsGenerator) writeFuncsFileHeader() {
	g.sb.WriteString(goPackageHeader)

	g.sb.WriteString(fmt.Sprintf(
		`// #include "extra_types.h"
// #include "%[1]s_structs_accessor.h"
// #include "%[1]s_wrapper.h"
import "C"
import "unsafe"

`, g.prefix))
}

func (g *goFuncsGenerator) GenerateFunction(f FuncDef, args []string, argWrappers []argOutput) (noErrors bool) {
	switch {
	case f.NonUDT == 1:
		noErrors = g.generateNonUDTFunc(f, args, argWrappers)
	default:
		noErrors = g.generateFunc(f, args, argWrappers)
	}

	if noErrors {
		g.sb.WriteString("}\n\n")
		g.convertedFuncCount += 1
	}

	return
}

// generateFunc will smartly generate a function basing on its return type and arguments.
func (g *goFuncsGenerator) generateFunc(f FuncDef, args []string, argWrappers []argOutput) (noErrors bool) {
	var returnType, returnStmt, receiver string
	funcName := f.FuncName

	// determine kind of function:
	returnTypeType := returnTypeUnknown
	rf, err := getReturnTypeWrapperFunc(f.Ret)
	if err == nil {
		returnTypeType = returnTypeKnown
	}

	goEnumName := f.Ret
	if f.Ret == "void" {
		if f.StructSetter {
			returnTypeType = returnTypeStructSetter
		} else {
			returnTypeType = returnTypeVoid
		}
	} else if g.enumNames[goEnumName] {
		returnTypeType = returnTypeEnum
	} else if strings.HasSuffix(f.Ret, "*") && (g.structNames[strings.TrimSuffix(f.Ret, "*")] || g.structNames[strings.TrimSuffix(strings.TrimPrefix(f.Ret, "const "), "*")]) {
		returnTypeType = returnTypeStructPtr
	} else if f.StructGetter && g.structNames[f.Ret] {
		returnTypeType = returnTypeStruct
	} else if f.Constructor {
		returnTypeType = returnTypeConstructor
	}

	// determine function name, return type (and return statement)
	switch returnTypeType {
	case returnTypeVoid:
		// noop
	case returnTypeStructSetter:
		funcParts := strings.Split(f.FuncName, "_")
		funcName = strings.TrimPrefix(f.FuncName, funcParts[0]+"_")
		if len(funcName) == 0 || !strings.HasPrefix(funcName, "Set") || skippedStructs[funcParts[0]] {
			return false
		}

		receiver = funcParts[0]
	case returnTypeKnown:
		returnType = rf.returnType
		returnStmt = rf.returnStmt
	case returnTypeEnum:
		returnType = goEnumName
	case returnTypeStructPtr:
		// return Im struct ptr
		returnType = strings.TrimPrefix(f.Ret, "const ")
		returnType = strings.TrimSuffix(returnType, "*")
	case returnTypeStruct:
		returnType = f.Ret
	case returnTypeConstructor:
		parts := strings.Split(f.FuncName, "_")

		returnType = parts[0]

		if !g.structNames[returnType] {
			return false
		}

		suffix := ""
		if len(parts) > 2 {
			suffix = strings.Join(parts[2:], "")
		}

		funcName = "New" + returnType + suffix
	default:
		fmt.Printf("Unknown return type \"%s\" in function %s\n", f.Ret, f.FuncName)
		return false
	}

	g.sb.WriteString(g.generateFuncDeclarationStmt(receiver, funcName, args, returnType, f))
	argInvokeStmt := g.generateFuncBody(argWrappers)

	switch returnTypeType {
	case returnTypeVoid:
		g.sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.CWrapperFuncName, argInvokeStmt))
	case returnTypeStructSetter:
		g.sb.WriteString(fmt.Sprintf("C.%s(self.handle(), %s)\n", f.CWrapperFuncName, argInvokeStmt))
	case returnTypeKnown:
		g.sb.WriteString(fmt.Sprintf(returnStmt, fmt.Sprintf("C.%s(%s)", f.CWrapperFuncName, argInvokeStmt)))
	case returnTypeEnum:
		g.sb.WriteString(fmt.Sprintf("return %s(C.%s(%s))", renameGoIdentifier(returnType), f.CWrapperFuncName, argInvokeStmt))
	case returnTypeStructPtr:
		g.sb.WriteString(fmt.Sprintf("return (%s)(unsafe.Pointer(C.%s(%s)))", renameGoIdentifier(returnType), f.CWrapperFuncName, argInvokeStmt))
	case returnTypeStruct:
		g.sb.WriteString(fmt.Sprintf("return new%sFromC(C.%s(%s))", renameGoIdentifier(f.Ret), f.CWrapperFuncName, argInvokeStmt))
	case returnTypeConstructor:
		g.sb.WriteString(fmt.Sprintf("return (%s)(unsafe.Pointer(C.%s(%s)))", renameGoIdentifier(returnType), f.CWrapperFuncName, argInvokeStmt))
	}

	return true
}

// generateNonUDTFunc generates a function tagged with NonUDT.
// it will consider a first argument of this function as a return type.
/*
	template:
	func FuncName(arg2 type2) typeOfArg1 {
		pOut := &typeOfArg1{}
		pOutArg, pOutFin := pOut.wrapped()
		defer pOutFin()
		C.FuncName(pOutArg, arg2)
		return *pOut
	}
*/
func (g *goFuncsGenerator) generateNonUDTFunc(f FuncDef, args []string, argWrappers []argOutput) (noErrors bool) {
	// find out the return type
	outArg := f.ArgsT[0]
	outArgT := strings.TrimSuffix(outArg.Type, "*")
	returnWrapper, err := getReturnTypeWrapperFunc(outArgT)
	if err != nil {
		fmt.Printf("Unknown return type \"%s\" in function %s\n", f.Ret, f.FuncName)
		return false
	}

	returnType := returnWrapper.returnType

	g.sb.WriteString(g.generateFuncDeclarationStmt("", f.FuncName, args[1:], returnType, f))

	// temporary out arg definition
	g.sb.WriteString(fmt.Sprintf("%s := &%s{}\n", outArg.Name, returnType))

	argInvokeStmt := g.generateFuncBody(argWrappers)

	// C function call
	g.sb.WriteString(fmt.Sprintf("C.%s(%s)\n", f.CWrapperFuncName, argInvokeStmt))

	// return statement
	g.sb.WriteString(fmt.Sprintf("return *%s", outArg.Name))

	return true
}

// this method is responsible for createing a function declaration statement.
// it takes function name, list of arguments and return type and returns go statement.
// e.g.: func (self *ImGuiType) FuncName(arg1 type1, arg2 type2) returnType {
func (g *goFuncsGenerator) generateFuncDeclarationStmt(receiver string, funcName string, args []string, returnType string, f FuncDef) (functionDeclaration string) {
	funcParts := strings.Split(funcName, "_")
	typeName := funcParts[0]

	if len(receiver) > 0 {
		receiver = fmt.Sprintf("(self %s)", renameGoIdentifier(receiver))
	}

	// Generate default param value hint
	var commentSb strings.Builder
	if len(f.Defaults) > 0 {
		commentSb.WriteString("// %s parameter default value hint:\n")

		// sort lexicographically for determenistic generation
		type defaultParam struct {
			name  string
			value string
		}
		defaults := make([]defaultParam, 0, len(f.Defaults))
		for n, v := range f.Defaults {
			defaults = append(defaults, defaultParam{name: n, value: v})
		}
		sort.Slice(defaults, func(i, j int) bool {
			return defaults[i].name < defaults[j].name
		})

		for _, p := range defaults {
			commentSb.WriteString(fmt.Sprintf("// %s: %s\n", p.name, p.value))
		}
	}

	if strings.Contains(funcName, "_") &&
		len(funcParts) > 1 &&
		len(args) > 0 && strings.Contains(args[0], "self ") {

		newFuncName := strings.TrimPrefix(funcName, typeName+"_")
		newArgs := args
		if len(newArgs) > 0 {
			newArgs = args[1:]
		}

		typeName = strings.TrimPrefix(args[0], "self ")
		return fmt.Sprintf("%sfunc %s (self %s) %s(%s) %s {\n",
			strings.Replace(commentSb.String(), "%s", renameGoIdentifier(newFuncName), 1),
			renameGoIdentifier(receiver),
			renameGoIdentifier(typeName),
			renameGoIdentifier(newFuncName),
			strings.Join(newArgs, ","),
			renameGoIdentifier(returnType))
	}

	return fmt.Sprintf("%sfunc %s %s(%s) %s {\n",
		strings.Replace(commentSb.String(), "%s", renameGoIdentifier(funcName), 1),
		renameGoIdentifier(receiver),
		renameGoIdentifier(funcName),
		strings.Join(args, ","),
		renameGoIdentifier(returnType))
}

func (g *goFuncsGenerator) generateFuncArgs(f FuncDef) (args []string, argWrappers []argOutput) {
	for i, a := range f.ArgsT {
		g.shouldGenerate = false

		if a.Name == "type" {
			a.Name = "typeArg"
		}

		if i == 0 && f.StructSetter {
			g.shouldGenerate = true
		}

		if f.StructGetter && g.structNames[a.Type] {
			args = append(args, fmt.Sprintf("%s %s", a.Name, renameGoIdentifier(a.Type)))
			argWrappers = append(argWrappers, argOutput{
				VarName: fmt.Sprintf("%s.handle()", a.Name),
			})

			g.shouldGenerate = true

			continue
		}

		if v, err := argWrapper(a.Type); err == nil {
			argType, argDef, varName := v(a)
			if goEnumName := argType; g.isEnum(goEnumName) {
				argType = goEnumName
			}

			argWrappers = append(argWrappers, argOutput{
				ArgType: argType,
				ArgDef:  argDef,
				VarName: varName,
			})

			args = append(args, fmt.Sprintf("%s %s", a.Name, renameGoIdentifier(argType)))

			g.shouldGenerate = true
			continue
		}

		if goEnumName := a.Type; g.isEnum(goEnumName) {
			args = append(args, fmt.Sprintf("%s %s", a.Name, renameGoIdentifier(goEnumName)))
			argWrappers = append(argWrappers, argOutput{
				VarName: fmt.Sprintf("C.%s(%s)", a.Type, a.Name),
			})

			g.shouldGenerate = true
			continue
		}

		if strings.HasSuffix(a.Type, "*") {
			pureType := strings.TrimPrefix(a.Type, "const ")
			pureType = strings.TrimSuffix(pureType, "*")

			if g.structNames[pureType] {
				args = append(args, fmt.Sprintf("%s %s", a.Name, renameGoIdentifier(pureType)))
				argWrappers = append(argWrappers, argOutput{
					VarName: fmt.Sprintf("%s.handle()", a.Name),
				})

				g.shouldGenerate = true
				continue
			}
		}

		if !g.shouldGenerate {
			fmt.Printf("Unknown argument type \"%s\" in function %s\n", a.Type, f.FuncName)
			break
		}
	}

	return args, argWrappers
}

// Generate function body
// and returns function call arguments
// e.g.:
// it will write the following into the buffer:
func (g *goFuncsGenerator) generateFuncBody(argWrappers []argOutput) string {
	var invokeStmt []string
	for _, aw := range argWrappers {
		invokeStmt = append(invokeStmt, aw.VarName)
		if len(aw.ArgDef) > 0 {
			g.sb.WriteString(fmt.Sprintf("%s\n\n", aw.ArgDef))
		}
	}

	return strings.Join(invokeStmt, ",")
}

// isEnum returns true when given string is a valid enum type.
func (g *goFuncsGenerator) isEnum(argType string) bool {
	for en := range g.enumNames {
		if argType == en {
			return true
		}
	}

	return false
}
