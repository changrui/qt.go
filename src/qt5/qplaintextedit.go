package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qplaintextedit.h
// dst-file: /src/widgets/qplaintextedit.go
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
  // proto:  void QPlainTextDocumentLayout::requestUpdate();
extern void _ZN24QPlainTextDocumentLayout13requestUpdateEv(void* qthis);
  // proto:  void QPlainTextDocumentLayout::setCursorWidth(int width);
extern void _ZN24QPlainTextDocumentLayout14setCursorWidthEi(void* qthis, int arg0);
  // proto:  QRectF QPlainTextDocumentLayout::frameBoundingRect(QTextFrame * );
extern void _ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame(void* qthis, void* arg0);
  // proto:  int QPlainTextDocumentLayout::pageCount();
extern void _ZNK24QPlainTextDocumentLayout9pageCountEv(void* qthis);
  // proto:  const QMetaObject * QPlainTextDocumentLayout::metaObject();
extern void _ZNK24QPlainTextDocumentLayout10metaObjectEv(void* qthis);
  // proto:  void QPlainTextDocumentLayout::ensureBlockLayout(const QTextBlock & block);
extern void _ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock(void* qthis, void* arg0);
  // proto:  void QPlainTextDocumentLayout::~QPlainTextDocumentLayout();
extern void _ZN24QPlainTextDocumentLayoutD0Ev(void* qthis);
  // proto:  QRectF QPlainTextDocumentLayout::blockBoundingRect(const QTextBlock & block);
extern void _ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock(void* qthis, void* arg0);
  // proto:  int QPlainTextDocumentLayout::cursorWidth();
extern void _ZNK24QPlainTextDocumentLayout11cursorWidthEv(void* qthis);
  // proto:  void QPlainTextDocumentLayout::QPlainTextDocumentLayout(QTextDocument * document);
extern void* dector_ZN24QPlainTextDocumentLayoutC1EP13QTextDocument(void* arg0);
extern void _ZN24QPlainTextDocumentLayoutC1EP13QTextDocument(void* qthis, void* arg0);
  // proto:  QSizeF QPlainTextDocumentLayout::documentSize();
extern void _ZNK24QPlainTextDocumentLayout12documentSizeEv(void* qthis);
  // proto:  QMenu * QPlainTextEdit::createStandardContextMenu(const QPoint & position);
extern void _ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::ensureCursorVisible();
extern void _ZN14QPlainTextEdit19ensureCursorVisibleEv(void* qthis);
  // proto:  QTextDocument * QPlainTextEdit::document();
extern void _ZNK14QPlainTextEdit8documentEv(void* qthis);
  // proto:  QRect QPlainTextEdit::cursorRect();
extern void _ZNK14QPlainTextEdit10cursorRectEv(void* qthis);
  // proto:  void QPlainTextEdit::setTabChangesFocus(bool b);
extern void _ZN14QPlainTextEdit18setTabChangesFocusEb(void* qthis, bool arg0);
  // proto:  QString QPlainTextEdit::toPlainText();
extern void demth_ZNK14QPlainTextEdit11toPlainTextEv(void* qthis);
  // proto:  QVariant QPlainTextEdit::loadResource(int type, const QUrl & name);
extern void _ZN14QPlainTextEdit12loadResourceEiRK4QUrl(void* qthis, int arg0, void* arg1);
  // proto:  int QPlainTextEdit::tabStopWidth();
extern void _ZNK14QPlainTextEdit12tabStopWidthEv(void* qthis);
  // proto:  bool QPlainTextEdit::isReadOnly();
extern void _ZNK14QPlainTextEdit10isReadOnlyEv(void* qthis);
  // proto:  void QPlainTextEdit::setReadOnly(bool ro);
extern void _ZN14QPlainTextEdit11setReadOnlyEb(void* qthis, bool arg0);
  // proto:  QTextCursor QPlainTextEdit::textCursor();
extern void _ZNK14QPlainTextEdit10textCursorEv(void* qthis);
  // proto:  void QPlainTextEdit::setCenterOnScroll(bool enabled);
extern void _ZN14QPlainTextEdit17setCenterOnScrollEb(void* qthis, bool arg0);
  // proto:  QString QPlainTextEdit::placeholderText();
extern void _ZNK14QPlainTextEdit15placeholderTextEv(void* qthis);
  // proto:  int QPlainTextEdit::blockCount();
extern void _ZNK14QPlainTextEdit10blockCountEv(void* qthis);
  // proto:  void QPlainTextEdit::setCurrentCharFormat(const QTextCharFormat & format);
extern void _ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::setDocument(QTextDocument * document);
extern void _ZN14QPlainTextEdit11setDocumentEP13QTextDocument(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::print(QPagedPaintDevice * printer);
extern void _ZNK14QPlainTextEdit5printEP17QPagedPaintDevice(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::setTabStopWidth(int width);
extern void _ZN14QPlainTextEdit15setTabStopWidthEi(void* qthis, int arg0);
  // proto:  bool QPlainTextEdit::backgroundVisible();
extern void _ZNK14QPlainTextEdit17backgroundVisibleEv(void* qthis);
  // proto:  void QPlainTextEdit::redo();
extern void _ZN14QPlainTextEdit4redoEv(void* qthis);
  // proto:  void QPlainTextEdit::QPlainTextEdit(const QString & text, QWidget * parent);
extern void* dector_ZN14QPlainTextEditC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN14QPlainTextEditC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QPlainTextEdit::setOverwriteMode(bool overwrite);
extern void _ZN14QPlainTextEdit16setOverwriteModeEb(void* qthis, bool arg0);
  // proto:  bool QPlainTextEdit::tabChangesFocus();
extern void _ZNK14QPlainTextEdit15tabChangesFocusEv(void* qthis);
  // proto:  void QPlainTextEdit::copy();
extern void _ZN14QPlainTextEdit4copyEv(void* qthis);
  // proto:  void QPlainTextEdit::mergeCurrentCharFormat(const QTextCharFormat & modifier);
extern void _ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(void* qthis, void* arg0);
  // proto:  int QPlainTextEdit::maximumBlockCount();
extern void demth_ZNK14QPlainTextEdit17maximumBlockCountEv(void* qthis);
  // proto:  void QPlainTextEdit::insertPlainText(const QString & text);
extern void _ZN14QPlainTextEdit15insertPlainTextERK7QString(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::setTextCursor(const QTextCursor & cursor);
extern void _ZN14QPlainTextEdit13setTextCursorERK11QTextCursor(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::paste();
extern void _ZN14QPlainTextEdit5pasteEv(void* qthis);
  // proto:  void QPlainTextEdit::zoomIn(int range);
extern void _ZN14QPlainTextEdit6zoomInEi(void* qthis, int arg0);
  // proto:  void QPlainTextEdit::setMaximumBlockCount(int maximum);
extern void demth_ZN14QPlainTextEdit20setMaximumBlockCountEi(void* qthis, int arg0);
  // proto:  QTextCharFormat QPlainTextEdit::currentCharFormat();
extern void _ZNK14QPlainTextEdit17currentCharFormatEv(void* qthis);
  // proto:  void QPlainTextEdit::setCursorWidth(int width);
extern void _ZN14QPlainTextEdit14setCursorWidthEi(void* qthis, int arg0);
  // proto:  QString QPlainTextEdit::documentTitle();
extern void demth_ZNK14QPlainTextEdit13documentTitleEv(void* qthis);
  // proto:  void QPlainTextEdit::selectAll();
extern void _ZN14QPlainTextEdit9selectAllEv(void* qthis);
  // proto:  void QPlainTextEdit::QPlainTextEdit(const QPlainTextEdit & );
extern void* dector_ZN14QPlainTextEditC1ERKS_(void* arg0);
extern void _ZN14QPlainTextEditC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::setPlainText(const QString & text);
extern void _ZN14QPlainTextEdit12setPlainTextERK7QString(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::setBackgroundVisible(bool visible);
extern void _ZN14QPlainTextEdit20setBackgroundVisibleEb(void* qthis, bool arg0);
  // proto:  void QPlainTextEdit::setUndoRedoEnabled(bool enable);
extern void demth_ZN14QPlainTextEdit18setUndoRedoEnabledEb(void* qthis, bool arg0);
  // proto:  bool QPlainTextEdit::overwriteMode();
extern void _ZNK14QPlainTextEdit13overwriteModeEv(void* qthis);
  // proto:  void QPlainTextEdit::centerCursor();
extern void _ZN14QPlainTextEdit12centerCursorEv(void* qthis);
  // proto:  const QMetaObject * QPlainTextEdit::metaObject();
extern void _ZNK14QPlainTextEdit10metaObjectEv(void* qthis);
  // proto:  QMenu * QPlainTextEdit::createStandardContextMenu();
extern void _ZN14QPlainTextEdit25createStandardContextMenuEv(void* qthis);
  // proto:  void QPlainTextEdit::setDocumentTitle(const QString & title);
extern void demth_ZN14QPlainTextEdit16setDocumentTitleERK7QString(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::~QPlainTextEdit();
extern void _ZN14QPlainTextEditD0Ev(void* qthis);
  // proto:  void QPlainTextEdit::clear();
extern void _ZN14QPlainTextEdit5clearEv(void* qthis);
  // proto:  QString QPlainTextEdit::anchorAt(const QPoint & pos);
extern void _ZNK14QPlainTextEdit8anchorAtERK6QPoint(void* qthis, void* arg0);
  // proto:  bool QPlainTextEdit::canPaste();
extern void _ZNK14QPlainTextEdit8canPasteEv(void* qthis);
  // proto:  void QPlainTextEdit::QPlainTextEdit(QWidget * parent);
extern void* dector_ZN14QPlainTextEditC1EP7QWidget(void* arg0);
extern void _ZN14QPlainTextEditC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::cut();
extern void _ZN14QPlainTextEdit3cutEv(void* qthis);
  // proto:  void QPlainTextEdit::appendHtml(const QString & html);
extern void _ZN14QPlainTextEdit10appendHtmlERK7QString(void* qthis, void* arg0);
  // proto:  bool QPlainTextEdit::isUndoRedoEnabled();
extern void demth_ZNK14QPlainTextEdit17isUndoRedoEnabledEv(void* qthis);
  // proto:  void QPlainTextEdit::zoomOut(int range);
extern void _ZN14QPlainTextEdit7zoomOutEi(void* qthis, int arg0);
  // proto:  void QPlainTextEdit::setPlaceholderText(const QString & placeholderText);
extern void _ZN14QPlainTextEdit18setPlaceholderTextERK7QString(void* qthis, void* arg0);
  // proto:  void QPlainTextEdit::undo();
extern void _ZN14QPlainTextEdit4undoEv(void* qthis);
  // proto:  QTextCursor QPlainTextEdit::cursorForPosition(const QPoint & pos);
extern void _ZNK14QPlainTextEdit17cursorForPositionERK6QPoint(void* qthis, void* arg0);
  // proto:  bool QPlainTextEdit::centerOnScroll();
extern void _ZNK14QPlainTextEdit14centerOnScrollEv(void* qthis);
  // proto:  void QPlainTextEdit::appendPlainText(const QString & text);
extern void _ZN14QPlainTextEdit15appendPlainTextERK7QString(void* qthis, void* arg0);
  // proto:  int QPlainTextEdit::cursorWidth();
extern void _ZNK14QPlainTextEdit11cursorWidthEv(void* qthis);
  // proto:  QRect QPlainTextEdit::cursorRect(const QTextCursor & cursor);
extern void _ZNK14QPlainTextEdit10cursorRectERK11QTextCursor(void* qthis, void* arg0);
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

// class sizeof(QPlainTextDocumentLayout)=1
type QPlainTextDocumentLayout struct {
  /*qbase*/ QAbstractTextDocumentLayout;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPlainTextEdit)=1
type QPlainTextEdit struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst uint64 /* *mut c_void*/;
//  _cursorPositionChanged QPlainTextEdit_cursorPositionChanged_signal;
//  _modificationChanged QPlainTextEdit_modificationChanged_signal;
//  _redoAvailable QPlainTextEdit_redoAvailable_signal;
//  _selectionChanged QPlainTextEdit_selectionChanged_signal;
//  _updateRequest QPlainTextEdit_updateRequest_signal;
//  _blockCountChanged QPlainTextEdit_blockCountChanged_signal;
//  _undoAvailable QPlainTextEdit_undoAvailable_signal;
//  _textChanged QPlainTextEdit_textChanged_signal;
//  _copyAvailable QPlainTextEdit_copyAvailable_signal;
}

  // proto:  void QPlainTextDocumentLayout::requestUpdate();
func (this *QPlainTextDocumentLayout) requestUpdate(args ...interface{}) () {
  // requestUpdate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QPlainTextDocumentLayout13requestUpdateEv
    // invoke: void requestUpdate()
    C._ZN24QPlainTextDocumentLayout13requestUpdateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "requestUpdate", args)
  }

}

  // proto:  void QPlainTextDocumentLayout::setCursorWidth(int width);
func (this *QPlainTextDocumentLayout) setCursorWidth(args ...interface{}) () {
  // setCursorWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QPlainTextDocumentLayout14setCursorWidthEi
    // invoke: void setCursorWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN24QPlainTextDocumentLayout14setCursorWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "setCursorWidth", args)
  }

}

  // proto:  QRectF QPlainTextDocumentLayout::frameBoundingRect(QTextFrame * );
func (this *QPlainTextDocumentLayout) frameBoundingRect(args ...interface{}) () {
  // frameBoundingRect(class QTextFrame *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame
    // invoke: QRectF frameBoundingRect(class QTextFrame *)
    var arg0 = args[0].(QTextFrame).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK24QPlainTextDocumentLayout17frameBoundingRectEP10QTextFrame(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "frameBoundingRect", args)
  }

}

  // proto:  int QPlainTextDocumentLayout::pageCount();
func (this *QPlainTextDocumentLayout) pageCount(args ...interface{}) () {
  // pageCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout9pageCountEv
    // invoke: int pageCount()
    C._ZNK24QPlainTextDocumentLayout9pageCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "pageCount", args)
  }

}

  // proto:  const QMetaObject * QPlainTextDocumentLayout::metaObject();
func (this *QPlainTextDocumentLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK24QPlainTextDocumentLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "metaObject", args)
  }

}

  // proto:  void QPlainTextDocumentLayout::ensureBlockLayout(const QTextBlock & block);
func (this *QPlainTextDocumentLayout) ensureBlockLayout(args ...interface{}) () {
  // ensureBlockLayout(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock
    // invoke: void ensureBlockLayout(const class QTextBlock &)
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK24QPlainTextDocumentLayout17ensureBlockLayoutERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "ensureBlockLayout", args)
  }

}

  // proto:  void QPlainTextDocumentLayout::~QPlainTextDocumentLayout();
func (this *QPlainTextDocumentLayout) FreeQPlainTextDocumentLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "~QPlainTextDocumentLayout", args)
  }

}

  // proto:  QRectF QPlainTextDocumentLayout::blockBoundingRect(const QTextBlock & block);
func (this *QPlainTextDocumentLayout) blockBoundingRect(args ...interface{}) () {
  // blockBoundingRect(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock
    // invoke: QRectF blockBoundingRect(const class QTextBlock &)
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK24QPlainTextDocumentLayout17blockBoundingRectERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "blockBoundingRect", args)
  }

}

  // proto:  int QPlainTextDocumentLayout::cursorWidth();
func (this *QPlainTextDocumentLayout) cursorWidth(args ...interface{}) () {
  // cursorWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout11cursorWidthEv
    // invoke: int cursorWidth()
    C._ZNK24QPlainTextDocumentLayout11cursorWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "cursorWidth", args)
  }

}

  // proto:  void QPlainTextDocumentLayout::QPlainTextDocumentLayout(QTextDocument * document);
func NewQPlainTextDocumentLayout(args ...interface{}) QPlainTextDocumentLayout {
  return QPlainTextDocumentLayout{}
}

  // proto:  QSizeF QPlainTextDocumentLayout::documentSize();
func (this *QPlainTextDocumentLayout) documentSize(args ...interface{}) () {
  // documentSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QPlainTextDocumentLayout12documentSizeEv
    // invoke: QSizeF documentSize()
    C._ZNK24QPlainTextDocumentLayout12documentSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextDocumentLayout", "documentSize", args)
  }

}

  // proto:  QMenu * QPlainTextEdit::createStandardContextMenu(const QPoint & position);
func (this *QPlainTextEdit) createStandardContextMenu(args ...interface{}) () {
  // createStandardContextMenu(const class QPoint &)
  // createStandardContextMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint
    // invoke: QMenu * createStandardContextMenu(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit25createStandardContextMenuERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN14QPlainTextEdit25createStandardContextMenuEv
    // invoke: QMenu * createStandardContextMenu()
    C._ZN14QPlainTextEdit25createStandardContextMenuEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "createStandardContextMenu", args)
  }

}

  // proto:  void QPlainTextEdit::ensureCursorVisible();
func (this *QPlainTextEdit) ensureCursorVisible(args ...interface{}) () {
  // ensureCursorVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit19ensureCursorVisibleEv
    // invoke: void ensureCursorVisible()
    C._ZN14QPlainTextEdit19ensureCursorVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "ensureCursorVisible", args)
  }

}

  // proto:  QTextDocument * QPlainTextEdit::document();
func (this *QPlainTextEdit) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit8documentEv
    // invoke: QTextDocument * document()
    C._ZNK14QPlainTextEdit8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "document", args)
  }

}

  // proto:  QRect QPlainTextEdit::cursorRect();
func (this *QPlainTextEdit) cursorRect(args ...interface{}) () {
  // cursorRect()
  // cursorRect(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10cursorRectEv
    // invoke: QRect cursorRect()
    C._ZNK14QPlainTextEdit10cursorRectEv(this.qclsinst)
  case 1:
    // invoke: _ZNK14QPlainTextEdit10cursorRectERK11QTextCursor
    // invoke: QRect cursorRect(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QPlainTextEdit10cursorRectERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorRect", args)
  }

}

  // proto:  void QPlainTextEdit::setTabChangesFocus(bool b);
func (this *QPlainTextEdit) setTabChangesFocus(args ...interface{}) () {
  // setTabChangesFocus(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit18setTabChangesFocusEb
    // invoke: void setTabChangesFocus(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit18setTabChangesFocusEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTabChangesFocus", args)
  }

}

  // proto:  QString QPlainTextEdit::toPlainText();
func (this *QPlainTextEdit) toPlainText(args ...interface{}) () {
  // toPlainText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit11toPlainTextEv
    // invoke: QString toPlainText()
    C.demth_ZNK14QPlainTextEdit11toPlainTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "toPlainText", args)
  }

}

  // proto:  QVariant QPlainTextEdit::loadResource(int type, const QUrl & name);
func (this *QPlainTextEdit) loadResource(args ...interface{}) () {
  // loadResource(int, const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit12loadResourceEiRK4QUrl
    // invoke: QVariant loadResource(int, const class QUrl &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN14QPlainTextEdit12loadResourceEiRK4QUrl(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "loadResource", args)
  }

}

  // proto:  int QPlainTextEdit::tabStopWidth();
func (this *QPlainTextEdit) tabStopWidth(args ...interface{}) () {
  // tabStopWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit12tabStopWidthEv
    // invoke: int tabStopWidth()
    C._ZNK14QPlainTextEdit12tabStopWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "tabStopWidth", args)
  }

}

  // proto:  bool QPlainTextEdit::isReadOnly();
func (this *QPlainTextEdit) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10isReadOnlyEv
    // invoke: bool isReadOnly()
    C._ZNK14QPlainTextEdit10isReadOnlyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "isReadOnly", args)
  }

}

  // proto:  void QPlainTextEdit::setReadOnly(bool ro);
func (this *QPlainTextEdit) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit11setReadOnlyEb
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit11setReadOnlyEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setReadOnly", args)
  }

}

  // proto:  QTextCursor QPlainTextEdit::textCursor();
func (this *QPlainTextEdit) textCursor(args ...interface{}) () {
  // textCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10textCursorEv
    // invoke: QTextCursor textCursor()
    C._ZNK14QPlainTextEdit10textCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "textCursor", args)
  }

}

  // proto:  void QPlainTextEdit::setCenterOnScroll(bool enabled);
func (this *QPlainTextEdit) setCenterOnScroll(args ...interface{}) () {
  // setCenterOnScroll(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit17setCenterOnScrollEb
    // invoke: void setCenterOnScroll(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit17setCenterOnScrollEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCenterOnScroll", args)
  }

}

  // proto:  QString QPlainTextEdit::placeholderText();
func (this *QPlainTextEdit) placeholderText(args ...interface{}) () {
  // placeholderText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit15placeholderTextEv
    // invoke: QString placeholderText()
    C._ZNK14QPlainTextEdit15placeholderTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "placeholderText", args)
  }

}

  // proto:  int QPlainTextEdit::blockCount();
func (this *QPlainTextEdit) blockCount(args ...interface{}) () {
  // blockCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10blockCountEv
    // invoke: int blockCount()
    C._ZNK14QPlainTextEdit10blockCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "blockCount", args)
  }

}

  // proto:  void QPlainTextEdit::setCurrentCharFormat(const QTextCharFormat & format);
func (this *QPlainTextEdit) setCurrentCharFormat(args ...interface{}) () {
  // setCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat
    // invoke: void setCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit20setCurrentCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCurrentCharFormat", args)
  }

}

  // proto:  void QPlainTextEdit::setDocument(QTextDocument * document);
func (this *QPlainTextEdit) setDocument(args ...interface{}) () {
  // setDocument(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit11setDocumentEP13QTextDocument
    // invoke: void setDocument(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit11setDocumentEP13QTextDocument(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setDocument", args)
  }

}

  // proto:  void QPlainTextEdit::print(QPagedPaintDevice * printer);
func (this *QPlainTextEdit) print(args ...interface{}) () {
  // print(class QPagedPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPagedPaintDevice{}) // "QPagedPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit5printEP17QPagedPaintDevice
    // invoke: void print(class QPagedPaintDevice *)
    var arg0 = args[0].(QPagedPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QPlainTextEdit5printEP17QPagedPaintDevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "print", args)
  }

}

  // proto:  void QPlainTextEdit::setTabStopWidth(int width);
func (this *QPlainTextEdit) setTabStopWidth(args ...interface{}) () {
  // setTabStopWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit15setTabStopWidthEi
    // invoke: void setTabStopWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit15setTabStopWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTabStopWidth", args)
  }

}

  // proto:  bool QPlainTextEdit::backgroundVisible();
func (this *QPlainTextEdit) backgroundVisible(args ...interface{}) () {
  // backgroundVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17backgroundVisibleEv
    // invoke: bool backgroundVisible()
    C._ZNK14QPlainTextEdit17backgroundVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "backgroundVisible", args)
  }

}

  // proto:  void QPlainTextEdit::redo();
func (this *QPlainTextEdit) redo(args ...interface{}) () {
  // redo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit4redoEv
    // invoke: void redo()
    C._ZN14QPlainTextEdit4redoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "redo", args)
  }

}

  // proto:  void QPlainTextEdit::QPlainTextEdit(const QString & text, QWidget * parent);
func NewQPlainTextEdit(args ...interface{}) QPlainTextEdit {
  return QPlainTextEdit{}
}

  // proto:  void QPlainTextEdit::setOverwriteMode(bool overwrite);
func (this *QPlainTextEdit) setOverwriteMode(args ...interface{}) () {
  // setOverwriteMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit16setOverwriteModeEb
    // invoke: void setOverwriteMode(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit16setOverwriteModeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setOverwriteMode", args)
  }

}

  // proto:  bool QPlainTextEdit::tabChangesFocus();
func (this *QPlainTextEdit) tabChangesFocus(args ...interface{}) () {
  // tabChangesFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit15tabChangesFocusEv
    // invoke: bool tabChangesFocus()
    C._ZNK14QPlainTextEdit15tabChangesFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "tabChangesFocus", args)
  }

}

  // proto:  void QPlainTextEdit::copy();
func (this *QPlainTextEdit) copy(args ...interface{}) () {
  // copy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit4copyEv
    // invoke: void copy()
    C._ZN14QPlainTextEdit4copyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "copy", args)
  }

}

  // proto:  void QPlainTextEdit::mergeCurrentCharFormat(const QTextCharFormat & modifier);
func (this *QPlainTextEdit) mergeCurrentCharFormat(args ...interface{}) () {
  // mergeCurrentCharFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat
    // invoke: void mergeCurrentCharFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit22mergeCurrentCharFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "mergeCurrentCharFormat", args)
  }

}

  // proto:  int QPlainTextEdit::maximumBlockCount();
func (this *QPlainTextEdit) maximumBlockCount(args ...interface{}) () {
  // maximumBlockCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17maximumBlockCountEv
    // invoke: int maximumBlockCount()
    C.demth_ZNK14QPlainTextEdit17maximumBlockCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "maximumBlockCount", args)
  }

}

  // proto:  void QPlainTextEdit::insertPlainText(const QString & text);
func (this *QPlainTextEdit) insertPlainText(args ...interface{}) () {
  // insertPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit15insertPlainTextERK7QString
    // invoke: void insertPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit15insertPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "insertPlainText", args)
  }

}

  // proto:  void QPlainTextEdit::setTextCursor(const QTextCursor & cursor);
func (this *QPlainTextEdit) setTextCursor(args ...interface{}) () {
  // setTextCursor(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit13setTextCursorERK11QTextCursor
    // invoke: void setTextCursor(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit13setTextCursorERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setTextCursor", args)
  }

}

  // proto:  void QPlainTextEdit::paste();
func (this *QPlainTextEdit) paste(args ...interface{}) () {
  // paste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit5pasteEv
    // invoke: void paste()
    C._ZN14QPlainTextEdit5pasteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "paste", args)
  }

}

  // proto:  void QPlainTextEdit::zoomIn(int range);
func (this *QPlainTextEdit) zoomIn(args ...interface{}) () {
  // zoomIn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit6zoomInEi
    // invoke: void zoomIn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit6zoomInEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "zoomIn", args)
  }

}

  // proto:  void QPlainTextEdit::setMaximumBlockCount(int maximum);
func (this *QPlainTextEdit) setMaximumBlockCount(args ...interface{}) () {
  // setMaximumBlockCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit20setMaximumBlockCountEi
    // invoke: void setMaximumBlockCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN14QPlainTextEdit20setMaximumBlockCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setMaximumBlockCount", args)
  }

}

  // proto:  QTextCharFormat QPlainTextEdit::currentCharFormat();
func (this *QPlainTextEdit) currentCharFormat(args ...interface{}) () {
  // currentCharFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17currentCharFormatEv
    // invoke: QTextCharFormat currentCharFormat()
    C._ZNK14QPlainTextEdit17currentCharFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "currentCharFormat", args)
  }

}

  // proto:  void QPlainTextEdit::setCursorWidth(int width);
func (this *QPlainTextEdit) setCursorWidth(args ...interface{}) () {
  // setCursorWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit14setCursorWidthEi
    // invoke: void setCursorWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit14setCursorWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setCursorWidth", args)
  }

}

  // proto:  QString QPlainTextEdit::documentTitle();
func (this *QPlainTextEdit) documentTitle(args ...interface{}) () {
  // documentTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit13documentTitleEv
    // invoke: QString documentTitle()
    C.demth_ZNK14QPlainTextEdit13documentTitleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "documentTitle", args)
  }

}

  // proto:  void QPlainTextEdit::selectAll();
func (this *QPlainTextEdit) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit9selectAllEv
    // invoke: void selectAll()
    C._ZN14QPlainTextEdit9selectAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "selectAll", args)
  }

}

  // proto:  void QPlainTextEdit::setPlainText(const QString & text);
func (this *QPlainTextEdit) setPlainText(args ...interface{}) () {
  // setPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit12setPlainTextERK7QString
    // invoke: void setPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit12setPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setPlainText", args)
  }

}

  // proto:  void QPlainTextEdit::setBackgroundVisible(bool visible);
func (this *QPlainTextEdit) setBackgroundVisible(args ...interface{}) () {
  // setBackgroundVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit20setBackgroundVisibleEb
    // invoke: void setBackgroundVisible(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit20setBackgroundVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setBackgroundVisible", args)
  }

}

  // proto:  void QPlainTextEdit::setUndoRedoEnabled(bool enable);
func (this *QPlainTextEdit) setUndoRedoEnabled(args ...interface{}) () {
  // setUndoRedoEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit18setUndoRedoEnabledEb
    // invoke: void setUndoRedoEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C.demth_ZN14QPlainTextEdit18setUndoRedoEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setUndoRedoEnabled", args)
  }

}

  // proto:  bool QPlainTextEdit::overwriteMode();
func (this *QPlainTextEdit) overwriteMode(args ...interface{}) () {
  // overwriteMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit13overwriteModeEv
    // invoke: bool overwriteMode()
    C._ZNK14QPlainTextEdit13overwriteModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "overwriteMode", args)
  }

}

  // proto:  void QPlainTextEdit::centerCursor();
func (this *QPlainTextEdit) centerCursor(args ...interface{}) () {
  // centerCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit12centerCursorEv
    // invoke: void centerCursor()
    C._ZN14QPlainTextEdit12centerCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "centerCursor", args)
  }

}

  // proto:  const QMetaObject * QPlainTextEdit::metaObject();
func (this *QPlainTextEdit) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK14QPlainTextEdit10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "metaObject", args)
  }

}

  // proto:  void QPlainTextEdit::setDocumentTitle(const QString & title);
func (this *QPlainTextEdit) setDocumentTitle(args ...interface{}) () {
  // setDocumentTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit16setDocumentTitleERK7QString
    // invoke: void setDocumentTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN14QPlainTextEdit16setDocumentTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setDocumentTitle", args)
  }

}

  // proto:  void QPlainTextEdit::~QPlainTextEdit();
func (this *QPlainTextEdit) FreeQPlainTextEdit(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "~QPlainTextEdit", args)
  }

}

  // proto:  void QPlainTextEdit::clear();
func (this *QPlainTextEdit) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit5clearEv
    // invoke: void clear()
    C._ZN14QPlainTextEdit5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "clear", args)
  }

}

  // proto:  QString QPlainTextEdit::anchorAt(const QPoint & pos);
func (this *QPlainTextEdit) anchorAt(args ...interface{}) () {
  // anchorAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit8anchorAtERK6QPoint
    // invoke: QString anchorAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QPlainTextEdit8anchorAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "anchorAt", args)
  }

}

  // proto:  bool QPlainTextEdit::canPaste();
func (this *QPlainTextEdit) canPaste(args ...interface{}) () {
  // canPaste()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit8canPasteEv
    // invoke: bool canPaste()
    C._ZNK14QPlainTextEdit8canPasteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "canPaste", args)
  }

}

  // proto:  void QPlainTextEdit::cut();
func (this *QPlainTextEdit) cut(args ...interface{}) () {
  // cut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit3cutEv
    // invoke: void cut()
    C._ZN14QPlainTextEdit3cutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cut", args)
  }

}

  // proto:  void QPlainTextEdit::appendHtml(const QString & html);
func (this *QPlainTextEdit) appendHtml(args ...interface{}) () {
  // appendHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit10appendHtmlERK7QString
    // invoke: void appendHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit10appendHtmlERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "appendHtml", args)
  }

}

  // proto:  bool QPlainTextEdit::isUndoRedoEnabled();
func (this *QPlainTextEdit) isUndoRedoEnabled(args ...interface{}) () {
  // isUndoRedoEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17isUndoRedoEnabledEv
    // invoke: bool isUndoRedoEnabled()
    C.demth_ZNK14QPlainTextEdit17isUndoRedoEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "isUndoRedoEnabled", args)
  }

}

  // proto:  void QPlainTextEdit::zoomOut(int range);
func (this *QPlainTextEdit) zoomOut(args ...interface{}) () {
  // zoomOut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit7zoomOutEi
    // invoke: void zoomOut(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit7zoomOutEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "zoomOut", args)
  }

}

  // proto:  void QPlainTextEdit::setPlaceholderText(const QString & placeholderText);
func (this *QPlainTextEdit) setPlaceholderText(args ...interface{}) () {
  // setPlaceholderText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit18setPlaceholderTextERK7QString
    // invoke: void setPlaceholderText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit18setPlaceholderTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "setPlaceholderText", args)
  }

}

  // proto:  void QPlainTextEdit::undo();
func (this *QPlainTextEdit) undo(args ...interface{}) () {
  // undo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit4undoEv
    // invoke: void undo()
    C._ZN14QPlainTextEdit4undoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "undo", args)
  }

}

  // proto:  QTextCursor QPlainTextEdit::cursorForPosition(const QPoint & pos);
func (this *QPlainTextEdit) cursorForPosition(args ...interface{}) () {
  // cursorForPosition(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit17cursorForPositionERK6QPoint
    // invoke: QTextCursor cursorForPosition(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QPlainTextEdit17cursorForPositionERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorForPosition", args)
  }

}

  // proto:  bool QPlainTextEdit::centerOnScroll();
func (this *QPlainTextEdit) centerOnScroll(args ...interface{}) () {
  // centerOnScroll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit14centerOnScrollEv
    // invoke: bool centerOnScroll()
    C._ZNK14QPlainTextEdit14centerOnScrollEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "centerOnScroll", args)
  }

}

  // proto:  void QPlainTextEdit::appendPlainText(const QString & text);
func (this *QPlainTextEdit) appendPlainText(args ...interface{}) () {
  // appendPlainText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QPlainTextEdit15appendPlainTextERK7QString
    // invoke: void appendPlainText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QPlainTextEdit15appendPlainTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "appendPlainText", args)
  }

}

  // proto:  int QPlainTextEdit::cursorWidth();
func (this *QPlainTextEdit) cursorWidth(args ...interface{}) () {
  // cursorWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QPlainTextEdit11cursorWidthEv
    // invoke: int cursorWidth()
    C._ZNK14QPlainTextEdit11cursorWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPlainTextEdit", "cursorWidth", args)
  }

}

// <= body block end

