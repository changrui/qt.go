package qtwidgets

// /usr/include/qt/QtWidgets/qspinbox.h
// #include <qspinbox.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
func (this *QSpinBox) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// QValidator::State validate(class QString &, int &)
func (this *QSpinBox) InheritValidate(f func(input *qtcore.QString, pos int) int) {
	ffiqt.SetAllInheritCallback(this, "validate", f)
}

// int valueFromText(const class QString &)
func (this *QSpinBox) InheritValueFromText(f func(text *qtcore.QString) int) {
	ffiqt.SetAllInheritCallback(this, "valueFromText", f)
}

// QString textFromValue(int)
func (this *QSpinBox) InheritTextFromValue(f func(val int) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "textFromValue", f)
}

// void fixup(class QString &)
func (this *QSpinBox) InheritFixup(f func(str *qtcore.QString)) {
	ffiqt.SetAllInheritCallback(this, "fixup", f)
}

type QSpinBox struct {
	*QAbstractSpinBox
}

func (this *QSpinBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSpinBox.GetCthis()
	}
}
func (this *QSpinBox) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSpinBox = NewQAbstractSpinBoxFromPointer(cthis)
}
func NewQSpinBoxFromPointer(cthis unsafe.Pointer) *QSpinBox {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QSpinBox{bcthis0}
}
func (*QSpinBox) NewFromPointer(cthis unsafe.Pointer) *QSpinBox {
	return NewQSpinBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qspinbox.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSpinBox) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSpinBox(QWidget *)
func NewQSpinBox(parent *QWidget /*777 QWidget **/) *QSpinBox {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBoxC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qspinbox.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSpinBox()
func DeleteQSpinBox(this *QSpinBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBoxD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qspinbox.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int value()
func (this *QSpinBox) Value() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString prefix()
func (this *QSpinBox) Prefix() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox6prefixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPrefix(const QString &)
func (this *QSpinBox) SetPrefix(prefix *qtcore.QString) {
	var convArg0 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox9setPrefixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QString suffix()
func (this *QSpinBox) Suffix() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox6suffixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSuffix(const QString &)
func (this *QSpinBox) SetSuffix(suffix *qtcore.QString) {
	var convArg0 = suffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox9setSuffixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString cleanText()
func (this *QSpinBox) CleanText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox9cleanTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int singleStep()
func (this *QSpinBox) SingleStep() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox10singleStepEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleStep(int)
func (this *QSpinBox) SetSingleStep(val int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox13setSingleStepEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimum()
func (this *QSpinBox) Minimum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox7minimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(int)
func (this *QSpinBox) SetMinimum(min int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox10setMinimumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), min)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximum()
func (this *QSpinBox) Maximum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox7maximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(int)
func (this *QSpinBox) SetMaximum(max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox10setMaximumEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(int, int)
func (this *QSpinBox) SetRange(min int, max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox8setRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), min, max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int displayIntegerBase()
func (this *QSpinBox) DisplayIntegerBase() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox18displayIntegerBaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDisplayIntegerBase(int)
func (this *QSpinBox) SetDisplayIntegerBase(base int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox21setDisplayIntegerBaseEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), base)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:93
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSpinBox) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qspinbox.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &)
func (this *QSpinBox) Validate(input *qtcore.QString, pos int) int {
	var convArg0 = input.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox8validateER7QStringRi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int valueFromText(const QString &)
func (this *QSpinBox) ValueFromText(text *qtcore.QString) int {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox13valueFromTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qspinbox.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QString textFromValue(int)
func (this *QSpinBox) TextFromValue(val int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox13textFromValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), val)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qspinbox.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &)
func (this *QSpinBox) Fixup(str *qtcore.QString) {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSpinBox5fixupER7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(int)
func (this *QSpinBox) SetValue(val int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox8setValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), val)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(int)
func (this *QSpinBox) ValueChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox12valueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qspinbox.h:105
// index:1
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(const QString &)
func (this *QSpinBox) ValueChanged_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSpinBox12valueChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
