package chrome

import "github.com/gopherjs/gopherjs/js"

const (
	WINDOW_ID_NONE    = -1
	WINDOW_ID_CURRENT = -2
)

type Windows struct {
	o js.Object
}

/*
* Methods
 */

func (w *Windows) Get(windowId int, getInfo interface{}, callback func(js.Object)) {
	w.o.Call("get", windowId, getInfo, callback)
}

func (w *Windows) GetCurrent(getInfo interface{}, callback func(js.Object)) {
	w.o.Call("getCurrent", getInfo, callback)
}

func (w *Windows) GetLastFocused(getInfo interface{}, callback func(js.Object)) {
	w.o.Call("getLastFocused", getInfo, callback)
}

func (w *Windows) GetAll(getInfo interface{}, callback func(js.Object)) {
	w.o.Call("getAll", getInfo, callback)
}

func (w *Windows) Create(createData interface{}, callback func(js.Object)) {
	w.o.Call("create", createData, callback)
}

func (w *Windows) Update(windowId int, updateInfo interface{}, callback func(js.Object)) {
	w.o.Call("update", windowId, updateInfo, callback)
}

func (w *Windows) Remove(windowId int, callback func(js.Object)) {
	w.o.Call("remove", windowId, callback)
}

/*
* Events
 */

func (w *Windows) OnCreated(callback func(window map[string]interface{})) {
	w.o.Get("onCreated").Call("addListener", callback)
}

func (w *Windows) OnRemoved(callback func(windowId int)) {
	w.o.Get("onRemoved").Call("addListener", callback)
}

func (w *Windows) onFocusChanged(callback func(windowId int)) {
	w.o.Get("onFocusChanged").Call("addListener", callback)
}
