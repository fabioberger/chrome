package chrome

import "github.com/gopherjs/gopherjs/js"

const (
	WINDOW_ID_NONE    = -1
	WINDOW_ID_CURRENT = -2
)

type Tabs struct {
	o js.Object
}

func (t *Tabs) Get(tabId int, callback func(js.Object)) {
	t.o.Call("get", tabId, callback)
}

func (t *Tabs) GetCurrent(callback func(js.Object)) {
	t.o.Call("getCurrent", callback)
}

func (t *Tabs) Connect(tabId int, connectInfo interface{}) {
	t.o.Call("connect", tabId, connectInfo)
}

func (t *Tabs) SendMessage(tabId int, message interface{}, responseCallback func(js.Object)) {
	t.o.Call("sendMessage", tabId, message, responseCallback)
}

func (t *Tabs) CurrentSelected(windowId int, callback func(js.Object)) {
	t.o.Call("getSelected", windowId, callback)
}

func (t *Tabs) GetAllInWindow(windowId int, callback func(js.Object)) {
	t.o.Call("getAllInWindow", windowId, callback)
}

func (t *Tabs) Create(createProperties interface{}, callback func(js.Object)) {
	t.o.Call("create", createProperties, callback)
}

func (t *Tabs) Duplicate(tabId int, callback func(js.Object)) {
	t.o.Call("duplicate", tabId, callback)
}

func (t *Tabs) Query(queryInfo interface{}, callback func(js.Object)) {
	t.o.Call("query", queryInfo, callback)
}

func (t *Tabs) Highlight(highlightInfo interface{}, callback func(js.Object)) {
	t.o.Call("highlight", highlightInfo, callback)
}

func (t *Tabs) Update(tabId int, updateProperties interface{}, callback func(js.Object)) {
	t.o.Call("highlight", updateProperties, callback)
}

func (t *Tabs) Move(tabIds []interface{}, moveProperties interface{}, callback func(js.Object)) {
	t.o.Call("move", tabIds, moveProperties, callback)
}

func (t *Tabs) Reload(tabId int, reloadProperties interface{}, callback func(js.Object)) {
	t.o.Call("reload", tabId, reloadProperties, callback)
}

func (t *Tabs) Remove(tabIds []interface{}, callback func(js.Object)) {
	t.o.Call("remove", tabIds, callback)
}

func (t *Tabs) DetectLanguage(tabId int, callback func(js.Object)) {
	t.o.Call("detectLanguage", tabId, callback)
}

func (t *Tabs) CaptureVisibleTab(windowId int, options interface{}, callback func(js.Object)) {
	t.o.Call("captureVisibleTab", windowId, options, callback)
}

func (t *Tabs) ExecuteScript(tabId int, details interface{}, callback func(js.Object)) {
	t.o.Call("executeScript", tabId, details, callback)
}

func (t *Tabs) InsertCss(tabId int, details interface{}, callback func(js.Object)) {
	t.o.Call("insertCss", tabId, details, callback)
}
