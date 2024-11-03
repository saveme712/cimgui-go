// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package implot

// #include <stdlib.h>
// #include <memory.h>
// #include "../imgui/extra_types.h"
// #include "cimplot_wrapper.h"
// #include "cimplot_typedefs.h"
// extern int callbackPlotFormatter0(double, char*, int, void*);
// extern int callbackPlotFormatter1(double, char*, int, void*);
// extern ImPlotPoint callbackPlotGetter0(int, void*);
// extern ImPlotPoint callbackPlotGetter1(int, void*);
// extern double callbackPlotTransform0(double, void*);
// extern double callbackPlotTransform1(double, void*);
import "C"
import (
	"unsafe"

	"github.com/AllenDang/cimgui-go/internal"
)

type FormatterTimeData struct {
	CData *C.Formatter_Time_Data
}

// Handle returns C version of FormatterTimeData and its finalizer func.
func (self *FormatterTimeData) Handle() (result *C.Formatter_Time_Data, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self FormatterTimeData) C() (C.Formatter_Time_Data, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewFormatterTimeDataFromC creates FormatterTimeData from its C pointer.
// SRC ~= *C.Formatter_Time_Data
func NewFormatterTimeDataFromC[SRC any](cvalue SRC) *FormatterTimeData {
	return &FormatterTimeData{CData: internal.ReinterpretCast[*C.Formatter_Time_Data](cvalue)}
}

type PlotAlignmentData struct {
	CData *C.ImPlotAlignmentData
}

// Handle returns C version of PlotAlignmentData and its finalizer func.
func (self *PlotAlignmentData) Handle() (result *C.ImPlotAlignmentData, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotAlignmentData) C() (C.ImPlotAlignmentData, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotAlignmentDataFromC creates PlotAlignmentData from its C pointer.
// SRC ~= *C.ImPlotAlignmentData
func NewPlotAlignmentDataFromC[SRC any](cvalue SRC) *PlotAlignmentData {
	return &PlotAlignmentData{CData: internal.ReinterpretCast[*C.ImPlotAlignmentData](cvalue)}
}

type PlotAnnotation struct {
	CData *C.ImPlotAnnotation
}

// Handle returns C version of PlotAnnotation and its finalizer func.
func (self *PlotAnnotation) Handle() (result *C.ImPlotAnnotation, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotAnnotation) C() (C.ImPlotAnnotation, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotAnnotationFromC creates PlotAnnotation from its C pointer.
// SRC ~= *C.ImPlotAnnotation
func NewPlotAnnotationFromC[SRC any](cvalue SRC) *PlotAnnotation {
	return &PlotAnnotation{CData: internal.ReinterpretCast[*C.ImPlotAnnotation](cvalue)}
}

type PlotAnnotationCollection struct {
	CData *C.ImPlotAnnotationCollection
}

// Handle returns C version of PlotAnnotationCollection and its finalizer func.
func (self *PlotAnnotationCollection) Handle() (result *C.ImPlotAnnotationCollection, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotAnnotationCollection) C() (C.ImPlotAnnotationCollection, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotAnnotationCollectionFromC creates PlotAnnotationCollection from its C pointer.
// SRC ~= *C.ImPlotAnnotationCollection
func NewPlotAnnotationCollectionFromC[SRC any](cvalue SRC) *PlotAnnotationCollection {
	return &PlotAnnotationCollection{CData: internal.ReinterpretCast[*C.ImPlotAnnotationCollection](cvalue)}
}

type PlotAxis struct {
	CData *C.ImPlotAxis
}

// Handle returns C version of PlotAxis and its finalizer func.
func (self *PlotAxis) Handle() (result *C.ImPlotAxis, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotAxis) C() (C.ImPlotAxis, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotAxisFromC creates PlotAxis from its C pointer.
// SRC ~= *C.ImPlotAxis
func NewPlotAxisFromC[SRC any](cvalue SRC) *PlotAxis {
	return &PlotAxis{CData: internal.ReinterpretCast[*C.ImPlotAxis](cvalue)}
}

type PlotAxisColor struct {
	CData *C.ImPlotAxisColor
}

// Handle returns C version of PlotAxisColor and its finalizer func.
func (self *PlotAxisColor) Handle() (result *C.ImPlotAxisColor, fin func()) {
	return self.CData, func() {}
}

// NewPlotAxisColorFromC creates PlotAxisColor from its C pointer.
// SRC ~= *C.ImPlotAxisColor
func NewPlotAxisColorFromC[SRC any](cvalue SRC) *PlotAxisColor {
	return &PlotAxisColor{CData: internal.ReinterpretCast[*C.ImPlotAxisColor](cvalue)}
}

type PlotColormapData struct {
	CData *C.ImPlotColormapData
}

// Handle returns C version of PlotColormapData and its finalizer func.
func (self *PlotColormapData) Handle() (result *C.ImPlotColormapData, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotColormapData) C() (C.ImPlotColormapData, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotColormapDataFromC creates PlotColormapData from its C pointer.
// SRC ~= *C.ImPlotColormapData
func NewPlotColormapDataFromC[SRC any](cvalue SRC) *PlotColormapData {
	return &PlotColormapData{CData: internal.ReinterpretCast[*C.ImPlotColormapData](cvalue)}
}

type PlotContext struct {
	CData *C.ImPlotContext
}

// Handle returns C version of PlotContext and its finalizer func.
func (self *PlotContext) Handle() (result *C.ImPlotContext, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotContext) C() (C.ImPlotContext, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotContextFromC creates PlotContext from its C pointer.
// SRC ~= *C.ImPlotContext
func NewPlotContextFromC[SRC any](cvalue SRC) *PlotContext {
	return &PlotContext{CData: internal.ReinterpretCast[*C.ImPlotContext](cvalue)}
}

type PlotDateTimeSpec struct {
	CData *C.ImPlotDateTimeSpec
}

// Handle returns C version of PlotDateTimeSpec and its finalizer func.
func (self *PlotDateTimeSpec) Handle() (result *C.ImPlotDateTimeSpec, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotDateTimeSpec) C() (C.ImPlotDateTimeSpec, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotDateTimeSpecFromC creates PlotDateTimeSpec from its C pointer.
// SRC ~= *C.ImPlotDateTimeSpec
func NewPlotDateTimeSpecFromC[SRC any](cvalue SRC) *PlotDateTimeSpec {
	return &PlotDateTimeSpec{CData: internal.ReinterpretCast[*C.ImPlotDateTimeSpec](cvalue)}
}

type (
	PlotFormatter  func(value float64, buff string, size int32, user_data unsafe.Pointer) int32
	cPlotFormatter func(value C.double, buff *C.char, size C.int, user_data unsafe.Pointer) C.int
)

func NewPlotFormatterFromC(cvalue *C.ImPlotFormatter) *PlotFormatter {
	result := poolPlotFormatter.Find(*cvalue)
	return &result
}

func (c PlotFormatter) C() (C.ImPlotFormatter, func()) {
	return poolPlotFormatter.Allocate(c), func() {}
}

func wrapPlotFormatter(cb PlotFormatter, value C.double, buff *C.char, size C.int, user_data unsafe.Pointer) C.int {
	result := cb(float64(value), C.GoString(buff), int32(size), unsafe.Pointer(user_data))

	return C.int(result)
}

//export callbackPlotFormatter0
func callbackPlotFormatter0(value C.double, buff *C.char, size C.int, user_data unsafe.Pointer) C.int {
	return wrapPlotFormatter(poolPlotFormatter.Get(0), value, buff, size, user_data)
}

//export callbackPlotFormatter1
func callbackPlotFormatter1(value C.double, buff *C.char, size C.int, user_data unsafe.Pointer) C.int {
	return wrapPlotFormatter(poolPlotFormatter.Get(1), value, buff, size, user_data)
}

var poolPlotFormatter *internal.Pool[PlotFormatter, C.ImPlotFormatter]

func init() {
	poolPlotFormatter = internal.NewPool[PlotFormatter, C.ImPlotFormatter](
		C.ImPlotFormatter(C.callbackPlotFormatter0),
		C.ImPlotFormatter(C.callbackPlotFormatter1),
	)
}

func ClearPlotFormatterPool() {
	poolPlotFormatter.Clear()
}

type (
	PlotGetter  func(idx int32, user_data unsafe.Pointer) PlotPoint
	cPlotGetter func(idx C.int, user_data unsafe.Pointer) C.ImPlotPoint
)

func NewPlotGetterFromC(cvalue *C.ImPlotGetter) *PlotGetter {
	result := poolPlotGetter.Find(*cvalue)
	return &result
}

func (c PlotGetter) C() (C.ImPlotGetter, func()) {
	return poolPlotGetter.Allocate(c), func() {}
}

func wrapPlotGetter(cb PlotGetter, idx C.int, user_data unsafe.Pointer) C.ImPlotPoint {
	result := cb(int32(idx), unsafe.Pointer(user_data))

	return internal.ReinterpretCast[C.ImPlotPoint](result.ToC())
}

//export callbackPlotGetter0
func callbackPlotGetter0(idx C.int, user_data unsafe.Pointer) C.ImPlotPoint {
	return wrapPlotGetter(poolPlotGetter.Get(0), idx, user_data)
}

//export callbackPlotGetter1
func callbackPlotGetter1(idx C.int, user_data unsafe.Pointer) C.ImPlotPoint {
	return wrapPlotGetter(poolPlotGetter.Get(1), idx, user_data)
}

var poolPlotGetter *internal.Pool[PlotGetter, C.ImPlotGetter]

func init() {
	poolPlotGetter = internal.NewPool[PlotGetter, C.ImPlotGetter](
		C.ImPlotGetter(C.callbackPlotGetter0),
		C.ImPlotGetter(C.callbackPlotGetter1),
	)
}

func ClearPlotGetterPool() {
	poolPlotGetter.Clear()
}

type PlotInputMap struct {
	CData *C.ImPlotInputMap
}

// Handle returns C version of PlotInputMap and its finalizer func.
func (self *PlotInputMap) Handle() (result *C.ImPlotInputMap, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotInputMap) C() (C.ImPlotInputMap, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotInputMapFromC creates PlotInputMap from its C pointer.
// SRC ~= *C.ImPlotInputMap
func NewPlotInputMapFromC[SRC any](cvalue SRC) *PlotInputMap {
	return &PlotInputMap{CData: internal.ReinterpretCast[*C.ImPlotInputMap](cvalue)}
}

type PlotItem struct {
	CData *C.ImPlotItem
}

// Handle returns C version of PlotItem and its finalizer func.
func (self *PlotItem) Handle() (result *C.ImPlotItem, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotItem) C() (C.ImPlotItem, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotItemFromC creates PlotItem from its C pointer.
// SRC ~= *C.ImPlotItem
func NewPlotItemFromC[SRC any](cvalue SRC) *PlotItem {
	return &PlotItem{CData: internal.ReinterpretCast[*C.ImPlotItem](cvalue)}
}

type PlotItemGroup struct {
	CData *C.ImPlotItemGroup
}

// Handle returns C version of PlotItemGroup and its finalizer func.
func (self *PlotItemGroup) Handle() (result *C.ImPlotItemGroup, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotItemGroup) C() (C.ImPlotItemGroup, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotItemGroupFromC creates PlotItemGroup from its C pointer.
// SRC ~= *C.ImPlotItemGroup
func NewPlotItemGroupFromC[SRC any](cvalue SRC) *PlotItemGroup {
	return &PlotItemGroup{CData: internal.ReinterpretCast[*C.ImPlotItemGroup](cvalue)}
}

type PlotLegend struct {
	CData *C.ImPlotLegend
}

// Handle returns C version of PlotLegend and its finalizer func.
func (self *PlotLegend) Handle() (result *C.ImPlotLegend, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotLegend) C() (C.ImPlotLegend, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotLegendFromC creates PlotLegend from its C pointer.
// SRC ~= *C.ImPlotLegend
func NewPlotLegendFromC[SRC any](cvalue SRC) *PlotLegend {
	return &PlotLegend{CData: internal.ReinterpretCast[*C.ImPlotLegend](cvalue)}
}

type PlotNextItemData struct {
	CData *C.ImPlotNextItemData
}

// Handle returns C version of PlotNextItemData and its finalizer func.
func (self *PlotNextItemData) Handle() (result *C.ImPlotNextItemData, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotNextItemData) C() (C.ImPlotNextItemData, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotNextItemDataFromC creates PlotNextItemData from its C pointer.
// SRC ~= *C.ImPlotNextItemData
func NewPlotNextItemDataFromC[SRC any](cvalue SRC) *PlotNextItemData {
	return &PlotNextItemData{CData: internal.ReinterpretCast[*C.ImPlotNextItemData](cvalue)}
}

type PlotNextPlotData struct {
	CData *C.ImPlotNextPlotData
}

// Handle returns C version of PlotNextPlotData and its finalizer func.
func (self *PlotNextPlotData) Handle() (result *C.ImPlotNextPlotData, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotNextPlotData) C() (C.ImPlotNextPlotData, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotNextPlotDataFromC creates PlotNextPlotData from its C pointer.
// SRC ~= *C.ImPlotNextPlotData
func NewPlotNextPlotDataFromC[SRC any](cvalue SRC) *PlotNextPlotData {
	return &PlotNextPlotData{CData: internal.ReinterpretCast[*C.ImPlotNextPlotData](cvalue)}
}

type PlotPlot struct {
	CData *C.ImPlotPlot
}

// Handle returns C version of PlotPlot and its finalizer func.
func (self *PlotPlot) Handle() (result *C.ImPlotPlot, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotPlot) C() (C.ImPlotPlot, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotPlotFromC creates PlotPlot from its C pointer.
// SRC ~= *C.ImPlotPlot
func NewPlotPlotFromC[SRC any](cvalue SRC) *PlotPlot {
	return &PlotPlot{CData: internal.ReinterpretCast[*C.ImPlotPlot](cvalue)}
}

type PlotPointError struct {
	CData *C.ImPlotPointError
}

// Handle returns C version of PlotPointError and its finalizer func.
func (self *PlotPointError) Handle() (result *C.ImPlotPointError, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotPointError) C() (C.ImPlotPointError, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotPointErrorFromC creates PlotPointError from its C pointer.
// SRC ~= *C.ImPlotPointError
func NewPlotPointErrorFromC[SRC any](cvalue SRC) *PlotPointError {
	return &PlotPointError{CData: internal.ReinterpretCast[*C.ImPlotPointError](cvalue)}
}

type PlotRange struct {
	CData *C.ImPlotRange
}

// Handle returns C version of PlotRange and its finalizer func.
func (self *PlotRange) Handle() (result *C.ImPlotRange, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotRange) C() (C.ImPlotRange, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotRangeFromC creates PlotRange from its C pointer.
// SRC ~= *C.ImPlotRange
func NewPlotRangeFromC[SRC any](cvalue SRC) *PlotRange {
	return &PlotRange{CData: internal.ReinterpretCast[*C.ImPlotRange](cvalue)}
}

type PlotRect struct {
	CData *C.ImPlotRect
}

// Handle returns C version of PlotRect and its finalizer func.
func (self *PlotRect) Handle() (result *C.ImPlotRect, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotRect) C() (C.ImPlotRect, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotRectFromC creates PlotRect from its C pointer.
// SRC ~= *C.ImPlotRect
func NewPlotRectFromC[SRC any](cvalue SRC) *PlotRect {
	return &PlotRect{CData: internal.ReinterpretCast[*C.ImPlotRect](cvalue)}
}

type PlotStyle struct {
	CData *C.ImPlotStyle
}

// Handle returns C version of PlotStyle and its finalizer func.
func (self *PlotStyle) Handle() (result *C.ImPlotStyle, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotStyle) C() (C.ImPlotStyle, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotStyleFromC creates PlotStyle from its C pointer.
// SRC ~= *C.ImPlotStyle
func NewPlotStyleFromC[SRC any](cvalue SRC) *PlotStyle {
	return &PlotStyle{CData: internal.ReinterpretCast[*C.ImPlotStyle](cvalue)}
}

type PlotSubplot struct {
	CData *C.ImPlotSubplot
}

// Handle returns C version of PlotSubplot and its finalizer func.
func (self *PlotSubplot) Handle() (result *C.ImPlotSubplot, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotSubplot) C() (C.ImPlotSubplot, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotSubplotFromC creates PlotSubplot from its C pointer.
// SRC ~= *C.ImPlotSubplot
func NewPlotSubplotFromC[SRC any](cvalue SRC) *PlotSubplot {
	return &PlotSubplot{CData: internal.ReinterpretCast[*C.ImPlotSubplot](cvalue)}
}

type PlotTag struct {
	CData *C.ImPlotTag
}

// Handle returns C version of PlotTag and its finalizer func.
func (self *PlotTag) Handle() (result *C.ImPlotTag, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotTag) C() (C.ImPlotTag, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotTagFromC creates PlotTag from its C pointer.
// SRC ~= *C.ImPlotTag
func NewPlotTagFromC[SRC any](cvalue SRC) *PlotTag {
	return &PlotTag{CData: internal.ReinterpretCast[*C.ImPlotTag](cvalue)}
}

type PlotTagCollection struct {
	CData *C.ImPlotTagCollection
}

// Handle returns C version of PlotTagCollection and its finalizer func.
func (self *PlotTagCollection) Handle() (result *C.ImPlotTagCollection, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotTagCollection) C() (C.ImPlotTagCollection, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotTagCollectionFromC creates PlotTagCollection from its C pointer.
// SRC ~= *C.ImPlotTagCollection
func NewPlotTagCollectionFromC[SRC any](cvalue SRC) *PlotTagCollection {
	return &PlotTagCollection{CData: internal.ReinterpretCast[*C.ImPlotTagCollection](cvalue)}
}

type PlotTick struct {
	CData *C.ImPlotTick
}

// Handle returns C version of PlotTick and its finalizer func.
func (self *PlotTick) Handle() (result *C.ImPlotTick, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotTick) C() (C.ImPlotTick, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotTickFromC creates PlotTick from its C pointer.
// SRC ~= *C.ImPlotTick
func NewPlotTickFromC[SRC any](cvalue SRC) *PlotTick {
	return &PlotTick{CData: internal.ReinterpretCast[*C.ImPlotTick](cvalue)}
}

type PlotTicker struct {
	CData *C.ImPlotTicker
}

// Handle returns C version of PlotTicker and its finalizer func.
func (self *PlotTicker) Handle() (result *C.ImPlotTicker, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self PlotTicker) C() (C.ImPlotTicker, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewPlotTickerFromC creates PlotTicker from its C pointer.
// SRC ~= *C.ImPlotTicker
func NewPlotTickerFromC[SRC any](cvalue SRC) *PlotTicker {
	return &PlotTicker{CData: internal.ReinterpretCast[*C.ImPlotTicker](cvalue)}
}

type (
	PlotTransform  func(value float64, user_data unsafe.Pointer) float64
	cPlotTransform func(value C.double, user_data unsafe.Pointer) C.double
)

func NewPlotTransformFromC(cvalue *C.ImPlotTransform) *PlotTransform {
	result := poolPlotTransform.Find(*cvalue)
	return &result
}

func (c PlotTransform) C() (C.ImPlotTransform, func()) {
	return poolPlotTransform.Allocate(c), func() {}
}

func wrapPlotTransform(cb PlotTransform, value C.double, user_data unsafe.Pointer) C.double {
	result := cb(float64(value), unsafe.Pointer(user_data))

	return C.double(result)
}

//export callbackPlotTransform0
func callbackPlotTransform0(value C.double, user_data unsafe.Pointer) C.double {
	return wrapPlotTransform(poolPlotTransform.Get(0), value, user_data)
}

//export callbackPlotTransform1
func callbackPlotTransform1(value C.double, user_data unsafe.Pointer) C.double {
	return wrapPlotTransform(poolPlotTransform.Get(1), value, user_data)
}

var poolPlotTransform *internal.Pool[PlotTransform, C.ImPlotTransform]

func init() {
	poolPlotTransform = internal.NewPool[PlotTransform, C.ImPlotTransform](
		C.ImPlotTransform(C.callbackPlotTransform0),
		C.ImPlotTransform(C.callbackPlotTransform1),
	)
}

func ClearPlotTransformPool() {
	poolPlotTransform.Clear()
}
