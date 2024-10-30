// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

#pragma once

#include "../cwrappers/cimCTE.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void wrap_TextEditor_ImGuiDebugPanel(TextEditor* self);
extern void wrap_TextEditor_Redo(TextEditor* self);
extern bool wrap_TextEditor_Render(TextEditor* self,const char* aTitle);
extern void wrap_TextEditor_SelectAllOccurrencesOf(TextEditor* self,const char* aText,int aTextSize);
extern void wrap_TextEditor_SelectNextOccurrenceOf(TextEditor* self,const char* aText,int aTextSize);
extern void wrap_TextEditor_Undo(TextEditor* self);

#ifdef __cplusplus
}
#endif
