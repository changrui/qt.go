package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

type QGraphicsItemGroup struct {
	*QGraphicsItem
}

func (this *QGraphicsItemGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsItem.GetCthis()
	}
}
func (this *QGraphicsItemGroup) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsItem = NewQGraphicsItemFromPointer(cthis)
}
func NewQGraphicsItemGroupFromPointer(cthis unsafe.Pointer) *QGraphicsItemGroup {
	bcthis0 := NewQGraphicsItemFromPointer(cthis)
	return &QGraphicsItemGroup{bcthis0}
}
func (*QGraphicsItemGroup) NewFromPointer(cthis unsafe.Pointer) *QGraphicsItemGroup {
	return NewQGraphicsItemGroupFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1004
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsItemGroup(QGraphicsItem *)
func NewQGraphicsItemGroup(parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsItemGroup {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroupC2EP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsItemGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsItemGroup)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1005
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsItemGroup()
func DeleteQGraphicsItemGroup(this *QGraphicsItemGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroupD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1007
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addToGroup(QGraphicsItem *)
func (this *QGraphicsItemGroup) AddToGroup(item *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1008
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeFromGroup(QGraphicsItem *)
func (this *QGraphicsItemGroup) RemoveFromGroup(item *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1010
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsItemGroup) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1011
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsItemGroup) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionGraphicsItem /*777 const QStyleOptionGraphicsItem **/, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1013
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
func (this *QGraphicsItemGroup) IsObscuredBy(item *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1014
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea()
func (this *QGraphicsItemGroup) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1017
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsItemGroup) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QGraphicsItemGroup__ = int

const QGraphicsItemGroup__Type QGraphicsItemGroup__ = 10

//  body block end
