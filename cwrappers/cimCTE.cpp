//This file is automatically generated by generator.lua from https://github.com/cimgui/imnodes
//based on imnodes.h file version XXX from https://github.com/Nelarius/imnodes
#include "imgui.h"
#include "imgui_internal.h"
#include "./ImGuiColorTextEdit/TextEditor.h"
#include "cimCTE.h"



CIMGUI_API TextEditor* TextEditor_TextEditor(void)
{
    return IM_NEW(TextEditor)();
}
CIMGUI_API void TextEditor_destroy(TextEditor* self)
{
    IM_DELETE(self);
}
CIMGUI_API void TextEditor_SetReadOnlyEnabled(TextEditor* self,bool aValue)
{
    return self->SetReadOnlyEnabled(aValue);
}
CIMGUI_API bool TextEditor_IsReadOnlyEnabled(TextEditor* self)
{
    return self->IsReadOnlyEnabled();
}
CIMGUI_API void TextEditor_SetAutoIndentEnabled(TextEditor* self,bool aValue)
{
    return self->SetAutoIndentEnabled(aValue);
}
CIMGUI_API bool TextEditor_IsAutoIndentEnabled(TextEditor* self)
{
    return self->IsAutoIndentEnabled();
}
CIMGUI_API void TextEditor_SetShowWhitespacesEnabled(TextEditor* self,bool aValue)
{
    return self->SetShowWhitespacesEnabled(aValue);
}
CIMGUI_API bool TextEditor_IsShowWhitespacesEnabled(TextEditor* self)
{
    return self->IsShowWhitespacesEnabled();
}
CIMGUI_API void TextEditor_SetShowLineNumbersEnabled(TextEditor* self,bool aValue)
{
    return self->SetShowLineNumbersEnabled(aValue);
}
CIMGUI_API bool TextEditor_IsShowLineNumbersEnabled(TextEditor* self)
{
    return self->IsShowLineNumbersEnabled();
}
CIMGUI_API void TextEditor_SetShortTabsEnabled(TextEditor* self,bool aValue)
{
    return self->SetShortTabsEnabled(aValue);
}
CIMGUI_API bool TextEditor_IsShortTabsEnabled(TextEditor* self)
{
    return self->IsShortTabsEnabled();
}
CIMGUI_API int TextEditor_GetLineCount(TextEditor* self)
{
    return self->GetLineCount();
}
CIMGUI_API void TextEditor_SetPalette(TextEditor* self,PaletteId aValue)
{
    return self->SetPalette(aValue);
}
CIMGUI_API PaletteId TextEditor_GetPalette(TextEditor* self)
{
    return self->GetPalette();
}
CIMGUI_API void TextEditor_SetLanguageDefinition(TextEditor* self,LanguageDefinitionId aValue)
{
    return self->SetLanguageDefinition(aValue);
}
CIMGUI_API LanguageDefinitionId TextEditor_GetLanguageDefinition(TextEditor* self)
{
    return self->GetLanguageDefinition();
}
CIMGUI_API const char* TextEditor_GetLanguageDefinitionName(TextEditor* self)
{
    return self->GetLanguageDefinitionName();
}
CIMGUI_API void TextEditor_SetTabSize(TextEditor* self,int aValue)
{
    return self->SetTabSize(aValue);
}
CIMGUI_API int TextEditor_GetTabSize(TextEditor* self)
{
    return self->GetTabSize();
}
CIMGUI_API void TextEditor_SetLineSpacing(TextEditor* self,float aValue)
{
    return self->SetLineSpacing(aValue);
}
CIMGUI_API float TextEditor_GetLineSpacing(TextEditor* self)
{
    return self->GetLineSpacing();
}
CIMGUI_API void TextEditor_SetDefaultPalette(PaletteId aValue)
{
    return TextEditor::SetDefaultPalette(aValue);
}
CIMGUI_API PaletteId TextEditor_GetDefaultPalette()
{
    return TextEditor::GetDefaultPalette();
}
CIMGUI_API void TextEditor_SelectAll(TextEditor* self)
{
    return self->SelectAll();
}
CIMGUI_API void TextEditor_SelectLine(TextEditor* self,int aLine)
{
    return self->SelectLine(aLine);
}
CIMGUI_API void TextEditor_SelectRegion(TextEditor* self,int aStartLine,int aStartChar,int aEndLine,int aEndChar)
{
    return self->SelectRegion(aStartLine,aStartChar,aEndLine,aEndChar);
}
CIMGUI_API void TextEditor_SelectNextOccurrenceOf(TextEditor* self,const char* aText,int aTextSize,bool aCaseSensitive)
{
    return self->SelectNextOccurrenceOf(aText,aTextSize,aCaseSensitive);
}
CIMGUI_API void TextEditor_SelectAllOccurrencesOf(TextEditor* self,const char* aText,int aTextSize,bool aCaseSensitive)
{
    return self->SelectAllOccurrencesOf(aText,aTextSize,aCaseSensitive);
}
CIMGUI_API bool TextEditor_AnyCursorHasSelection(TextEditor* self)
{
    return self->AnyCursorHasSelection();
}
CIMGUI_API bool TextEditor_AllCursorsHaveSelection(TextEditor* self)
{
    return self->AllCursorsHaveSelection();
}
CIMGUI_API void TextEditor_ClearExtraCursors(TextEditor* self)
{
    return self->ClearExtraCursors();
}
CIMGUI_API void TextEditor_ClearSelections(TextEditor* self)
{
    return self->ClearSelections();
}
CIMGUI_API void TextEditor_SetCursorPosition(TextEditor* self,int aLine,int aCharIndex)
{
    return self->SetCursorPosition(aLine,aCharIndex);
}
CIMGUI_API void TextEditor_GetCursorPosition(TextEditor* self,int* outLine,int* outColumn)
{
    return self->GetCursorPosition(*outLine,*outColumn);
}
CIMGUI_API int TextEditor_GetFirstVisibleLine(TextEditor* self)
{
    return self->GetFirstVisibleLine();
}
CIMGUI_API int TextEditor_GetLastVisibleLine(TextEditor* self)
{
    return self->GetLastVisibleLine();
}
CIMGUI_API void TextEditor_SetViewAtLine(TextEditor* self,int aLine,SetViewAtLineMode aMode)
{
    return self->SetViewAtLine(aLine,aMode);
}
CIMGUI_API void TextEditor_Copy(TextEditor* self)
{
    return self->Copy();
}
CIMGUI_API void TextEditor_Cut(TextEditor* self)
{
    return self->Cut();
}
CIMGUI_API void TextEditor_Paste(TextEditor* self)
{
    return self->Paste();
}
CIMGUI_API void TextEditor_Undo(TextEditor* self,int aSteps)
{
    return self->Undo(aSteps);
}
CIMGUI_API void TextEditor_Redo(TextEditor* self,int aSteps)
{
    return self->Redo(aSteps);
}
CIMGUI_API bool TextEditor_CanUndo(TextEditor* self)
{
    return self->CanUndo();
}
CIMGUI_API bool TextEditor_CanRedo(TextEditor* self)
{
    return self->CanRedo();
}
CIMGUI_API int TextEditor_GetUndoIndex(TextEditor* self)
{
    return self->GetUndoIndex();
}
CIMGUI_API bool TextEditor_Render(TextEditor* self,const char* aTitle,bool aParentIsFocused,const ImVec2 aSize,bool aBorder)
{
    return self->Render(aTitle,aParentIsFocused,aSize,aBorder);
}
CIMGUI_API void TextEditor_UnitTests(TextEditor* self)
{
    return self->UnitTests();
}

////////////////manually generated
CIMGUI_API void TextEditor_SetText(TextEditor* self,const char* aText)
{
    return self->SetText(std::string(aText));
}
CIMGUI_API char* TextEditor_GetText_alloc(TextEditor* self)
{
    std::string str = self->GetText();
    char* cStr = (char*)IM_ALLOC(str.size() + 1); // Allocate memory
    std::strcpy(cStr, str.c_str()); // Copy string contents
    return cStr; // Return new C-style string
}
CIMGUI_API void TextEditor_GetText_free(char* ptr)
{
    IM_FREE(ptr); // free memory
}
CIMGUI_API const char* TextEditor_GetText_static(TextEditor* self)
{
    static std::string str = self->GetText();
    return str.c_str();
}
CIMGUI_API const char* TextEditor_GetText(TextEditor* self)
{
    static std::string str = self->GetText();
    return str.c_str();
}
CIMGUI_API void TextEditor_ImGuiDebugPanel(TextEditor* self,const char* panelName)
{
    return self->ImGuiDebugPanel(std::string(panelName));
}


