// Code generated by cmd/codegen from https://github.com/saveme712/cimgui-go.
// DO NOT EDIT.

#include "wrapper.h"
#include "../cwrappers/cimguizmo.h"

ImGuiID wrap_ImGuizmo_GetID_Ptr(const uintptr_t ptr_id) { return ImGuizmo_GetID_Ptr((const void*)(uintptr_t)ptr_id); }
void wrap_ImGuizmo_PushID_Ptr(const uintptr_t ptr_id) { ImGuizmo_PushID_Ptr((const void*)(uintptr_t)ptr_id); }
bool wrap_ImGuizmo_Manipulate(const float* view,const float* projection,OPERATION operation,MODE mode,float* matrix) { return ImGuizmo_Manipulate(view,projection,operation,mode,matrix,0,0,0,0); }
void wrap_ImGuizmo_SetDrawlist() { ImGuizmo_SetDrawlist(0); }
