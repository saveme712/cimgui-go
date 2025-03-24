// Code generated by cmd/codegen from https://github.com/saveme712/cimgui-go.
// DO NOT EDIT.

#include "wrapper.h"
#include "../cwrappers/cimmarkdown.h"

bool wrap_RenderLinkText(TextRegion* self,const char* text_,const Link link_,const char* markdown_,const MarkdownConfig mdConfig_,const char** linkHoverStart_) { return RenderLinkText(self,text_,0,link_,markdown_,mdConfig_,linkHoverStart_); }
void wrap_RenderLinkTextWrappedV(TextRegion* self,const char* text_,const Link link_,const char* markdown_,const MarkdownConfig mdConfig_,const char** linkHoverStart_,bool bIndentToHere_) { RenderLinkTextWrapped(self,text_,0,link_,markdown_,mdConfig_,linkHoverStart_,bIndentToHere_); }
void wrap_RenderListTextWrapped(TextRegion* self,const char* text_) { RenderListTextWrapped(self,text_,0); }
void wrap_RenderTextWrappedV(TextRegion* self,const char* text_,bool bIndentToHere_) { RenderTextWrapped(self,text_,0,bIndentToHere_); }
void wrap_RenderLinkTextWrapped(TextRegion* self,const char* text_,const Link link_,const char* markdown_,const MarkdownConfig mdConfig_,const char** linkHoverStart_) { wrap_RenderLinkTextWrappedV(self,text_,link_,markdown_,mdConfig_,linkHoverStart_,false); }
void wrap_RenderTextWrapped(TextRegion* self,const char* text_) { wrap_RenderTextWrappedV(self,text_,false); }
