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

// Get gets details about a window.
func (w *Windows) Get(windowId int, getInfo interface{}, callback func(js.Object)) {
	w.o.Call("get", windowId, getInfo, callback)
}

// GetCurrent gets the current window.
func (w *Windows) GetCurrent(getInfo interface{}, callback func(js.Object)) {
	w.o.Call("getCurrent", getInfo, callback)
}

// GetLastFocused gets the window that was most recently focused — typically the window 'on top'.
func (w *Windows) GetLastFocused(getInfo interface{}, callback func(js.Object)) {
	w.o.Call("getLastFocused", getInfo, callback)
}

// GetAll gets all windows.
func (w *Windows) GetAll(getInfo interface{}, callback func(js.Object)) {
	w.o.Call("getAll", getInfo, callback)
}

// Create creates (opens) a new browser with any optional sizing, position or default URL provided.
func (w *Windows) Create(createData interface{}, callback func(js.Object)) {
	w.o.Call("create", createData, callback)
}

// Update updates the properties of a window. Specify only the properties that you
// want to change; unspecified properties will be left unchanged.
func (w *Windows) Update(windowId int, updateInfo interface{}, callback func(js.Object)) {
	w.o.Call("update", windowId, updateInfo, callback)
}

// Remove removes (closes) a window, and all the tabs inside it.
func (w *Windows) Remove(windowId int, callback func(js.Object)) {
	w.o.Call("remove", windowId, callback)
}

/*
* Events
 */

// OnCreated fired when a window is created.
func (w *Windows) OnCreated(callback func(window map[string]interface{})) {
	w.o.Get("onCreated").Call("addListener", callback)
}

// OnRemoved fired when a window is removed (closed).
func (w *Windows) OnRemoved(callback func(windowId int)) {
	w.o.Get("onRemoved").Call("addListener", callback)
}

// onFocusChanged fired when the currently focused window changes.
// Will be chrome.windows.WINDOW_ID_NONE if all chrome windows have lost focus.
// Note: On some Linux window managers, WINDOW_ID_NONE will always be sent immediately
// preceding a switch from one chrome window to another.
func (w *Windows) onFocusChanged(callback func(windowId int)) {
	w.o.Get("onFocusChanged").Call("addListener", callback)
}
