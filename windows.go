package chrome

import "github.com/gopherjs/gopherjs/js"

type Windows struct {
	o                 *js.Object
	WINDOW_ID_NONE    int
	WINDOW_ID_CURRENT int
}

func NewWindows(windowsObj *js.Object) *Windows {
	w := new(Windows)
	w.o = windowsObj
	if w.o.String() != "undefined" {
		w.WINDOW_ID_CURRENT = w.o.Get("WINDOW_ID_CURRENT").Int()
		w.WINDOW_ID_NONE = w.o.Get("WINDOW_ID_NONE").Int()
	}
	return w
}

/*
* Types
 */

type Window struct {
	*js.Object
	Id          int    `js:"id"`
	Focused     bool   `js:"focused"`
	Top         int    `js:"top"`
	Left        int    `js:"left"`
	Width       int    `js:"width"`
	Height      int    `js:"height"`
	Tabs        []Tab  `js:"tabs"`
	Incognito   bool   `js:"incognito"`
	Type        string `js:"type"`
	State       string `js:"state"`
	AlwaysOnTop bool   `js:"alwaysOnTop"`
	SessionId   string `js:"sessionId"`
}

/*
* Methods
 */

// Get gets details about a window.
func (w *Windows) Get(windowId int, getInfo Object, callback func(window Window)) {
	w.o.Call("get", windowId, getInfo, callback)
}

// GetCurrent gets the current window.
func (w *Windows) GetCurrent(getInfo Object, callback func(window Window)) {
	w.o.Call("getCurrent", getInfo, callback)
}

// GetLastFocused gets the window that was most recently focused — typically the window 'on top'.
func (w *Windows) GetLastFocused(getInfo Object, callback func(window Window)) {
	w.o.Call("getLastFocused", getInfo, callback)
}

// GetAll gets all windows.
func (w *Windows) GetAll(getInfo Object, callback func(windows []Window)) {
	w.o.Call("getAll", getInfo, callback)
}

// Create creates (opens) a new browser with any optional sizing, position or default URL provided.
func (w *Windows) Create(createData Object, callback func(window Window)) {
	w.o.Call("create", createData, callback)
}

// Update updates the properties of a window. Specify only the properties that you
// want to change; unspecified properties will be left unchanged.
func (w *Windows) Update(windowId int, updateInfo Object, callback func(window Window)) {
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
func (w *Windows) OnCreated(callback func(window Window)) {
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
