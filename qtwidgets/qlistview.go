package qtwidgets

// /usr/include/qt/QtWidgets/qlistview.h
// #include <qlistview.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QListView) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// void scrollContentsBy(int, int)
func (this *QListView) InheritScrollContentsBy(f func(dx int, dy int)) {
	ffiqt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// void resizeContents(int, int)
func (this *QListView) InheritResizeContents(f func(width int, height int)) {
	ffiqt.SetAllInheritCallback(this, "resizeContents", f)
}

// QSize contentsSize()
func (this *QListView) InheritContentsSize(f func() unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "contentsSize", f)
}

// void rowsInserted(const class QModelIndex &, int, int)
func (this *QListView) InheritRowsInserted(f func(parent *qtcore.QModelIndex, start int, end int)) {
	ffiqt.SetAllInheritCallback(this, "rowsInserted", f)
}

// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QListView) InheritRowsAboutToBeRemoved(f func(parent *qtcore.QModelIndex, start int, end int)) {
	ffiqt.SetAllInheritCallback(this, "rowsAboutToBeRemoved", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QListView) InheritMouseMoveEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QListView) InheritMouseReleaseEvent(f func(e *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QListView) InheritWheelEvent(f func(e *qtgui.QWheelEvent /*777 QWheelEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QListView) InheritTimerEvent(f func(e *qtcore.QTimerEvent /*777 QTimerEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "timerEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QListView) InheritResizeEvent(f func(e *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QListView) InheritDragMoveEvent(f func(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QListView) InheritDragLeaveEvent(f func(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QListView) InheritDropEvent(f func(e *qtgui.QDropEvent /*777 QDropEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dropEvent", f)
}

// QStyleOptionViewItem viewOptions()
func (this *QListView) InheritViewOptions(f func() unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "viewOptions", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QListView) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// int horizontalOffset()
func (this *QListView) InheritHorizontalOffset(f func() int) {
	ffiqt.SetAllInheritCallback(this, "horizontalOffset", f)
}

// int verticalOffset()
func (this *QListView) InheritVerticalOffset(f func() int) {
	ffiqt.SetAllInheritCallback(this, "verticalOffset", f)
}

// QRect rectForIndex(const class QModelIndex &)
func (this *QListView) InheritRectForIndex(f func(index *qtcore.QModelIndex) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "rectForIndex", f)
}

// void setPositionForIndex(const class QPoint &, const class QModelIndex &)
func (this *QListView) InheritSetPositionForIndex(f func(position *qtcore.QPoint, index *qtcore.QModelIndex)) {
	ffiqt.SetAllInheritCallback(this, "setPositionForIndex", f)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QListView) InheritSetSelection(f func(rect *qtcore.QRect, command int)) {
	ffiqt.SetAllInheritCallback(this, "setSelection", f)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QListView) InheritVisualRegionForSelection(f func(selection *qtcore.QItemSelection) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "visualRegionForSelection", f)
}

// void updateGeometries()
func (this *QListView) InheritUpdateGeometries(f func()) {
	ffiqt.SetAllInheritCallback(this, "updateGeometries", f)
}

// bool isIndexHidden(const class QModelIndex &)
func (this *QListView) InheritIsIndexHidden(f func(index *qtcore.QModelIndex) bool) {
	ffiqt.SetAllInheritCallback(this, "isIndexHidden", f)
}

// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QListView) InheritSelectionChanged(f func(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection)) {
	ffiqt.SetAllInheritCallback(this, "selectionChanged", f)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QListView) InheritCurrentChanged(f func(current *qtcore.QModelIndex, previous *qtcore.QModelIndex)) {
	ffiqt.SetAllInheritCallback(this, "currentChanged", f)
}

// QSize viewportSizeHint()
func (this *QListView) InheritViewportSizeHint(f func() unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

type QListView struct {
	*QAbstractItemView
}

func (this *QListView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemView.GetCthis()
	}
}
func (this *QListView) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemView = NewQAbstractItemViewFromPointer(cthis)
}
func NewQListViewFromPointer(cthis unsafe.Pointer) *QListView {
	bcthis0 := NewQAbstractItemViewFromPointer(cthis)
	return &QListView{bcthis0}
}
func (*QListView) NewFromPointer(cthis unsafe.Pointer) *QListView {
	return NewQListViewFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlistview.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QListView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListView(QWidget *)
func NewQListView(parent *QWidget /*777 QWidget **/) *QListView {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListViewC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQListViewFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlistview.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QListView()
func DeleteQListView(this *QListView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListViewD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlistview.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovement(enum QListView::Movement)
func (this *QListView) SetMovement(movement int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setMovementENS_8MovementE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), movement)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::Movement movement()
func (this *QListView) Movement() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8movementEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlow(enum QListView::Flow)
func (this *QListView) SetFlow(flow int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView7setFlowENS_4FlowE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::Flow flow()
func (this *QListView) Flow() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView4flowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWrapping(_Bool)
func (this *QListView) SetWrapping(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setWrappingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWrapping()
func (this *QListView) IsWrapping() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10isWrappingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResizeMode(enum QListView::ResizeMode)
func (this *QListView) SetResizeMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13setResizeModeENS_10ResizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::ResizeMode resizeMode()
func (this *QListView) ResizeMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10resizeModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutMode(enum QListView::LayoutMode)
func (this *QListView) SetLayoutMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13setLayoutModeENS_10LayoutModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::LayoutMode layoutMode()
func (this *QListView) LayoutMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10layoutModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)
func (this *QListView) SetSpacing(space int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10setSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), space)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing()
func (this *QListView) Spacing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView7spacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBatchSize(int)
func (this *QListView) SetBatchSize(batchSize int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setBatchSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), batchSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int batchSize()
func (this *QListView) BatchSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView9batchSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGridSize(const QSize &)
func (this *QListView) SetGridSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setGridSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize gridSize()
func (this *QListView) GridSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8gridSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewMode(enum QListView::ViewMode)
func (this *QListView) SetViewMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setViewModeENS_8ViewModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QListView::ViewMode viewMode()
func (this *QListView) ViewMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8viewModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearPropertyFlags()
func (this *QListView) ClearPropertyFlags() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView18clearPropertyFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRowHidden(int)
func (this *QListView) IsRowHidden(row int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView11isRowHiddenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowHidden(int, _Bool)
func (this *QListView) SetRowHidden(row int, hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setRowHiddenEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModelColumn(int)
func (this *QListView) SetModelColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14setModelColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int modelColumn()
func (this *QListView) ModelColumn() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView11modelColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUniformItemSizes(_Bool)
func (this *QListView) SetUniformItemSizes(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView19setUniformItemSizesEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool uniformItemSizes()
func (this *QListView) UniformItemSizes() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView16uniformItemSizesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordWrap(_Bool)
func (this *QListView) SetWordWrap(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11setWordWrapEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wordWrap()
func (this *QListView) WordWrap() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView8wordWrapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectionRectVisible(_Bool)
func (this *QListView) SetSelectionRectVisible(show bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView23setSelectionRectVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), show)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSelectionRectVisible()
func (this *QListView) IsSelectionRectVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView22isSelectionRectVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:128
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect visualRect(const QModelIndex &)
func (this *QListView) VisualRect(index *qtcore.QModelIndex) *qtcore.QRect /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void scrollTo(const QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QListView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:130
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex indexAt(const QPoint &)
func (this *QListView) IndexAt(p *qtcore.QPoint) *qtcore.QModelIndex /*123*/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView7indexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void doItemsLayout()
func (this *QListView) DoItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13doItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()
func (this *QListView) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:134
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)
func (this *QListView) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QListView) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:144
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)
func (this *QListView) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:146
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void resizeContents(int, int)
func (this *QListView) ResizeContents(width int, height int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14resizeContentsEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width, height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:147
// index:0
// Protected Visibility=Default Availability=Available
// [8] QSize contentsSize()
func (this *QListView) ContentsSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView12contentsSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:150
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsInserted(const QModelIndex &, int, int)
func (this *QListView) RowsInserted(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:151
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void rowsAboutToBeRemoved(const QModelIndex &, int, int)
func (this *QListView) RowsAboutToBeRemoved(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView20rowsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, start, end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:153
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QListView) MouseMoveEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:154
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QListView) MouseReleaseEvent(e *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:156
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QListView) WheelEvent(e *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:159
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QListView) TimerEvent(e *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:160
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QListView) ResizeEvent(e *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QListView) DragMoveEvent(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:163
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)
func (this *QListView) DragLeaveEvent(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:164
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QListView) DropEvent(e *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:168
// index:0
// Protected virtual Visibility=Default Availability=Available
// [192] QStyleOptionViewItem viewOptions()
func (this *QListView) ViewOptions() *QStyleOptionViewItem /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView11viewOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStyleOptionViewItem)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:169
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QListView) PaintEvent(e *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:171
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int horizontalOffset()
func (this *QListView) HorizontalOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView16horizontalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:172
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int verticalOffset()
func (this *QListView) VerticalOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView14verticalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistview.h:174
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect rectForIndex(const QModelIndex &)
func (this *QListView) RectForIndex(index *qtcore.QModelIndex) *qtcore.QRect /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView12rectForIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:175
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setPositionForIndex(const QPoint &, const QModelIndex &)
func (this *QListView) SetPositionForIndex(position *qtcore.QPoint, index *qtcore.QModelIndex) {
	var convArg0 = position.GetCthis()
	var convArg1 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView19setPositionForIndexERK6QPointRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:177
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setSelection(const QRect &, QItemSelectionModel::SelectionFlags)
func (this *QListView) SetSelection(rect *qtcore.QRect, command int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:178
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QRegion visualRegionForSelection(const QItemSelection &)
func (this *QListView) VisualRegionForSelection(selection *qtcore.QItemSelection) *qtgui.QRegion /*123*/ {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistview.h:181
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateGeometries()
func (this *QListView) UpdateGeometries() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView16updateGeometriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool isIndexHidden(const QModelIndex &)
func (this *QListView) IsIndexHidden(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistview.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QListView) SelectionChanged(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) {
	var convArg0 = selected.GetCthis()
	var convArg1 = deselected.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:186
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QListView) CurrentChanged(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QListView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistview.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint()
func (this *QListView) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QListView16viewportSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

type QListView__Movement = int

const QListView__Static QListView__Movement = 0
const QListView__Free QListView__Movement = 1
const QListView__Snap QListView__Movement = 2

type QListView__Flow = int

const QListView__LeftToRight QListView__Flow = 0
const QListView__TopToBottom QListView__Flow = 1

type QListView__ResizeMode = int

const QListView__Fixed QListView__ResizeMode = 0
const QListView__Adjust QListView__ResizeMode = 1

type QListView__LayoutMode = int

const QListView__SinglePass QListView__LayoutMode = 0
const QListView__Batched QListView__LayoutMode = 1

type QListView__ViewMode = int

const QListView__ListMode QListView__ViewMode = 0
const QListView__IconMode QListView__ViewMode = 1

//  body block end
