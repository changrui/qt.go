package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 100
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QVariantComparisonHelper struct {
	*qtrt.CObject
}

func (this *QVariantComparisonHelper) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVariantComparisonHelper) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVariantComparisonHelperFromPointer(cthis unsafe.Pointer) *QVariantComparisonHelper {
	return &QVariantComparisonHelper{&qtrt.CObject{cthis}}
}
func (*QVariantComparisonHelper) NewFromPointer(cthis unsafe.Pointer) *QVariantComparisonHelper {
	return NewQVariantComparisonHelperFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:560
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVariantComparisonHelper(const QVariant &)
func NewQVariantComparisonHelper(var_ *QVariant) *QVariantComparisonHelper {
	var convArg0 = var_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QVariantComparisonHelperC2ERK8QVariant", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantComparisonHelperFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariantComparisonHelper)
	return gothis
}

func DeleteQVariantComparisonHelper(this *QVariantComparisonHelper) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QVariantComparisonHelperD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
