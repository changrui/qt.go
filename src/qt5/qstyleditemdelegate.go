package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qstyleditemdelegate.h
// dst-file: /src/widgets/qstyleditemdelegate.go
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
  // proto:  QSize QStyledItemDelegate::sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1);
  // proto:  QWidget * QStyledItemDelegate::createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QStyledItemDelegate::~QStyledItemDelegate();
extern void _ZN19QStyledItemDelegateD0Ev(void* qthis);
  // proto:  void QStyledItemDelegate::updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QStyledItemDelegate::setEditorData(QWidget * editor, const QModelIndex & index);
extern void _ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(void* qthis, void* arg0, void* arg1);
  // proto:  void QStyledItemDelegate::setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index);
extern void _ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QStyledItemDelegate::QStyledItemDelegate(const QStyledItemDelegate & );
extern void* dector_ZN19QStyledItemDelegateC1ERKS_(void* arg0);
extern void _ZN19QStyledItemDelegateC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyledItemDelegate::setItemEditorFactory(QItemEditorFactory * factory);
extern void _ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory(void* qthis, void* arg0);
  // proto:  void QStyledItemDelegate::paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  QItemEditorFactory * QStyledItemDelegate::itemEditorFactory();
extern void _ZNK19QStyledItemDelegate17itemEditorFactoryEv(void* qthis);
  // proto:  void QStyledItemDelegate::QStyledItemDelegate(QObject * parent);
extern void* dector_ZN19QStyledItemDelegateC1EP7QObject(void* arg0);
extern void _ZN19QStyledItemDelegateC1EP7QObject(void* qthis, void* arg0);
  // proto:  QString QStyledItemDelegate::displayText(const QVariant & value, const QLocale & locale);
extern void _ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale(void* qthis, void* arg0, void* arg1);
  // proto:  const QMetaObject * QStyledItemDelegate::metaObject();
extern void _ZNK19QStyledItemDelegate10metaObjectEv(void* qthis);
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

// class sizeof(QStyledItemDelegate)=1
type QStyledItemDelegate struct {
  /*qbase*/ QAbstractItemDelegate;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QSize QStyledItemDelegate::sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index);
func (this *QStyledItemDelegate) sizeHint(args ...interface{}) () {
  // sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex
    // invoke: QSize sizeHint(const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK19QStyledItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "sizeHint", args)
  }

}

  // proto:  QWidget * QStyledItemDelegate::createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index);
func (this *QStyledItemDelegate) createEditor(args ...interface{}) () {
  // createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK19QStyledItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "createEditor", args)
  }

}

  // proto:  void QStyledItemDelegate::~QStyledItemDelegate();
func (this *QStyledItemDelegate) FreeQStyledItemDelegate(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "~QStyledItemDelegate", args)
  }

}

  // proto:  void QStyledItemDelegate::updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index);
func (this *QStyledItemDelegate) updateEditorGeometry(args ...interface{}) () {
  // updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK19QStyledItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "updateEditorGeometry", args)
  }

}

  // proto:  void QStyledItemDelegate::setEditorData(QWidget * editor, const QModelIndex & index);
func (this *QStyledItemDelegate) setEditorData(args ...interface{}) () {
  // setEditorData(class QWidget *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex
    // invoke: void setEditorData(class QWidget *, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK19QStyledItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "setEditorData", args)
  }

}

  // proto:  void QStyledItemDelegate::setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index);
func (this *QStyledItemDelegate) setModelData(args ...interface{}) () {
  // setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex
    // invoke: void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK19QStyledItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "setModelData", args)
  }

}

  // proto:  void QStyledItemDelegate::QStyledItemDelegate(const QStyledItemDelegate & );
func NewQStyledItemDelegate(args ...interface{}) QStyledItemDelegate {
  return QStyledItemDelegate{}
}

  // proto:  void QStyledItemDelegate::setItemEditorFactory(QItemEditorFactory * factory);
func (this *QStyledItemDelegate) setItemEditorFactory(args ...interface{}) () {
  // setItemEditorFactory(class QItemEditorFactory *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemEditorFactory{}) // "QItemEditorFactory *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory
    // invoke: void setItemEditorFactory(class QItemEditorFactory *)
    var arg0 = args[0].(QItemEditorFactory).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QStyledItemDelegate20setItemEditorFactoryEP18QItemEditorFactory(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "setItemEditorFactory", args)
  }

}

  // proto:  void QStyledItemDelegate::paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index);
func (this *QStyledItemDelegate) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: void paint(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK19QStyledItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "paint", args)
  }

}

  // proto:  QItemEditorFactory * QStyledItemDelegate::itemEditorFactory();
func (this *QStyledItemDelegate) itemEditorFactory(args ...interface{}) () {
  // itemEditorFactory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate17itemEditorFactoryEv
    // invoke: QItemEditorFactory * itemEditorFactory()
    C._ZNK19QStyledItemDelegate17itemEditorFactoryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "itemEditorFactory", args)
  }

}

  // proto:  QString QStyledItemDelegate::displayText(const QVariant & value, const QLocale & locale);
func (this *QStyledItemDelegate) displayText(args ...interface{}) () {
  // displayText(const class QVariant &, const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][1] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale
    // invoke: QString displayText(const class QVariant &, const class QLocale &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QLocale).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK19QStyledItemDelegate11displayTextERK8QVariantRK7QLocale(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "displayText", args)
  }

}

  // proto:  const QMetaObject * QStyledItemDelegate::metaObject();
func (this *QStyledItemDelegate) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QStyledItemDelegate10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK19QStyledItemDelegate10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyledItemDelegate", "metaObject", args)
  }

}

// <= body block end

