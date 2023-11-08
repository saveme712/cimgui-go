// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include <stdlib.h>
// #include <memory.h>
// #include "extra_types.h"
// #include "cimplot_wrapper.h"
import "C"

type FormatterTimeData struct {
	data *C.Formatter_Time_Data
}

func (self *FormatterTimeData) handle() (result *C.Formatter_Time_Data, fin func()) {
	return self.data, func() {}
}

func (self FormatterTimeData) c() (C.Formatter_Time_Data, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newFormatterTimeDataFromC(cvalue *C.Formatter_Time_Data) *FormatterTimeData {
	return &FormatterTimeData{data: cvalue}
}

type PlotAlignmentData struct {
	data *C.ImPlotAlignmentData
}

func (self *PlotAlignmentData) handle() (result *C.ImPlotAlignmentData, fin func()) {
	return self.data, func() {}
}

func (self PlotAlignmentData) c() (C.ImPlotAlignmentData, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotAlignmentDataFromC(cvalue *C.ImPlotAlignmentData) *PlotAlignmentData {
	return &PlotAlignmentData{data: cvalue}
}

type PlotAnnotation struct {
	data *C.ImPlotAnnotation
}

func (self *PlotAnnotation) handle() (result *C.ImPlotAnnotation, fin func()) {
	return self.data, func() {}
}

func (self PlotAnnotation) c() (C.ImPlotAnnotation, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotAnnotationFromC(cvalue *C.ImPlotAnnotation) *PlotAnnotation {
	return &PlotAnnotation{data: cvalue}
}

type PlotAnnotationCollection struct {
	data *C.ImPlotAnnotationCollection
}

func (self *PlotAnnotationCollection) handle() (result *C.ImPlotAnnotationCollection, fin func()) {
	return self.data, func() {}
}

func (self PlotAnnotationCollection) c() (C.ImPlotAnnotationCollection, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotAnnotationCollectionFromC(cvalue *C.ImPlotAnnotationCollection) *PlotAnnotationCollection {
	return &PlotAnnotationCollection{data: cvalue}
}

type PlotAxis struct {
	data *C.ImPlotAxis
}

func (self *PlotAxis) handle() (result *C.ImPlotAxis, fin func()) {
	return self.data, func() {}
}

func (self PlotAxis) c() (C.ImPlotAxis, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotAxisFromC(cvalue *C.ImPlotAxis) *PlotAxis {
	return &PlotAxis{data: cvalue}
}

type PlotAxisColor struct {
	data *C.ImPlotAxisColor
}

func (self *PlotAxisColor) handle() (result *C.ImPlotAxisColor, fin func()) {
	return self.data, func() {}
}

func (self PlotAxisColor) c() (C.ImPlotAxisColor, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotAxisColorFromC(cvalue *C.ImPlotAxisColor) *PlotAxisColor {
	return &PlotAxisColor{data: cvalue}
}

type PlotColormapData struct {
	data *C.ImPlotColormapData
}

func (self *PlotColormapData) handle() (result *C.ImPlotColormapData, fin func()) {
	return self.data, func() {}
}

func (self PlotColormapData) c() (C.ImPlotColormapData, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotColormapDataFromC(cvalue *C.ImPlotColormapData) *PlotColormapData {
	return &PlotColormapData{data: cvalue}
}

type PlotContext struct {
	data *C.ImPlotContext
}

func (self *PlotContext) handle() (result *C.ImPlotContext, fin func()) {
	return self.data, func() {}
}

func (self PlotContext) c() (C.ImPlotContext, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotContextFromC(cvalue *C.ImPlotContext) *PlotContext {
	return &PlotContext{data: cvalue}
}

type PlotDateTimeSpec struct {
	data *C.ImPlotDateTimeSpec
}

func (self *PlotDateTimeSpec) handle() (result *C.ImPlotDateTimeSpec, fin func()) {
	return self.data, func() {}
}

func (self PlotDateTimeSpec) c() (C.ImPlotDateTimeSpec, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotDateTimeSpecFromC(cvalue *C.ImPlotDateTimeSpec) *PlotDateTimeSpec {
	return &PlotDateTimeSpec{data: cvalue}
}

type PlotInputMap struct {
	data *C.ImPlotInputMap
}

func (self *PlotInputMap) handle() (result *C.ImPlotInputMap, fin func()) {
	return self.data, func() {}
}

func (self PlotInputMap) c() (C.ImPlotInputMap, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotInputMapFromC(cvalue *C.ImPlotInputMap) *PlotInputMap {
	return &PlotInputMap{data: cvalue}
}

type PlotItem struct {
	data *C.ImPlotItem
}

func (self *PlotItem) handle() (result *C.ImPlotItem, fin func()) {
	return self.data, func() {}
}

func (self PlotItem) c() (C.ImPlotItem, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotItemFromC(cvalue *C.ImPlotItem) *PlotItem {
	return &PlotItem{data: cvalue}
}

type PlotItemGroup struct {
	data *C.ImPlotItemGroup
}

func (self *PlotItemGroup) handle() (result *C.ImPlotItemGroup, fin func()) {
	return self.data, func() {}
}

func (self PlotItemGroup) c() (C.ImPlotItemGroup, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotItemGroupFromC(cvalue *C.ImPlotItemGroup) *PlotItemGroup {
	return &PlotItemGroup{data: cvalue}
}

type PlotLegend struct {
	data *C.ImPlotLegend
}

func (self *PlotLegend) handle() (result *C.ImPlotLegend, fin func()) {
	return self.data, func() {}
}

func (self PlotLegend) c() (C.ImPlotLegend, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotLegendFromC(cvalue *C.ImPlotLegend) *PlotLegend {
	return &PlotLegend{data: cvalue}
}

type PlotNextItemData struct {
	data *C.ImPlotNextItemData
}

func (self *PlotNextItemData) handle() (result *C.ImPlotNextItemData, fin func()) {
	return self.data, func() {}
}

func (self PlotNextItemData) c() (C.ImPlotNextItemData, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotNextItemDataFromC(cvalue *C.ImPlotNextItemData) *PlotNextItemData {
	return &PlotNextItemData{data: cvalue}
}

type PlotNextPlotData struct {
	data *C.ImPlotNextPlotData
}

func (self *PlotNextPlotData) handle() (result *C.ImPlotNextPlotData, fin func()) {
	return self.data, func() {}
}

func (self PlotNextPlotData) c() (C.ImPlotNextPlotData, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotNextPlotDataFromC(cvalue *C.ImPlotNextPlotData) *PlotNextPlotData {
	return &PlotNextPlotData{data: cvalue}
}

type PlotPlot struct {
	data *C.ImPlotPlot
}

func (self *PlotPlot) handle() (result *C.ImPlotPlot, fin func()) {
	return self.data, func() {}
}

func (self PlotPlot) c() (C.ImPlotPlot, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotPlotFromC(cvalue *C.ImPlotPlot) *PlotPlot {
	return &PlotPlot{data: cvalue}
}

type PlotPointError struct {
	data *C.ImPlotPointError
}

func (self *PlotPointError) handle() (result *C.ImPlotPointError, fin func()) {
	return self.data, func() {}
}

func (self PlotPointError) c() (C.ImPlotPointError, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotPointErrorFromC(cvalue *C.ImPlotPointError) *PlotPointError {
	return &PlotPointError{data: cvalue}
}

type PlotRange struct {
	data *C.ImPlotRange
}

func (self *PlotRange) handle() (result *C.ImPlotRange, fin func()) {
	return self.data, func() {}
}

func (self PlotRange) c() (C.ImPlotRange, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotRangeFromC(cvalue *C.ImPlotRange) *PlotRange {
	return &PlotRange{data: cvalue}
}

type PlotRect struct {
	data *C.ImPlotRect
}

func (self *PlotRect) handle() (result *C.ImPlotRect, fin func()) {
	return self.data, func() {}
}

func (self PlotRect) c() (C.ImPlotRect, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotRectFromC(cvalue *C.ImPlotRect) *PlotRect {
	return &PlotRect{data: cvalue}
}

type PlotStyle struct {
	data *C.ImPlotStyle
}

func (self *PlotStyle) handle() (result *C.ImPlotStyle, fin func()) {
	return self.data, func() {}
}

func (self PlotStyle) c() (C.ImPlotStyle, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotStyleFromC(cvalue *C.ImPlotStyle) *PlotStyle {
	return &PlotStyle{data: cvalue}
}

type PlotSubplot struct {
	data *C.ImPlotSubplot
}

func (self *PlotSubplot) handle() (result *C.ImPlotSubplot, fin func()) {
	return self.data, func() {}
}

func (self PlotSubplot) c() (C.ImPlotSubplot, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotSubplotFromC(cvalue *C.ImPlotSubplot) *PlotSubplot {
	return &PlotSubplot{data: cvalue}
}

type PlotTag struct {
	data *C.ImPlotTag
}

func (self *PlotTag) handle() (result *C.ImPlotTag, fin func()) {
	return self.data, func() {}
}

func (self PlotTag) c() (C.ImPlotTag, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotTagFromC(cvalue *C.ImPlotTag) *PlotTag {
	return &PlotTag{data: cvalue}
}

type PlotTagCollection struct {
	data *C.ImPlotTagCollection
}

func (self *PlotTagCollection) handle() (result *C.ImPlotTagCollection, fin func()) {
	return self.data, func() {}
}

func (self PlotTagCollection) c() (C.ImPlotTagCollection, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotTagCollectionFromC(cvalue *C.ImPlotTagCollection) *PlotTagCollection {
	return &PlotTagCollection{data: cvalue}
}

type PlotTick struct {
	data *C.ImPlotTick
}

func (self *PlotTick) handle() (result *C.ImPlotTick, fin func()) {
	return self.data, func() {}
}

func (self PlotTick) c() (C.ImPlotTick, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotTickFromC(cvalue *C.ImPlotTick) *PlotTick {
	return &PlotTick{data: cvalue}
}

type PlotTicker struct {
	data *C.ImPlotTicker
}

func (self *PlotTicker) handle() (result *C.ImPlotTicker, fin func()) {
	return self.data, func() {}
}

func (self PlotTicker) c() (C.ImPlotTicker, func()) {
	result, fn := self.handle()
	return *result, fn
}

func newPlotTickerFromC(cvalue *C.ImPlotTicker) *PlotTicker {
	return &PlotTicker{data: cvalue}
}
