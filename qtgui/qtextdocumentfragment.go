package qtgui

// /usr/include/qt/QtGui/qtextdocumentfragment.h
// #include <qtextdocumentfragment.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QTextDocumentFragment struct {
	*qtrt.CObject
}

func (this *QTextDocumentFragment) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextDocumentFragment) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextDocumentFragmentFromPointer(cthis unsafe.Pointer) *QTextDocumentFragment {
	return &QTextDocumentFragment{&qtrt.CObject{cthis}}
}
func (*QTextDocumentFragment) NewFromPointer(cthis unsafe.Pointer) *QTextDocumentFragment {
	return NewQTextDocumentFragmentFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentFragment()
func NewQTextDocumentFragment() *QTextDocumentFragment {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentFragment)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentFragment(const QTextDocument *)
func NewQTextDocumentFragment_1(document *QTextDocument /*777 const QTextDocument **/) *QTextDocumentFragment {
	var convArg0 = document.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2EPK13QTextDocument", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentFragment)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:59
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentFragment(const QTextCursor &)
func NewQTextDocumentFragment_2(range_ *QTextCursor) *QTextDocumentFragment {
	var convArg0 = range_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2ERK11QTextCursor", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentFragment)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextDocumentFragment()
func DeleteQTextDocumentFragment(this *QTextDocumentFragment) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QTextDocumentFragment) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toPlainText()
func (this *QTextDocumentFragment) ToPlainText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment11toPlainTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml(const QByteArray &)
func (this *QTextDocumentFragment) ToHtml(encoding *qtcore.QByteArray) *qtcore.QString /*123*/ {
	var convArg0 = encoding.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment6toHtmlERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:71
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextDocumentFragment fromPlainText(const QString &)
func (this *QTextDocumentFragment) FromPlainText(plainText *qtcore.QString) *QTextDocumentFragment /*123*/ {
	var convArg0 = plainText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment13fromPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextDocumentFragment)
	return rv2
}
func QTextDocumentFragment_FromPlainText(plainText *qtcore.QString) *QTextDocumentFragment /*123*/ {
	var nilthis *QTextDocumentFragment
	rv := nilthis.FromPlainText(plainText)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:73
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextDocumentFragment fromHtml(const QString &)
func (this *QTextDocumentFragment) FromHtml(html *qtcore.QString) *QTextDocumentFragment /*123*/ {
	var convArg0 = html.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment8fromHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextDocumentFragment)
	return rv2
}
func QTextDocumentFragment_FromHtml(html *qtcore.QString) *QTextDocumentFragment /*123*/ {
	var nilthis *QTextDocumentFragment
	rv := nilthis.FromHtml(html)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:74
// index:1
// Public static Visibility=Default Availability=Available
// [8] QTextDocumentFragment fromHtml(const QString &, const QTextDocument *)
func (this *QTextDocumentFragment) FromHtml_1(html *qtcore.QString, resourceProvider *QTextDocument /*777 const QTextDocument **/) *QTextDocumentFragment /*123*/ {
	var convArg0 = html.GetCthis()
	var convArg1 = resourceProvider.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment8fromHtmlERK7QStringPK13QTextDocument", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextDocumentFragment)
	return rv2
}
func QTextDocumentFragment_FromHtml_1(html *qtcore.QString, resourceProvider *QTextDocument /*777 const QTextDocument **/) *QTextDocumentFragment /*123*/ {
	var nilthis *QTextDocumentFragment
	rv := nilthis.FromHtml_1(html, resourceProvider)
	return rv
}

//  body block end
