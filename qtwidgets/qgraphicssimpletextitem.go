package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsSimpleTextItem) InheritSupportsExtension(f func(extension int) bool) {
	ffiqt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsSimpleTextItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant)) {
	ffiqt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const class QVariant &)
func (this *QGraphicsSimpleTextItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "extension", f)
}

type QGraphicsSimpleTextItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsSimpleTextItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractGraphicsShapeItem.GetCthis()
	}
}
func (this *QGraphicsSimpleTextItem) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractGraphicsShapeItem = NewQAbstractGraphicsShapeItemFromPointer(cthis)
}
func NewQGraphicsSimpleTextItemFromPointer(cthis unsafe.Pointer) *QGraphicsSimpleTextItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsSimpleTextItem{bcthis0}
}
func (*QGraphicsSimpleTextItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSimpleTextItem {
	return NewQGraphicsSimpleTextItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:968
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSimpleTextItem(QGraphicsItem *)
func NewQGraphicsSimpleTextItem(parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsSimpleTextItem {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSimpleTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSimpleTextItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:969
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSimpleTextItem(const QString &, QGraphicsItem *)
func NewQGraphicsSimpleTextItem_1(text *qtcore.QString, parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsSimpleTextItem {
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemC2ERK7QStringP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSimpleTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSimpleTextItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:970
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSimpleTextItem()
func DeleteQGraphicsSimpleTextItem(this *QGraphicsSimpleTextItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:972
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QGraphicsSimpleTextItem) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:973
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QGraphicsSimpleTextItem) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:975
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QGraphicsSimpleTextItem) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:976
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QGraphicsSimpleTextItem) Font() *qtgui.QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:978
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsSimpleTextItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:979
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape()
func (this *QGraphicsSimpleTextItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:980
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &)
func (this *QGraphicsSimpleTextItem) Contains(point *qtcore.QPointF) bool {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:982
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsSimpleTextItem) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionGraphicsItem /*777 const QStyleOptionGraphicsItem **/, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:984
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
func (this *QGraphicsSimpleTextItem) IsObscuredBy(item *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:985
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea()
func (this *QGraphicsSimpleTextItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:988
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsSimpleTextItem) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:991
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsSimpleTextItem) SupportsExtension(extension int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:992
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(enum QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsSimpleTextItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:993
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &)
func (this *QGraphicsSimpleTextItem) Extension(variant *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

type QGraphicsSimpleTextItem__ = int

const QGraphicsSimpleTextItem__Type QGraphicsSimpleTextItem__ = 9

//  body block end
