package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qopenglpixeltransferoptions.h
// dst-file: /src/gui/qopenglpixeltransferoptions.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QOpenGLPixelTransferOptions::QOpenGLPixelTransferOptions();
extern void* dector_ZN27QOpenGLPixelTransferOptionsC1Ev();
extern void _ZN27QOpenGLPixelTransferOptionsC1Ev(void* qthis);
  // proto:  void QOpenGLPixelTransferOptions::~QOpenGLPixelTransferOptions();
extern void _ZN27QOpenGLPixelTransferOptionsD0Ev(void* qthis);
  // proto:  bool QOpenGLPixelTransferOptions::isSwapBytesEnabled();
extern void _ZNK27QOpenGLPixelTransferOptions18isSwapBytesEnabledEv(void* qthis);
  // proto:  void QOpenGLPixelTransferOptions::swap(QOpenGLPixelTransferOptions & other);
extern void _ZN27QOpenGLPixelTransferOptions4swapERS_(void* qthis, void* arg0);
  // proto:  void QOpenGLPixelTransferOptions::QOpenGLPixelTransferOptions(const QOpenGLPixelTransferOptions & );
extern void* dector_ZN27QOpenGLPixelTransferOptionsC1ERKS_(void* arg0);
extern void _ZN27QOpenGLPixelTransferOptionsC1ERKS_(void* qthis, void* arg0);
  // proto:  int QOpenGLPixelTransferOptions::skipImages();
extern void _ZNK27QOpenGLPixelTransferOptions10skipImagesEv(void* qthis);
  // proto:  void QOpenGLPixelTransferOptions::setSkipRows(int skipRows);
extern void _ZN27QOpenGLPixelTransferOptions11setSkipRowsEi(void* qthis, int arg0);
  // proto:  int QOpenGLPixelTransferOptions::skipPixels();
extern void _ZNK27QOpenGLPixelTransferOptions10skipPixelsEv(void* qthis);
  // proto:  void QOpenGLPixelTransferOptions::setRowLength(int rowLength);
extern void _ZN27QOpenGLPixelTransferOptions12setRowLengthEi(void* qthis, int arg0);
  // proto:  int QOpenGLPixelTransferOptions::imageHeight();
extern void _ZNK27QOpenGLPixelTransferOptions11imageHeightEv(void* qthis);
  // proto:  void QOpenGLPixelTransferOptions::setImageHeight(int imageHeight);
extern void _ZN27QOpenGLPixelTransferOptions14setImageHeightEi(void* qthis, int arg0);
  // proto:  int QOpenGLPixelTransferOptions::skipRows();
extern void _ZNK27QOpenGLPixelTransferOptions8skipRowsEv(void* qthis);
  // proto:  void QOpenGLPixelTransferOptions::setAlignment(int alignment);
extern void _ZN27QOpenGLPixelTransferOptions12setAlignmentEi(void* qthis, int arg0);
  // proto:  void QOpenGLPixelTransferOptions::setSkipImages(int skipImages);
extern void _ZN27QOpenGLPixelTransferOptions13setSkipImagesEi(void* qthis, int arg0);
  // proto:  int QOpenGLPixelTransferOptions::alignment();
extern void _ZNK27QOpenGLPixelTransferOptions9alignmentEv(void* qthis);
  // proto:  void QOpenGLPixelTransferOptions::setSkipPixels(int skipPixels);
extern void _ZN27QOpenGLPixelTransferOptions13setSkipPixelsEi(void* qthis, int arg0);
  // proto:  void QOpenGLPixelTransferOptions::setSwapBytesEnabled(bool swapBytes);
extern void _ZN27QOpenGLPixelTransferOptions19setSwapBytesEnabledEb(void* qthis, bool arg0);
  // proto:  void QOpenGLPixelTransferOptions::setLeastSignificantByteFirst(bool lsbFirst);
extern void _ZN27QOpenGLPixelTransferOptions28setLeastSignificantByteFirstEb(void* qthis, bool arg0);
  // proto:  bool QOpenGLPixelTransferOptions::isLeastSignificantBitFirst();
extern void _ZNK27QOpenGLPixelTransferOptions26isLeastSignificantBitFirstEv(void* qthis);
  // proto:  int QOpenGLPixelTransferOptions::rowLength();
extern void _ZNK27QOpenGLPixelTransferOptions9rowLengthEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QOpenGLPixelTransferOptions)=1
type QOpenGLPixelTransferOptions struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QOpenGLPixelTransferOptions::QOpenGLPixelTransferOptions();
func NewQOpenGLPixelTransferOptions(args ...interface{}) QOpenGLPixelTransferOptions {
  return QOpenGLPixelTransferOptions{}
}

  // proto:  void QOpenGLPixelTransferOptions::~QOpenGLPixelTransferOptions();
func (this *QOpenGLPixelTransferOptions) FreeQOpenGLPixelTransferOptions(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "~QOpenGLPixelTransferOptions", args)
  }

}

  // proto:  bool QOpenGLPixelTransferOptions::isSwapBytesEnabled();
func (this *QOpenGLPixelTransferOptions) isSwapBytesEnabled(args ...interface{}) () {
  // isSwapBytesEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions18isSwapBytesEnabledEv
    // invoke: bool isSwapBytesEnabled()
    C._ZNK27QOpenGLPixelTransferOptions18isSwapBytesEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "isSwapBytesEnabled", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::swap(QOpenGLPixelTransferOptions & other);
func (this *QOpenGLPixelTransferOptions) swap(args ...interface{}) () {
  // swap(class QOpenGLPixelTransferOptions &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLPixelTransferOptions{}) // "QOpenGLPixelTransferOptions &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions4swapERS_
    // invoke: void swap(class QOpenGLPixelTransferOptions &)
    var arg0 = args[0].(QOpenGLPixelTransferOptions).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "swap", args)
  }

}

  // proto:  int QOpenGLPixelTransferOptions::skipImages();
func (this *QOpenGLPixelTransferOptions) skipImages(args ...interface{}) () {
  // skipImages()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions10skipImagesEv
    // invoke: int skipImages()
    C._ZNK27QOpenGLPixelTransferOptions10skipImagesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipImages", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::setSkipRows(int skipRows);
func (this *QOpenGLPixelTransferOptions) setSkipRows(args ...interface{}) () {
  // setSkipRows(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions11setSkipRowsEi
    // invoke: void setSkipRows(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions11setSkipRowsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipRows", args)
  }

}

  // proto:  int QOpenGLPixelTransferOptions::skipPixels();
func (this *QOpenGLPixelTransferOptions) skipPixels(args ...interface{}) () {
  // skipPixels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions10skipPixelsEv
    // invoke: int skipPixels()
    C._ZNK27QOpenGLPixelTransferOptions10skipPixelsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipPixels", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::setRowLength(int rowLength);
func (this *QOpenGLPixelTransferOptions) setRowLength(args ...interface{}) () {
  // setRowLength(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions12setRowLengthEi
    // invoke: void setRowLength(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions12setRowLengthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setRowLength", args)
  }

}

  // proto:  int QOpenGLPixelTransferOptions::imageHeight();
func (this *QOpenGLPixelTransferOptions) imageHeight(args ...interface{}) () {
  // imageHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions11imageHeightEv
    // invoke: int imageHeight()
    C._ZNK27QOpenGLPixelTransferOptions11imageHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "imageHeight", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::setImageHeight(int imageHeight);
func (this *QOpenGLPixelTransferOptions) setImageHeight(args ...interface{}) () {
  // setImageHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions14setImageHeightEi
    // invoke: void setImageHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions14setImageHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setImageHeight", args)
  }

}

  // proto:  int QOpenGLPixelTransferOptions::skipRows();
func (this *QOpenGLPixelTransferOptions) skipRows(args ...interface{}) () {
  // skipRows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions8skipRowsEv
    // invoke: int skipRows()
    C._ZNK27QOpenGLPixelTransferOptions8skipRowsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "skipRows", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::setAlignment(int alignment);
func (this *QOpenGLPixelTransferOptions) setAlignment(args ...interface{}) () {
  // setAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions12setAlignmentEi
    // invoke: void setAlignment(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions12setAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setAlignment", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::setSkipImages(int skipImages);
func (this *QOpenGLPixelTransferOptions) setSkipImages(args ...interface{}) () {
  // setSkipImages(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions13setSkipImagesEi
    // invoke: void setSkipImages(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions13setSkipImagesEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipImages", args)
  }

}

  // proto:  int QOpenGLPixelTransferOptions::alignment();
func (this *QOpenGLPixelTransferOptions) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions9alignmentEv
    // invoke: int alignment()
    C._ZNK27QOpenGLPixelTransferOptions9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "alignment", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::setSkipPixels(int skipPixels);
func (this *QOpenGLPixelTransferOptions) setSkipPixels(args ...interface{}) () {
  // setSkipPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions13setSkipPixelsEi
    // invoke: void setSkipPixels(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions13setSkipPixelsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSkipPixels", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::setSwapBytesEnabled(bool swapBytes);
func (this *QOpenGLPixelTransferOptions) setSwapBytesEnabled(args ...interface{}) () {
  // setSwapBytesEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions19setSwapBytesEnabledEb
    // invoke: void setSwapBytesEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions19setSwapBytesEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setSwapBytesEnabled", args)
  }

}

  // proto:  void QOpenGLPixelTransferOptions::setLeastSignificantByteFirst(bool lsbFirst);
func (this *QOpenGLPixelTransferOptions) setLeastSignificantByteFirst(args ...interface{}) () {
  // setLeastSignificantByteFirst(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QOpenGLPixelTransferOptions28setLeastSignificantByteFirstEb
    // invoke: void setLeastSignificantByteFirst(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN27QOpenGLPixelTransferOptions28setLeastSignificantByteFirstEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "setLeastSignificantByteFirst", args)
  }

}

  // proto:  bool QOpenGLPixelTransferOptions::isLeastSignificantBitFirst();
func (this *QOpenGLPixelTransferOptions) isLeastSignificantBitFirst(args ...interface{}) () {
  // isLeastSignificantBitFirst()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions26isLeastSignificantBitFirstEv
    // invoke: bool isLeastSignificantBitFirst()
    C._ZNK27QOpenGLPixelTransferOptions26isLeastSignificantBitFirstEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "isLeastSignificantBitFirst", args)
  }

}

  // proto:  int QOpenGLPixelTransferOptions::rowLength();
func (this *QOpenGLPixelTransferOptions) rowLength(args ...interface{}) () {
  // rowLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QOpenGLPixelTransferOptions9rowLengthEv
    // invoke: int rowLength()
    C._ZNK27QOpenGLPixelTransferOptions9rowLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLPixelTransferOptions", "rowLength", args)
  }

}

// <= body block end

