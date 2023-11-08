// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "cimmarkdown_wrapper.h"
import "C"

type Emphasis struct {
	data *C.Emphasis
}

func (self *Emphasis) handle() (result *C.Emphasis, fin func()) {
	return self.data, func() {}
}

func (self Emphasis) c() (C.Emphasis, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newEmphasisFromC(cvalue *C.Emphasis) *Emphasis {
	return &Emphasis{data: cvalue}
}

type Line struct {
	data *C.Line
}

func (self *Line) handle() (result *C.Line, fin func()) {
	return self.data, func() {}
}

func (self Line) c() (C.Line, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newLineFromC(cvalue *C.Line) *Line {
	return &Line{data: cvalue}
}

type Link struct {
	data *C.Link
}

func (self *Link) handle() (result *C.Link, fin func()) {
	return self.data, func() {}
}

func (self Link) c() (C.Link, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newLinkFromC(cvalue *C.Link) *Link {
	return &Link{data: cvalue}
}

type MarkdownConfig struct {
	data *C.MarkdownConfig
}

func (self *MarkdownConfig) handle() (result *C.MarkdownConfig, fin func()) {
	return self.data, func() {}
}

func (self MarkdownConfig) c() (C.MarkdownConfig, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newMarkdownConfigFromC(cvalue *C.MarkdownConfig) *MarkdownConfig {
	return &MarkdownConfig{data: cvalue}
}

type MarkdownFormatInfo struct {
	data *C.MarkdownFormatInfo
}

func (self *MarkdownFormatInfo) handle() (result *C.MarkdownFormatInfo, fin func()) {
	return self.data, func() {}
}

func (self MarkdownFormatInfo) c() (C.MarkdownFormatInfo, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newMarkdownFormatInfoFromC(cvalue *C.MarkdownFormatInfo) *MarkdownFormatInfo {
	return &MarkdownFormatInfo{data: cvalue}
}

type MarkdownHeadingFormat struct {
	data *C.MarkdownHeadingFormat
}

func (self *MarkdownHeadingFormat) handle() (result *C.MarkdownHeadingFormat, fin func()) {
	return self.data, func() {}
}

func (self MarkdownHeadingFormat) c() (C.MarkdownHeadingFormat, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newMarkdownHeadingFormatFromC(cvalue *C.MarkdownHeadingFormat) *MarkdownHeadingFormat {
	return &MarkdownHeadingFormat{data: cvalue}
}

type MarkdownImageData struct {
	data *C.MarkdownImageData
}

func (self *MarkdownImageData) handle() (result *C.MarkdownImageData, fin func()) {
	return self.data, func() {}
}

func (self MarkdownImageData) c() (C.MarkdownImageData, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newMarkdownImageDataFromC(cvalue *C.MarkdownImageData) *MarkdownImageData {
	return &MarkdownImageData{data: cvalue}
}

type MarkdownLinkCallbackData struct {
	data *C.MarkdownLinkCallbackData
}

func (self *MarkdownLinkCallbackData) handle() (result *C.MarkdownLinkCallbackData, fin func()) {
	return self.data, func() {}
}

func (self MarkdownLinkCallbackData) c() (C.MarkdownLinkCallbackData, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newMarkdownLinkCallbackDataFromC(cvalue *C.MarkdownLinkCallbackData) *MarkdownLinkCallbackData {
	return &MarkdownLinkCallbackData{data: cvalue}
}

type MarkdownTooltipCallbackData struct {
	data *C.MarkdownTooltipCallbackData
}

func (self *MarkdownTooltipCallbackData) handle() (result *C.MarkdownTooltipCallbackData, fin func()) {
	return self.data, func() {}
}

func (self MarkdownTooltipCallbackData) c() (C.MarkdownTooltipCallbackData, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newMarkdownTooltipCallbackDataFromC(cvalue *C.MarkdownTooltipCallbackData) *MarkdownTooltipCallbackData {
	return &MarkdownTooltipCallbackData{data: cvalue}
}

type TextBlock struct {
	data *C.TextBlock
}

func (self *TextBlock) handle() (result *C.TextBlock, fin func()) {
	return self.data, func() {}
}

func (self TextBlock) c() (C.TextBlock, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newTextBlockFromC(cvalue *C.TextBlock) *TextBlock {
	return &TextBlock{data: cvalue}
}

type TextRegion struct {
	data *C.TextRegion
}

func (self *TextRegion) handle() (result *C.TextRegion, fin func()) {
	return self.data, func() {}
}

func (self TextRegion) c() (C.TextRegion, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newTextRegionFromC(cvalue *C.TextRegion) *TextRegion {
	return &TextRegion{data: cvalue}
}
