/*
 * This content is Auto-Generated by cimgui-go's Make file.
 * DO NOT EDIT
 *
 * cimgui-go: https://github.com/saveme712/cimgui-go
 * From: templates/assert.h
 */
extern "C" void goImguiAssertHandler(char const *expression, char const *file, int line);
#define IM_ASSERT(_EXPR)                                   \
   do                                                      \
   {                                                       \
      if ((_EXPR) == 0)                                    \
      {                                                    \
         goImguiAssertHandler(#_EXPR, __FILE__, __LINE__); \
      }                                                    \
   } while (false)
