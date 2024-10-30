// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package ImGuiColorTextEdit

// #include <stdlib.h>
// #include <memory.h>
// #include "../imgui/extra_types.h"
// #include "cimcte_wrapper.h"
// #include "cimcte_typedefs.h"
import "C"
import "github.com/AllenDang/cimgui-go/internal"

type Breakpoint struct {
	CData *C.Breakpoint
}

// Handle returns C version of Breakpoint and its finalizer func.
func (self *Breakpoint) Handle() (result *C.Breakpoint, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self Breakpoint) C() (C.Breakpoint, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewBreakpointFromC creates Breakpoint from its C pointer.
// SRC ~= *C.Breakpoint
func NewBreakpointFromC[SRC any](cvalue SRC) *Breakpoint {
	return &Breakpoint{CData: internal.ReinterpretCast[*C.Breakpoint](cvalue)}
}

type Coordinates struct {
	CData *C.Coordinates
}

// Handle returns C version of Coordinates and its finalizer func.
func (self *Coordinates) Handle() (result *C.Coordinates, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self Coordinates) C() (C.Coordinates, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewCoordinatesFromC creates Coordinates from its C pointer.
// SRC ~= *C.Coordinates
func NewCoordinatesFromC[SRC any](cvalue SRC) *Coordinates {
	return &Coordinates{CData: internal.ReinterpretCast[*C.Coordinates](cvalue)}
}

type Glyph struct {
	CData *C.Glyph
}

// Handle returns C version of Glyph and its finalizer func.
func (self *Glyph) Handle() (result *C.Glyph, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self Glyph) C() (C.Glyph, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewGlyphFromC creates Glyph from its C pointer.
// SRC ~= *C.Glyph
func NewGlyphFromC[SRC any](cvalue SRC) *Glyph {
	return &Glyph{CData: internal.ReinterpretCast[*C.Glyph](cvalue)}
}

type Identifier struct {
	CData *C.Identifier
}

// Handle returns C version of Identifier and its finalizer func.
func (self *Identifier) Handle() (result *C.Identifier, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self Identifier) C() (C.Identifier, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewIdentifierFromC creates Identifier from its C pointer.
// SRC ~= *C.Identifier
func NewIdentifierFromC[SRC any](cvalue SRC) *Identifier {
	return &Identifier{CData: internal.ReinterpretCast[*C.Identifier](cvalue)}
}

type LanguageDefinition struct {
	CData *C.LanguageDefinition
}

// Handle returns C version of LanguageDefinition and its finalizer func.
func (self *LanguageDefinition) Handle() (result *C.LanguageDefinition, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self LanguageDefinition) C() (C.LanguageDefinition, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewLanguageDefinitionFromC creates LanguageDefinition from its C pointer.
// SRC ~= *C.LanguageDefinition
func NewLanguageDefinitionFromC[SRC any](cvalue SRC) *LanguageDefinition {
	return &LanguageDefinition{CData: internal.ReinterpretCast[*C.LanguageDefinition](cvalue)}
}

type TextEditor struct {
	CData *C.TextEditor
}

// Handle returns C version of TextEditor and its finalizer func.
func (self *TextEditor) Handle() (result *C.TextEditor, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self TextEditor) C() (C.TextEditor, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewTextEditorFromC creates TextEditor from its C pointer.
// SRC ~= *C.TextEditor
func NewTextEditorFromC[SRC any](cvalue SRC) *TextEditor {
	return &TextEditor{CData: internal.ReinterpretCast[*C.TextEditor](cvalue)}
}
