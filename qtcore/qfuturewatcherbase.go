package qtcore

// /usr/include/qt/QtCore/qfuturewatcher.h
// #include <qfuturewatcher.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
// void connectNotify(const class QMetaMethod &)
func (this *QFutureWatcherBase) InheritConnectNotify(f func(signal *QMetaMethod)) {
	ffiqt.SetAllInheritCallback(this, "connectNotify", f)
}

// void disconnectNotify(const class QMetaMethod &)
func (this *QFutureWatcherBase) InheritDisconnectNotify(f func(signal *QMetaMethod)) {
	ffiqt.SetAllInheritCallback(this, "disconnectNotify", f)
}

// void connectOutputInterface()
func (this *QFutureWatcherBase) InheritConnectOutputInterface(f func()) {
	ffiqt.SetAllInheritCallback(this, "connectOutputInterface", f)
}

// void disconnectOutputInterface(_Bool)
func (this *QFutureWatcherBase) InheritDisconnectOutputInterface(f func(pendingAssignment bool)) {
	ffiqt.SetAllInheritCallback(this, "disconnectOutputInterface", f)
}

type QFutureWatcherBase struct {
	*QObject
}

func (this *QFutureWatcherBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QFutureWatcherBase) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQFutureWatcherBaseFromPointer(cthis unsafe.Pointer) *QFutureWatcherBase {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFutureWatcherBase{bcthis0}
}
func (*QFutureWatcherBase) NewFromPointer(cthis unsafe.Pointer) *QFutureWatcherBase {
	return NewQFutureWatcherBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFutureWatcherBase) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qfuturewatcher.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFutureWatcherBase(QObject *)
func NewQFutureWatcherBase(parent *QObject /*777 QObject **/) *QFutureWatcherBase {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBaseC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFutureWatcherBaseFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfuturewatcher.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressValue()
func (this *QFutureWatcherBase) ProgressValue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase13progressValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfuturewatcher.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressMinimum()
func (this *QFutureWatcherBase) ProgressMinimum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase15progressMinimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfuturewatcher.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressMaximum()
func (this *QFutureWatcherBase) ProgressMaximum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase15progressMaximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfuturewatcher.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString progressText()
func (this *QFutureWatcherBase) ProgressText() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase12progressTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qfuturewatcher.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStarted()
func (this *QFutureWatcherBase) IsStarted() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase9isStartedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished()
func (this *QFutureWatcherBase) IsFinished() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10isFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning()
func (this *QFutureWatcherBase) IsRunning() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase9isRunningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCanceled()
func (this *QFutureWatcherBase) IsCanceled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10isCanceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:73
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPaused()
func (this *QFutureWatcherBase) IsPaused() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase8isPausedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void waitForFinished()
func (this *QFutureWatcherBase) WaitForFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase15waitForFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPendingResultsLimit(int)
func (this *QFutureWatcherBase) SetPendingResultsLimit(limit int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase22setPendingResultsLimitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), limit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QFutureWatcherBase) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfuturewatcher.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void started()
func (this *QFutureWatcherBase) Started() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase7startedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished()
func (this *QFutureWatcherBase) Finished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase8finishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canceled()
func (this *QFutureWatcherBase) Canceled() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase8canceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paused()
func (this *QFutureWatcherBase) Paused() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6pausedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resumed()
func (this *QFutureWatcherBase) Resumed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase7resumedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resultReadyAt(int)
func (this *QFutureWatcherBase) ResultReadyAt(resultIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase13resultReadyAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), resultIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resultsReadyAt(int, int)
func (this *QFutureWatcherBase) ResultsReadyAt(beginIndex int, endIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase14resultsReadyAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), beginIndex, endIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void progressRangeChanged(int, int)
func (this *QFutureWatcherBase) ProgressRangeChanged(minimum int, maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase20progressRangeChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void progressValueChanged(int)
func (this *QFutureWatcherBase) ProgressValueChanged(progressValue int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase20progressValueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), progressValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void progressTextChanged(const QString &)
func (this *QFutureWatcherBase) ProgressTextChanged(progressText *QString) {
	var convArg0 = progressText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase19progressTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel()
func (this *QFutureWatcherBase) Cancel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6cancelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(_Bool)
func (this *QFutureWatcherBase) SetPaused(paused bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase9setPausedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pause()
func (this *QFutureWatcherBase) Pause() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase5pauseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()
func (this *QFutureWatcherBase) Resume() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6resumeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void togglePaused()
func (this *QFutureWatcherBase) TogglePaused() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase12togglePausedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:101
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void connectNotify(const QMetaMethod &)
func (this *QFutureWatcherBase) ConnectNotify(signal *QMetaMethod) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:102
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void disconnectNotify(const QMetaMethod &)
func (this *QFutureWatcherBase) DisconnectNotify(signal *QMetaMethod) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void connectOutputInterface()
func (this *QFutureWatcherBase) ConnectOutputInterface() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase22connectOutputInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void disconnectOutputInterface(_Bool)
func (this *QFutureWatcherBase) DisconnectOutputInterface(pendingAssignment bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pendingAssignment)
	gopp.ErrPrint(err, rv)
}

func DeleteQFutureWatcherBase(this *QFutureWatcherBase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBaseD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
