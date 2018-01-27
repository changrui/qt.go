package qtquick

// /usr/include/qt/QtQuick/qsgrendernode.h
// #include <qsgrendernode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSGRenderNode struct {
	*QSGNode
}

func (this *QSGRenderNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGNode.GetCthis()
	}
}
func (this *QSGRenderNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGNode = NewQSGNodeFromPointer(cthis)
}
func NewQSGRenderNodeFromPointer(cthis unsafe.Pointer) *QSGRenderNode {
	bcthis0 := NewQSGNodeFromPointer(cthis)
	return &QSGRenderNode{bcthis0}
}
func (*QSGRenderNode) NewFromPointer(cthis unsafe.Pointer) *QSGRenderNode {
	return NewQSGRenderNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:82
// index:0
// Public
// void QSGRenderNode()
func NewQSGRenderNode() *QSGRenderNode {
	cthis := qtrt.Calloc(1, 256) // 88
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSGRenderNodeC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGRenderNodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgrendernode.h:83
// index:0
// Public virtual
// void ~QSGRenderNode()
func DeleteQSGRenderNode(*QSGRenderNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSGRenderNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:85
// index:0
// Public virtual
// QSGRenderNode::StateFlags changedStates()
func (this *QSGRenderNode) ChangedStates() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSGRenderNode13changedStatesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:87
// index:0
// Public virtual
// void releaseResources()
func (this *QSGRenderNode) ReleaseResources() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QSGRenderNode16releaseResourcesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:88
// index:0
// Public virtual
// QSGRenderNode::RenderingFlags flags()
func (this *QSGRenderNode) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSGRenderNode5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgrendernode.h:89
// index:0
// Public virtual
// QRectF rect()
func (this *QSGRenderNode) Rect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSGRenderNode4rectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgrendernode.h:91
// index:0
// Public
// const QMatrix4x4 * matrix()
func (this *QSGRenderNode) Matrix() *qtgui.QMatrix4x4 /*777 const QMatrix4x4 **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSGRenderNode6matrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgrendernode.h:92
// index:0
// Public
// const QSGClipNode * clipList()
func (this *QSGRenderNode) ClipList() *QSGClipNode /*777 const QSGClipNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSGRenderNode8clipListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGClipNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgrendernode.h:93
// index:0
// Public
// qreal inheritedOpacity()
func (this *QSGRenderNode) InheritedOpacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QSGRenderNode16inheritedOpacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

type QSGRenderNode__StateFlag = int

const QSGRenderNode__DepthState QSGRenderNode__StateFlag = 1
const QSGRenderNode__StencilState QSGRenderNode__StateFlag = 2
const QSGRenderNode__ScissorState QSGRenderNode__StateFlag = 4
const QSGRenderNode__ColorState QSGRenderNode__StateFlag = 8
const QSGRenderNode__BlendState QSGRenderNode__StateFlag = 16
const QSGRenderNode__CullState QSGRenderNode__StateFlag = 32
const QSGRenderNode__ViewportState QSGRenderNode__StateFlag = 64
const QSGRenderNode__RenderTargetState QSGRenderNode__StateFlag = 128

type QSGRenderNode__RenderingFlag = int

const QSGRenderNode__BoundedRectRendering QSGRenderNode__RenderingFlag = 1
const QSGRenderNode__DepthAwareRendering QSGRenderNode__RenderingFlag = 2
const QSGRenderNode__OpaqueRendering QSGRenderNode__RenderingFlag = 4

//  body block end