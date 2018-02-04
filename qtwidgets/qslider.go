package qtwidgets

// /usr/include/qt/QtWidgets/qslider.h
// #include <qslider.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 41
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
// void paintEvent(class QPaintEvent *)
func (this *QSlider) InheritPaintEvent(f func(ev *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QSlider) InheritMousePressEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QSlider) InheritMouseReleaseEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QSlider) InheritMouseMoveEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void initStyleOption(class QStyleOptionSlider *)
func (this *QSlider) InheritInitStyleOption(f func(option *QStyleOptionSlider /*777 QStyleOptionSlider **/)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QSlider struct {
	*QAbstractSlider
}

func (this *QSlider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSlider.GetCthis()
	}
}
func (this *QSlider) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSlider = NewQAbstractSliderFromPointer(cthis)
}
func NewQSliderFromPointer(cthis unsafe.Pointer) *QSlider {
	bcthis0 := NewQAbstractSliderFromPointer(cthis)
	return &QSlider{bcthis0}
}
func (*QSlider) NewFromPointer(cthis unsafe.Pointer) *QSlider {
	return NewQSliderFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qslider.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSlider) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qslider.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSlider(QWidget *)
func NewQSlider(parent *QWidget /*777 QWidget **/) *QSlider {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSliderC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qslider.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSlider(Qt::Orientation, QWidget *)
func NewQSlider_1(orientation int, parent *QWidget /*777 QWidget **/) *QSlider {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSliderC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_POINTER, orientation, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qslider.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSlider()
func DeleteQSlider(this *QSlider) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSliderD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qslider.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QSlider) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qslider.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QSlider) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qslider.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTickPosition(enum QSlider::TickPosition)
func (this *QSlider) SetTickPosition(position int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider15setTickPositionENS_12TickPositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] QSlider::TickPosition tickPosition()
func (this *QSlider) TickPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider12tickPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qslider.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTickInterval(int)
func (this *QSlider) SetTickInterval(ti int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider15setTickIntervalEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ti)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int tickInterval()
func (this *QSlider) TickInterval() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider12tickIntervalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qslider.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSlider) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qslider.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QSlider) PaintEvent(ev *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QSlider) MousePressEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QSlider) MouseReleaseEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QSlider) MouseMoveEvent(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = ev.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSlider14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qslider.h:92
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSlider *)
func (this *QSlider) InitStyleOption(option *QStyleOptionSlider /*777 QStyleOptionSlider **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSlider15initStyleOptionEP18QStyleOptionSlider", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QSlider__TickPosition = int

const QSlider__NoTicks QSlider__TickPosition = 0
const QSlider__TicksAbove QSlider__TickPosition = 1
const QSlider__TicksLeft QSlider__TickPosition = 1
const QSlider__TicksBelow QSlider__TickPosition = 2
const QSlider__TicksRight QSlider__TickPosition = 2
const QSlider__TicksBothSides QSlider__TickPosition = 3

//  body block end
