package main

import "C"
import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/kpango/glg"
)

type typedefsGenerator struct {
	GoSb        *strings.Builder
	CGoHeaderSb *strings.Builder
	HSb         *strings.Builder
	CppSb       *strings.Builder
	ctx         *Context
}

// this function will proceed the following typedefs:
// - all structs thatare not present in struct_and_enums.json (they are supposed to be epaque)
// - everything that satisfies IsCallbackTypedef
func proceedTypedefs(
	typedefs *Typedefs,
	structs []StructDef,
	ctx *Context,
) (validTypeNames []CIdentifier, err error) {
	// quick counter for coverage control
	generatedTypedefs := 0
	maxTypedefs := len(typedefs.data)

	generator := &typedefsGenerator{
		// we need FILES
		GoSb:        &strings.Builder{},
		CGoHeaderSb: &strings.Builder{},
		HSb:         &strings.Builder{},
		CppSb:       &strings.Builder{},
		ctx:         ctx,
	}

	generator.writeHeaders()

	// because go ranges through maps as if it was drunken, we need to sort keys.
	keys := make([]CIdentifier, 0, len(typedefs.data))
	for k := range typedefs.data {
		keys = append(keys, k)
	}

	// sort keys
	SortStrings(keys)

typedefsGeneration:
	for _, k := range keys {
		typedef := typedefs.data[k]
		if shouldSkip, ok := skippedTypedefs[k]; ok && shouldSkip {
			if ctx.flags.showNotGenerated {
				glg.Infof("Arbitrarly skipping typedef %s", k)
			}

			maxTypedefs--
			continue
		}

		if _, exists := ctx.refTypedefs[k]; exists {
			if ctx.flags.showNotGenerated {
				glg.Infof("Duplicate of %s in reference typedefs. Skipping.", k)
			}

			maxTypedefs--
			continue
		}

		if shouldSkipStruct(k) {
			if ctx.flags.showNotGenerated {
				glg.Infof("Arbitrarly skipping struct %s", k)
			}

			maxTypedefs--
			continue
		}

		if IsEnumName(k, ctx.enumNames) /*|| IsStructName(k, structs)*/ {
			if ctx.flags.showGenerated {
				glg.Infof("typedef %s has extended deffinition in structs_and_enums.json. Will generate later", k)
			}

			maxTypedefs--
			continue
		}

		if IsTemplateTypedef(typedef) {
			if ctx.flags.showNotGenerated {
				glg.Warnf("typedef %s is a template. not implemented yet", k)
			}

			continue
		}

		isPtr := HasSuffix(typedef, "*")

		known := generator.parseArgDef(CIdentifier(typedef), ctx)

		switch {
		case typedefs.data[k] == "void*":
			if ctx.flags.showGenerated {
				glg.Successf("typedef %s is an alias to void*.", k)
			}

			generator.writeVoidPtrTypedef(k, known)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case !isPtr &&
			known.ptrReturnTypeErr == nil &&
			known.argTypeErr == nil &&
			known.ptrArgTypeErr == nil:
			if ctx.flags.showGenerated {
				glg.Successf("typedef %s is an alias typedef.", k)
			}

			fmt.Fprintf(generator.GoSb, `
type %[1]s %[2]s

// Handle returns C version of %[1]s and its finalizer func.
func (selfSrc *%[1]s) Handle() (result *C.%[6]s, fin func()) {
	self := (*%[2]s)(selfSrc)
    %[3]s
    return (*C.%[6]s)(%[4]s), func() { %[5]s }
}

// C is like Handle but returns plain type instead of pointer.
func (self %[1]s) C() (C.%[6]s, func()) {
    %[7]s
    return (C.%[6]s)(%[8]s), func() { %[9]s }
}

// New%[1]sFromC creates %[1]s from its C pointer.
// SRC ~= *C.%[6]s
func New%[1]sFromC[SRC any](cvalue SRC) *%[1]s {
	return (*%[1]s)(%[10]s)
}
`,
				k.renameGoIdentifier(),
				known.argType.ArgType,

				known.ptrArgType.ArgDef,
				known.ptrArgType.VarName,
				known.ptrArgType.Finalizer,

				k,

				known.argType.ArgDef,
				known.argType.VarName,
				known.argType.Finalizer,

				fmt.Sprintf(known.ptrReturnType.returnStmt, fmt.Sprintf("internal.ReinterpretCast[*C.%s](cvalue)", k)),
			)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case isPtr &&
			known.returnTypeErr == nil &&
			known.argTypeErr == nil:
			// if it's a pointer type, I think we can proceed as above, but without Handle() method...
			// (handle proceeds pointer values and we don't want double pointers, really)
			fmt.Fprintf(generator.GoSb, `
type %[1]s  struct {
	Data %[2]s
}

// Handle returns C version of %[1]s and its finalizer func.
func (self *%[1]s) Handle() (*C.%[6]s, func()) {
	result, fn := self.C()
	return &result, fn
}

// C is like Handle but returns plain type instead of pointer.
func (selfStruct *%[1]s) C() (result C.%[6]s, fin func()) {
	self := selfStruct.Data
    %[3]s
    return (C.%[6]s)(%[4]s), func() { %[5]s }
}

// New%[1]sFromC creates %[1]s from its C pointer.
// SRC ~= *C.%[6]s
func New%[1]sFromC[SRC any](cvalue SRC) *%[1]s {
	v := (%[8]s)(*internal.ReinterpretCast[*C.%[6]s](cvalue))
	return &%[1]s{Data: %[7]s}
}
`,
				k.renameGoIdentifier(),
				known.argType.ArgType,

				known.argType.ArgDef,
				known.argType.VarName,
				known.argType.Finalizer,

				k,

				fmt.Sprintf(known.returnType.returnStmt, "v"),
				known.argType.CType,
			)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		case IsCallbackTypedef(typedefs.data[k]):
			// see https://github.com/AllenDang/cimgui-go/issues/224#issuecomment-2452156237
			// 1: preprocessing - parse typedefs.data[k] to get return type and arguments
			typedefName := CIdentifier(k)
			// now let me use a bit of regex.
			// We have 2 possibilities:
			// - returnType(*<funcName>)(args1 arg1Name, args2 arg2Name, args3 arg3Name);
			// - returnType <funcName>(args1 arg1Name, args2 arg2Name, args3 arg3Name);
			// NOTE: the second is uesed mainly in immarkdown
			// NOTE: in the 1st, spaces does not matter so we'll trim them
			expr1, err := regexp.Compile("([a-zA-Z0-9_]+\\*?)\\(\\*.*\\)\\((.*)\\);")
			if err != nil {
				panic(fmt.Sprintf("Cannot compile regex expr1!: %v", err))
			}

			expr2, err := regexp.Compile("([a-zA-Z0-9_]+)\\s+([a-zA-Z0-9_]+)\\((.*)\\);")
			if err != nil {
				panic(fmt.Sprintf("Cannot compile regex expr2!: %v", err))
			}

			// we need the following from them:
			var returnTypeC CIdentifier
			argsC := make([]ArgDef, 0)

			if ok := expr1.MatchString(typedefs.data[k]); ok {
				glg.Debugf("callback typedef \"%s\" is in form 1", k)
				// now split by "("
				// it should be something like this:
				// ["returnType", "*<optional func name)", "args1 arg1Name, args2 arg2Name, args3 arg3Name);"]
				parts := strings.Split(typedefs.data[k], "(")
				if len(parts) != 3 {
					panic("Cannot split by (, check implementation in cmd/codegen!")
				}

				returnTypeC = TrimSuffix(CIdentifier(parts[0]), " ")
				argsStr := parts[2]
				argsStr = TrimSuffix(argsStr, ");")
				argsStr = ReplaceAll(argsStr, ", ", ",")
				argsStr = ReplaceAll(argsStr, "&", "")
				argsListStr := Split(argsStr, ",")
				for a, argStr := range argsListStr {
					// get name
					argParts := Split(argStr, " ")
					var name, typeName string
					switch len(argParts) {
					case 1:
						name = fmt.Sprintf("arg%d", a)
						typeName = strings.Join(argParts, " ")
					case 2: // like "int arg1" or "const int"
						if argParts[0] == "const" {
							name = fmt.Sprintf("arg%d", a)
							typeName = strings.Join(argParts, " ")
							break
						}

						fallthrough
					case 3: // something like "int" or "const int arg1"
						name = argParts[len(argParts)-1]
						typeName = strings.Join(argParts[:len(argParts)-1], " ")
					}

					argsC = append(argsC, ArgDef{
						Name: CIdentifier(name),
						Type: CIdentifier(typeName),
					})
				}
			} else if ok := expr2.MatchString(typedefs.data[k]); ok {
				glg.Warnf("Callback option 2 for %s not implemented yet", k)
				continue
			} else {
				if ctx.flags.showNotGenerated {
					glg.Failf("cannot parse callback typedef %s: \"%s\".", k, typedefs.data[k])
					continue
				}
			}

			// 2: Find wrappers:
			// We need to figure out how to wrap returnType and args.
			// In fact, we need to swap meaning of them, because we want to convert C argument type to Go argument type
			// so we are supposed to use returnWrapper for that.
			glg.Debugf("From %s got \"%s\" and \"%v\"", k, returnTypeC, argsC)

			// 2.1: get return wrapper
			var returnType ArgumentWrapperData
			if returnTypeC == "void" {
				returnType = ArgumentWrapperData{}
			} else {
				_, returnType, err = getArgWrapper(
					&ArgDef{
						Name: "result",
						Type: returnTypeC,
					},
					false, false,
					ctx,
				)
				if err != nil {
					if ctx.flags.showNotGenerated {
						glg.Failf("cannot get return wrapper for %s - \"%s\": %v", k, returnTypeC, err)
					}
					continue typedefsGeneration
				}
			}

			// 2.1: get arg wrappers
			args := make([]returnWrapper, len(argsC))
			for i, arg := range argsC {
				rw, err := getReturnWrapper(arg.Type, ctx)
				if err != nil {
					if ctx.flags.showNotGenerated {
						glg.Failf("cannot get arg wrapper for \"%s\" - \"%s\": %v", k, arg.Type, err)
					}

					continue typedefsGeneration
				}

				// fill rw return stmt
				rw.returnStmt = fmt.Sprintf(rw.returnStmt, arg.Name)

				args[i] = rw
			}

			// 3: Prepare call statement
			cCallStmt := ""
			goCallStmt := ""
			valuePassStmt := ""
			for i, arg := range args {
				cCallStmt += fmt.Sprintf("%s %s, ", argsC[i].Name, arg.CType)
				goCallStmt += fmt.Sprintf("%s %s, ", argsC[i].Name, arg.returnType)
				valuePassStmt += fmt.Sprintf("%s, ", argsC[i].Name)
			}

			cCallStmt = TrimSuffix(cCallStmt, ", ")
			goCallStmt = TrimSuffix(goCallStmt, ", ")
			valuePassStmt = TrimSuffix(valuePassStmt, ", ")

			// 4: Write code
			fmt.Fprintf(generator.GoSb, `
type %[1]s func(%[4]s) %[2]s
type c%[1]s func(%[5]s) %[3]s

func New%[1]sFromC(cvalue *C.%[6]s) *%[1]s {
	result := pool%[1]s.Find(*cvalue)
	return &result
}

func (c %[1]s) C() (C.%[6]s, func()) {
	return pool%[1]s.Allocate(c), func() { }
}
`,
				typedefName.renameGoIdentifier(),
				returnType.ArgType,
				returnType.CType,
				goCallStmt,
				cCallStmt,
				k,
			)

			cCallStmt2 := cCallStmt
			if cCallStmt2 != "" {
				cCallStmt2 = ", " + cCallStmt2
			}

			if returnType.ArgType == "" {
				fmt.Fprintf(generator.GoSb, `
func wrap%[1]s(cb %[1]s %[3]s) %[2]s {
	cb(%[4]s)
}
`,
					typedefName.renameGoIdentifier(),
					returnType.CType,
					cCallStmt2,
					func() string {
						result := ""
						for _, a := range args {
							result += a.returnStmt + ", "
						}
						result = TrimSuffix(result, ", ")
						return result
					}(),
				)
			} else {
				fmt.Fprintf(generator.GoSb, `
func wrap%[1]s(cb %[1]s %[5]s) %[2]s {
	result := cb(%[6]s)
	%[3]s
	return %[4]s
}
`,
					typedefName.renameGoIdentifier(),
					returnType.CType,
					returnType.ArgDef,
					returnType.VarName,
					cCallStmt2,
					func() string {
						result := ""
						for _, a := range args {
							result += a.returnStmt + ", "
						}
						result = TrimSuffix(result, ", ")
						return result
					}(),
				)
			}

			size := TypedefsPoolSize
			if h, ok := customPoolSize[k]; ok {
				size = h
			}

			poolNames := make([]string, size)
			// now write N functions
			for i := 0; i < size; i++ {
				fmt.Fprintf(generator.GoSb,
					`//export callback%[1]s%[2]d
func callback%[1]s%[2]d(%[5]s) %[3]s { %[4]s wrap%[1]s(pool%[1]s.Get(%[2]d), %[6]s) }
					`,
					typedefName.renameGoIdentifier(),
					i,
					returnType.CType,
					func() string {
						if returnType.ArgType != "" {
							return "return"
						}

						return ""
					}(),
					cCallStmt,
					valuePassStmt,
				)

				fmt.Fprintf(generator.CGoHeaderSb,
					`// extern %[1]s callback%[2]s%[3]d(%[4]s);
`,
					returnTypeC,
					typedefName.renameGoIdentifier(),
					i,
					func() string {
						result := ""
						for _, a := range argsC {
							result += TrimPrefix(string(a.Type), "const ") + ", "
						}

						return TrimSuffix(result, ", ")
						return result
					}(),
				)

				poolNames[i] = fmt.Sprintf("C.%[3]s(C.callback%[1]s%[2]d)", typedefName.renameGoIdentifier(), i, k)
			}

			fmt.Fprintf(generator.GoSb, `
var pool%[1]s *internal.Pool[%[1]s, C.%[3]s]
func init() {
	pool%[1]s = internal.NewPool[%[1]s, C.%[3]s](
%[2]s
)
}

func Clear%[1]sPool() {
	pool%[1]s.Clear()
}
`,
				typedefName.renameGoIdentifier(),
				strings.Join(poolNames, ",\n")+",\n",
				k,
			)

			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)

			glg.Successf("typedef %s is a callback. Implemented.", k)
		case HasPrefix(typedefs.data[k], "struct"):
			isOpaque := !IsStructName(k, ctx)
			if ctx.flags.showGenerated {
				glg.Successf("typedef %s is a struct (is opaque? %v).", k, isOpaque)
			}

			generator.writeStruct(k, isOpaque)
			generatedTypedefs++
			validTypeNames = append(validTypeNames, k)
		default:
			if ctx.flags.showNotGenerated {
				glg.Failf("unknown situation happened for type %s; not implemented. Probably unknown Arg (err: %v), Ret (err; %v) PtrArg (err: %v) or PtrRet (err: %v) type wrappers for isPointer: %v for %s. Check out source code for more details",
					k, known.argTypeErr, known.returnTypeErr, known.ptrArgTypeErr, known.ptrReturnTypeErr, isPtr, typedefs.data[k])
			}
		}
	}

	fmt.Fprint(generator.HSb,
		`
#ifdef __cplusplus
}
#endif`)

	typedefsGoResultSb := &strings.Builder{}
	typedefsGoResultSb.WriteString(getGoPackageHeader(ctx))
	fmt.Fprintf(typedefsGoResultSb,
		`// #include <stdlib.h>
// #include <memory.h>
// #include "../imgui/extra_types.h"
// #include "%[1]s_wrapper.h"
// #include "%[1]s_typedefs.h"
`, ctx.prefix)

	typedefsGoResultSb.WriteString(generator.CGoHeaderSb.String())

	fmt.Fprintf(typedefsGoResultSb,
		`import "C"
import "unsafe"
`)

	typedefsGoResultSb.WriteString(generator.GoSb.String())

	if err := os.WriteFile(fmt.Sprintf("%s_typedefs.go", ctx.prefix), []byte(typedefsGoResultSb.String()), 0644); err != nil {
		return nil, fmt.Errorf("cannot write %s_typedefs.go: %w", ctx.prefix, err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s_typedefs.cpp", ctx.prefix), []byte(generator.CppSb.String()), 0644); err != nil {
		return nil, fmt.Errorf("cannot write %s_typedefs.cpp: %w", ctx.prefix, err)
	}

	if err := os.WriteFile(fmt.Sprintf("%s_typedefs.h", ctx.prefix), []byte(generator.HSb.String()), 0644); err != nil {
		return nil, fmt.Errorf("cannot write %s_typedefs.h: %w", ctx.prefix, err)
	}

	glg.Infof("Typedefs generation complete. Generated %d/%d (%.2f%%) typedefs.", generatedTypedefs, maxTypedefs, float32(generatedTypedefs*100)/float32(maxTypedefs))
	return validTypeNames, nil
}

// Let's say our pureType is of form "short"
// the following code needs to handle two things:
// - int16 -> short (to know go type AND know how to proceed in C() func)
// - *int16 -> short* (for Handle())
// - short* -> *int16 (for NewXXXFromC)
type typedefTypeContext struct {
	argType,
	ptrArgType ArgumentWrapperData
	returnType,
	ptrReturnType returnWrapper

	argTypeErr, ptrArgTypeErr,
	returnTypeErr, ptrReturnTypeErr error
}

func (g *typedefsGenerator) parseArgDef(typedef CIdentifier, ctx *Context) *typedefTypeContext {
	result := &typedefTypeContext{}
	result.returnType, result.returnTypeErr = getReturnWrapper(
		CIdentifier(typedef),
		ctx,
	)

	result.ptrReturnType, result.ptrReturnTypeErr = getReturnWrapper(
		CIdentifier(typedef)+"*",
		ctx,
	)

	_, result.argType, result.argTypeErr = getArgWrapper(
		&ArgDef{
			Name: "self",
			Type: CIdentifier(typedef),
		},
		false, false,
		ctx,
	)

	_, result.ptrArgType, result.ptrArgTypeErr = getArgWrapper(
		&ArgDef{
			Name: "self",
			Type: CIdentifier(typedef) + "*",
		},
		false, false,
		ctx,
	)

	return result
}

func (g *typedefsGenerator) writeHeaders() {
	g.HSb.WriteString(cppFileHeader)
	fmt.Fprintf(g.HSb,
		`
#pragma once

#include "%s"

#ifdef __cplusplus
extern "C" {
#endif
`, g.ctx.flags.include)
	g.CppSb.WriteString(cppFileHeader)
	fmt.Fprintf(g.CppSb,
		`
#include "%[1]s_typedefs.h"
#include "%[2]s"
`, g.ctx.prefix, g.ctx.flags.include)
}

// k is plain C name of the typedef (key in typedefs_dict.json)
// known is parsed value of k
func (g *typedefsGenerator) writeVoidPtrTypedef(k CIdentifier, known *typedefTypeContext) {
	fmt.Fprintf(g.CppSb,
		`
uintptr_t %[1]s_toUintptr(%[1]s ptr) {
	return (uintptr_t)ptr;
}

%[1]s %[1]s_fromUintptr(uintptr_t ptr) {
	return (%[1]s)ptr;
}
`, k)
	fmt.Fprintf(g.HSb, `extern uintptr_t %[1]s_toUintptr(%[1]s ptr);
extern %[1]s %[1]s_fromUintptr(uintptr_t ptr);`, k)

	// NOTE: in case of problems e.g. with Textures, here might be potential issue:
	// Handle() is incomplete - it doesn't have right finalizer (for now I think this will not affect code)
	fmt.Fprintf(g.GoSb, `
type %[1]s struct {
	Data uintptr
}

// Handle returns C version of %[1]s and its finalizer func.
func (self *%[1]s) Handle() (result *C.%[6]s, fin func()) {
	r, f := self.C()
    return &r, f
}

// C is like Handle but returns plain type instead of pointer.
func (self %[1]s) C() (C.%[6]s, func()) {
    return (C.%[6]s)(C.%[6]s_fromUintptr(C.uintptr_t(self.Data))), func() { }
}

// New%[1]sFromC creates %[1]s from its C pointer.
// SRC ~= *C.%[6]s
func New%[1]sFromC[SRC any](cvalue SRC) *%[1]s {
	return &%[1]s{Data: (uintptr)(C.%[6]s_toUintptr(*internal.ReinterpretCast[*C.%[6]s](cvalue)))}
}
`,
		k.renameGoIdentifier(),
		known.argType.ArgType,

		known.ptrArgType.ArgDef,
		known.ptrArgType.VarName,
		known.ptrArgType.Finalizer,

		k,

		known.argType.ArgDef,
		known.argType.VarName,
		known.argType.Finalizer,

		fmt.Sprintf(known.ptrReturnType.returnStmt, "cvalue"),
	)
}

func (g *typedefsGenerator) writeStruct(
	name CIdentifier,
	isOpaque bool, // if this is true, we assume that we cannot use the exact value of a struct (only pointer)
) {
	// this will be put only for structs that are NOT opaque (w can know the exact definition)
	var toPlainValue string
	if !isOpaque {
		toPlainValue = fmt.Sprintf(`
// C is like Handle but returns plain type instead of pointer.
func (self %[1]s) C() (C.%[2]s, func()) {
	result, fn := self.Handle()
	return *result, fn
}
`, name.renameGoIdentifier(), name)
	}

	// we need to make it a struct, because we need to hide C type (otherwise it will duplicate methods)
	fmt.Fprintf(g.GoSb, `
type %[1]s struct {
	CData *C.%[2]s
}

// Handle returns C version of %[1]s and its finalizer func.
func (self *%[1]s) Handle() (result *C.%[2]s, fin func()) {
	return self.CData, func() {}
}

%[3]s

// New%[1]sFromC creates %[1]s from its C pointer.
// SRC ~= *C.%[2]s
func New%[1]sFromC[SRC any](cvalue SRC) *%[1]s {
	return &%[1]s{CData: internal.ReinterpretCast[*C.%[2]s](cvalue)}
}
`, name.renameGoIdentifier(), name, toPlainValue)
}
