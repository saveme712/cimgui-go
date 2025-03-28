// Code generated by cmd/codegen from https://github.com/saveme712/cimgui-go.
// DO NOT EDIT.

package immarkdown

// original name: EmphasisState
type EmphasisState int32

const (
	NONE   EmphasisState = 0
	LEFT   EmphasisState = 1
	MIDDLE EmphasisState = 2
	RIGHT  EmphasisState = 3
)

// original name: LinkState
type LinkState int32

const (
	NOLINK                            LinkState = 0
	HASSQUAREBRACKETOPEN              LinkState = 1
	HASSQUAREBRACKETS                 LinkState = 2
	HASSQUAREBRACKETSROUNDBRACKETOPEN LinkState = 3
)

// original name: MarkdownFormatType
type MarkdownFormatType int32

const (
	NORMALTEXT    MarkdownFormatType = 0
	HEADING       MarkdownFormatType = 1
	UNORDEREDLIST MarkdownFormatType = 2
	LINK          MarkdownFormatType = 3
	EMPHASIS      MarkdownFormatType = 4
)
