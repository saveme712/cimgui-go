// Code generated by cmd/codegen from https://github.com/saveme712/cimgui-go.
// DO NOT EDIT.

#pragma once

#include "../cwrappers/cimnodes.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void wrap_imnodes_BeginInputAttribute(int id);
extern void wrap_imnodes_BeginOutputAttribute(int id);
extern void wrap_imnodes_DestroyContext();
extern bool wrap_imnodes_IsAnyAttributeActive();
extern bool wrap_imnodes_IsLinkCreated_BoolPtr(int* started_at_attribute_id,int* ended_at_attribute_id);
extern bool wrap_imnodes_IsLinkCreated_IntPtr(int* started_at_node_id,int* started_at_attribute_id,int* ended_at_node_id,int* ended_at_attribute_id);
extern bool wrap_imnodes_IsLinkDropped();
extern void wrap_imnodes_MiniMap();
extern void wrap_imnodes_PopStyleVar();
extern const char* wrap_imnodes_SaveCurrentEditorStateToIniString();
extern const char* wrap_imnodes_SaveEditorStateToIniString(const ImNodesEditorContext* editor);
extern void wrap_imnodes_StyleColorsClassic();
extern void wrap_imnodes_StyleColorsDark();
extern void wrap_imnodes_StyleColorsLight();

#ifdef __cplusplus
}
#endif
