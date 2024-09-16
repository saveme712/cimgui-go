// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "cimnodes_wrapper.h"
// #include "cimnodes_typedefs.h"
import "C"

type EmulateThreeButtonMouse struct {
	CData *C.EmulateThreeButtonMouse
}

// Handle returns C version of EmulateThreeButtonMouse and its finalizer func.
func (self *EmulateThreeButtonMouse) Handle() (result *C.EmulateThreeButtonMouse, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self EmulateThreeButtonMouse) C() (C.EmulateThreeButtonMouse, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewEmulateThreeButtonMouseFromC creates EmulateThreeButtonMouse from its C pointer.
// SRC ~= *C.EmulateThreeButtonMouse
func NewEmulateThreeButtonMouseFromC[SRC any](cvalue SRC) *EmulateThreeButtonMouse {
	return &EmulateThreeButtonMouse{CData: ConvertCTypes[*C.EmulateThreeButtonMouse](cvalue)}
}

type NodesContext struct {
	CData *C.ImNodesContext
}

// Handle returns C version of NodesContext and its finalizer func.
func (self *NodesContext) Handle() (result *C.ImNodesContext, fin func()) {
	return self.CData, func() {}
}

// NewNodesContextFromC creates NodesContext from its C pointer.
// SRC ~= *C.ImNodesContext
func NewNodesContextFromC[SRC any](cvalue SRC) *NodesContext {
	return &NodesContext{CData: ConvertCTypes[*C.ImNodesContext](cvalue)}
}

type NodesEditorContext struct {
	CData *C.ImNodesEditorContext
}

// Handle returns C version of NodesEditorContext and its finalizer func.
func (self *NodesEditorContext) Handle() (result *C.ImNodesEditorContext, fin func()) {
	return self.CData, func() {}
}

// NewNodesEditorContextFromC creates NodesEditorContext from its C pointer.
// SRC ~= *C.ImNodesEditorContext
func NewNodesEditorContextFromC[SRC any](cvalue SRC) *NodesEditorContext {
	return &NodesEditorContext{CData: ConvertCTypes[*C.ImNodesEditorContext](cvalue)}
}

type NodesIO struct {
	CData *C.ImNodesIO
}

// Handle returns C version of NodesIO and its finalizer func.
func (self *NodesIO) Handle() (result *C.ImNodesIO, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self NodesIO) C() (C.ImNodesIO, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewNodesIOFromC creates NodesIO from its C pointer.
// SRC ~= *C.ImNodesIO
func NewNodesIOFromC[SRC any](cvalue SRC) *NodesIO {
	return &NodesIO{CData: ConvertCTypes[*C.ImNodesIO](cvalue)}
}

type NodesMiniMapNodeHoveringCallbackUserData struct {
	Data uintptr
}

// Handle returns C version of NodesMiniMapNodeHoveringCallbackUserData and its finalizer func.
func (self *NodesMiniMapNodeHoveringCallbackUserData) Handle() (result *C.ImNodesMiniMapNodeHoveringCallbackUserData, fin func()) {
	r, f := self.C()
	return &r, f
}

// C is like Handle but returns plain type instead of pointer.
func (self NodesMiniMapNodeHoveringCallbackUserData) C() (C.ImNodesMiniMapNodeHoveringCallbackUserData, func()) {
	return (C.ImNodesMiniMapNodeHoveringCallbackUserData)(C.ImNodesMiniMapNodeHoveringCallbackUserData_fromUintptr(C.uintptr_t(self.Data))), func() {}
}

// NewNodesMiniMapNodeHoveringCallbackUserDataFromC creates NodesMiniMapNodeHoveringCallbackUserData from its C pointer.
// SRC ~= *C.ImNodesMiniMapNodeHoveringCallbackUserData
func NewNodesMiniMapNodeHoveringCallbackUserDataFromC[SRC any](cvalue SRC) *NodesMiniMapNodeHoveringCallbackUserData {
	return &NodesMiniMapNodeHoveringCallbackUserData{Data: (uintptr)(C.ImNodesMiniMapNodeHoveringCallbackUserData_toUintptr(*ConvertCTypes[*C.ImNodesMiniMapNodeHoveringCallbackUserData](cvalue)))}
}

type NodesStyle struct {
	CData *C.ImNodesStyle
}

// Handle returns C version of NodesStyle and its finalizer func.
func (self *NodesStyle) Handle() (result *C.ImNodesStyle, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self NodesStyle) C() (C.ImNodesStyle, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewNodesStyleFromC creates NodesStyle from its C pointer.
// SRC ~= *C.ImNodesStyle
func NewNodesStyleFromC[SRC any](cvalue SRC) *NodesStyle {
	return &NodesStyle{CData: ConvertCTypes[*C.ImNodesStyle](cvalue)}
}

type LinkDetachWithModifierClick struct {
	CData *C.LinkDetachWithModifierClick
}

// Handle returns C version of LinkDetachWithModifierClick and its finalizer func.
func (self *LinkDetachWithModifierClick) Handle() (result *C.LinkDetachWithModifierClick, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self LinkDetachWithModifierClick) C() (C.LinkDetachWithModifierClick, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewLinkDetachWithModifierClickFromC creates LinkDetachWithModifierClick from its C pointer.
// SRC ~= *C.LinkDetachWithModifierClick
func NewLinkDetachWithModifierClickFromC[SRC any](cvalue SRC) *LinkDetachWithModifierClick {
	return &LinkDetachWithModifierClick{CData: ConvertCTypes[*C.LinkDetachWithModifierClick](cvalue)}
}

type MultipleSelectModifier struct {
	CData *C.MultipleSelectModifier
}

// Handle returns C version of MultipleSelectModifier and its finalizer func.
func (self *MultipleSelectModifier) Handle() (result *C.MultipleSelectModifier, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self MultipleSelectModifier) C() (C.MultipleSelectModifier, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewMultipleSelectModifierFromC creates MultipleSelectModifier from its C pointer.
// SRC ~= *C.MultipleSelectModifier
func NewMultipleSelectModifierFromC[SRC any](cvalue SRC) *MultipleSelectModifier {
	return &MultipleSelectModifier{CData: ConvertCTypes[*C.MultipleSelectModifier](cvalue)}
}
