package qtwidgets

// /usr/include/qt/QtWidgets/qlineedit.h
// #include <qlineedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 107
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
// void mousePressEvent(class QMouseEvent *)
func (this *QLineEdit) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QLineEdit) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QLineEdit) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QLineEdit) InheritMouseDoubleClickEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QLineEdit) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QLineEdit) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QLineEdit) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QLineEdit) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void dragEnterEvent(class QDragEnterEvent *)
func (this *QLineEdit) InheritDragEnterEvent(f func(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QLineEdit) InheritDragMoveEvent(f func(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QLineEdit) InheritDragLeaveEvent(f func(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QLineEdit) InheritDropEvent(f func(arg0 *qtgui.QDropEvent /*777 QDropEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dropEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QLineEdit) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QLineEdit) InheritContextMenuEvent(f func(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QLineEdit) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void initStyleOption(class QStyleOptionFrame *)
func (this *QLineEdit) InheritInitStyleOption(f func(option *QStyleOptionFrame /*777 QStyleOptionFrame **/)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

// QRect cursorRect()
func (this *QLineEdit) InheritCursorRect(f func() unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "cursorRect", f)
}

type QLineEdit struct {
	*QWidget
}

func (this *QLineEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QLineEdit) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQLineEditFromPointer(cthis unsafe.Pointer) *QLineEdit {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QLineEdit{bcthis0}
}
func (*QLineEdit) NewFromPointer(cthis unsafe.Pointer) *QLineEdit {
	return NewQLineEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlineedit.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QLineEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLineEdit(QWidget *)
func NewQLineEdit(parent *QWidget /*777 QWidget **/) *QLineEdit {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEditC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlineedit.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLineEdit(const QString &, QWidget *)
func NewQLineEdit_1(arg0 *qtcore.QString, parent *QWidget /*777 QWidget **/) *QLineEdit {
	var convArg0 = arg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEditC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQLineEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlineedit.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLineEdit()
func DeleteQLineEdit(this *QLineEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEditD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlineedit.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QLineEdit) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayText()
func (this *QLineEdit) DisplayText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit11displayTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QString placeholderText()
func (this *QLineEdit) PlaceholderText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15placeholderTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlaceholderText(const QString &)
func (this *QLineEdit) SetPlaceholderText(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit18setPlaceholderTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxLength()
func (this *QLineEdit) MaxLength() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9maxLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxLength(int)
func (this *QLineEdit) SetMaxLength(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setMaxLengthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrame(_Bool)
func (this *QLineEdit) SetFrame(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit8setFrameEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFrame()
func (this *QLineEdit) HasFrame() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit8hasFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearButtonEnabled(_Bool)
func (this *QLineEdit) SetClearButtonEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit21setClearButtonEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClearButtonEnabled()
func (this *QLineEdit) IsClearButtonEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit20isClearButtonEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QLineEdit::EchoMode echoMode()
func (this *QLineEdit) EchoMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit8echoModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEchoMode(enum QLineEdit::EchoMode)
func (this *QLineEdit) SetEchoMode(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11setEchoModeENS_8EchoModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly()
func (this *QLineEdit) IsReadOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(_Bool)
func (this *QLineEdit) SetReadOnly(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValidator(const QValidator *)
func (this *QLineEdit) SetValidator(arg0 *qtgui.QValidator /*777 const QValidator **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setValidatorEPK10QValidator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] const QValidator * validator()
func (this *QLineEdit) Validator() *qtgui.QValidator /*777 const QValidator **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9validatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQValidatorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompleter(QCompleter *)
func (this *QLineEdit) SetCompleter(completer *QCompleter /*777 QCompleter **/) {
	var convArg0 = completer.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setCompleterEP10QCompleter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QCompleter * completer()
func (this *QLineEdit) Completer() *QCompleter /*777 QCompleter **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9completerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:131
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QLineEdit) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QLineEdit) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorPosition()
func (this *QLineEdit) CursorPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit14cursorPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorPosition(int)
func (this *QLineEdit) SetCursorPosition(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit17setCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorPositionAt(const QPoint &)
func (this *QLineEdit) CursorPositionAt(pos *qtcore.QPoint) int {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit16cursorPositionAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment()
func (this *QLineEdit) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorForward(_Bool, int)
func (this *QLineEdit) CursorForward(mark bool, steps int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit13cursorForwardEbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mark, steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorBackward(_Bool, int)
func (this *QLineEdit) CursorBackward(mark bool, steps int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14cursorBackwardEbi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mark, steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorWordForward(_Bool)
func (this *QLineEdit) CursorWordForward(mark bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit17cursorWordForwardEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mark)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorWordBackward(_Bool)
func (this *QLineEdit) CursorWordBackward(mark bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit18cursorWordBackwardEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mark)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void backspace()
func (this *QLineEdit) Backspace() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9backspaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void del()
func (this *QLineEdit) Del() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit3delEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void home(_Bool)
func (this *QLineEdit) Home(mark bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit4homeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mark)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void end(_Bool)
func (this *QLineEdit) End(mark bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit3endEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mark)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:150
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isModified()
func (this *QLineEdit) IsModified() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit10isModifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModified(_Bool)
func (this *QLineEdit) SetModified(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11setModifiedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelection(int, int)
func (this *QLineEdit) SetSelection(arg0 int, arg1 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setSelectionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:154
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelectedText()
func (this *QLineEdit) HasSelectedText() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15hasSelectedTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedText()
func (this *QLineEdit) SelectedText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit12selectedTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionStart()
func (this *QLineEdit) SelectionStart() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit14selectionStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:157
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionEnd()
func (this *QLineEdit) SelectionEnd() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit12selectionEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:158
// index:0
// Public Visibility=Default Availability=Available
// [4] int selectionLength()
func (this *QLineEdit) SelectionLength() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15selectionLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlineedit.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndoAvailable()
func (this *QLineEdit) IsUndoAvailable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15isUndoAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:161
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRedoAvailable()
func (this *QLineEdit) IsRedoAvailable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15isRedoAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragEnabled(_Bool)
func (this *QLineEdit) SetDragEnabled(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14setDragEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dragEnabled()
func (this *QLineEdit) DragEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit11dragEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QLineEdit) SetCursorMoveStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit18setCursorMoveStyleEN2Qt15CursorMoveStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorMoveStyle cursorMoveStyle()
func (this *QLineEdit) CursorMoveStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15cursorMoveStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QString inputMask()
func (this *QLineEdit) InputMask() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit9inputMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputMask(const QString &)
func (this *QLineEdit) SetInputMask(inputMask *qtcore.QString) {
	var convArg0 = inputMask.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12setInputMaskERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:171
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAcceptableInput()
func (this *QLineEdit) HasAcceptableInput() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit18hasAcceptableInputEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextMargins(int, int, int, int)
func (this *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14setTextMarginsEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:174
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setTextMargins(const QMargins &)
func (this *QLineEdit) SetTextMargins_1(margins *qtcore.QMargins) {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14setTextMarginsERK8QMargins", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getTextMargins(int *, int *, int *, int *)
func (this *QLineEdit) GetTextMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit14getTextMarginsEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:176
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins textMargins()
func (this *QLineEdit) TextMargins() *qtcore.QMargins /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit11textMarginsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAction(QAction *, enum QLineEdit::ActionPosition)
func (this *QLineEdit) AddAction(action *QAction /*777 QAction **/, position int) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9addActionEP7QActionNS_14ActionPositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:181
// index:1
// Public Visibility=Default Availability=Available
// [8] QAction * addAction(const QIcon &, enum QLineEdit::ActionPosition)
func (this *QLineEdit) AddAction_1(icon *qtgui.QIcon, position int) *QAction /*777 QAction **/ {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9addActionERK5QIconNS_14ActionPositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, position)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QLineEdit) SetText(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QLineEdit) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectAll()
func (this *QLineEdit) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo()
func (this *QLineEdit) Undo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit4undoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo()
func (this *QLineEdit) Redo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit4redoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cut()
func (this *QLineEdit) Cut() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit3cutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copy()
func (this *QLineEdit) Copy() {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit4copyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paste()
func (this *QLineEdit) Paste() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit5pasteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void deselect()
func (this *QLineEdit) Deselect() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit8deselectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * createStandardContextMenu()
func (this *QLineEdit) CreateStandardContextMenu() *QMenu /*777 QMenu **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit25createStandardContextMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textChanged(const QString &)
func (this *QLineEdit) TextChanged(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11textChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textEdited(const QString &)
func (this *QLineEdit) TextEdited(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit10textEditedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorPositionChanged(int, int)
func (this *QLineEdit) CursorPositionChanged(arg0 int, arg1 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit21cursorPositionChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void returnPressed()
func (this *QLineEdit) ReturnPressed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit13returnPressedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editingFinished()
func (this *QLineEdit) EditingFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit15editingFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:209
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged()
func (this *QLineEdit) SelectionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit16selectionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:212
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QLineEdit) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:213
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QLineEdit) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QLineEdit) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:215
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QLineEdit) MouseDoubleClickEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:216
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QLineEdit) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:217
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QLineEdit) FocusInEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QLineEdit) FocusOutEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QLineEdit) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)
func (this *QLineEdit) DragEnterEvent(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QLineEdit) DragMoveEvent(e *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:223
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)
func (this *QLineEdit) DragLeaveEvent(e *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:224
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QLineEdit) DropEvent(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:226
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QLineEdit) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:228
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QLineEdit) ContextMenuEvent(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:231
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)
func (this *QLineEdit) InputMethodEvent(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:232
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionFrame *)
func (this *QLineEdit) InitStyleOption(option *QStyleOptionFrame /*777 QStyleOptionFrame **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit15initStyleOptionEP17QStyleOptionFrame", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlineedit.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QLineEdit) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:235
// index:1
// Public Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery, QVariant)
func (this *QLineEdit) InputMethodQuery_1(property int, argument *qtcore.QVariant /*123*/) *qtcore.QVariant /*123*/ {
	var convArg1 = argument.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit16inputMethodQueryEN2Qt16InputMethodQueryE8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), property, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qlineedit.h:236
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QLineEdit) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLineEdit5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlineedit.h:238
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect cursorRect()
func (this *QLineEdit) CursorRect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLineEdit10cursorRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

type QLineEdit__ActionPosition = int

const QLineEdit__LeadingPosition QLineEdit__ActionPosition = 0
const QLineEdit__TrailingPosition QLineEdit__ActionPosition = 1

type QLineEdit__EchoMode = int

const QLineEdit__Normal QLineEdit__EchoMode = 0
const QLineEdit__NoEcho QLineEdit__EchoMode = 1
const QLineEdit__Password QLineEdit__EchoMode = 2
const QLineEdit__PasswordEchoOnEdit QLineEdit__EchoMode = 3

//  body block end
