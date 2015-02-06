package chrome

import "github.com/gopherjs/gopherjs/js"

type Tabs struct {
	o js.Object
}

/*
* Methods:
 */

// Get retrieves details about the specified tab.
func (t *Tabs) Get(tabId int, callback func(js.Object)) {
	t.o.Call("get", tabId, callback)
}

// GetCurrent gets the tab that this script call is being made from.
// May be undefined if called from a non-tab context (for example: a background page or popup view).
func (t *Tabs) GetCurrent(callback func(js.Object)) {
	t.o.Call("getCurrent", callback)
}

// Connect Connects to the content script(s) in the specified tab.
// The runtime.onConnect event is fired in each content script running in the specified tab
// for the current extension. For more details, see Content Script Messaging (https://developer.chrome.com/extensions/messaging)
func (t *Tabs) Connect(tabId int, connectInfo interface{}) {
	t.o.Call("connect", tabId, connectInfo)
}

// SendMessage sends a single message to the content script(s) in the specified tab,
//  with an optional callback to run when a response is sent back. The runtime.onMessage
// event is fired in each content script running in the specified tab for the current extension.
func (t *Tabs) SendMessage(tabId int, message interface{}, responseCallback func(js.Object)) {
	t.o.Call("sendMessage", tabId, message, responseCallback)
}

// GetSelected gets the tab that is selected in the specified window.
func (t *Tabs) GetSelected(windowId int, callback func(js.Object)) {
	t.o.Call("getSelected", windowId, callback)
}

// GetAllInWindow gets details about all tabs in the specified window.
func (t *Tabs) GetAllInWindow(windowId int, callback func(js.Object)) {
	t.o.Call("getAllInWindow", windowId, callback)
}

// Create creates a new tab.
func (t *Tabs) Create(createProperties interface{}, callback func(js.Object)) {
	t.o.Call("create", createProperties, callback)
}

// Duplicate duplicates a tab.
func (t *Tabs) Duplicate(tabId int, callback func(js.Object)) {
	t.o.Call("duplicate", tabId, callback)
}

// Query gets all tabs that have the specified properties, or all tabs if no properties are specified.
func (t *Tabs) Query(queryInfo interface{}, callback func(js.Object)) {
	t.o.Call("query", queryInfo, callback)
}

// Highlight highlights the given tabs
func (t *Tabs) Highlight(highlightInfo interface{}, callback func(js.Object)) {
	t.o.Call("highlight", highlightInfo, callback)
}

// Update modifies the properties of a tab. Properties that are not specified in updateProperties are not modified.
func (t *Tabs) Update(tabId int, updateProperties interface{}, callback func(js.Object)) {
	t.o.Call("highlight", updateProperties, callback)
}

// Move moves one or more tabs to a new position within its window, or to a new window.
// Note that tabs can only be moved to and from normal windows.
func (t *Tabs) Move(tabIds []interface{}, moveProperties interface{}, callback func(js.Object)) {
	t.o.Call("move", tabIds, moveProperties, callback)
}

// Reload reloads a tab.
func (t *Tabs) Reload(tabId int, reloadProperties interface{}, callback func(js.Object)) {
	t.o.Call("reload", tabId, reloadProperties, callback)
}

// Remove closes one or more tabs.
func (t *Tabs) Remove(tabIds []interface{}, callback func(js.Object)) {
	t.o.Call("remove", tabIds, callback)
}

// DetectLanguage detects the primary language of the content in a tab.
func (t *Tabs) DetectLanguage(tabId int, callback func(js.Object)) {
	t.o.Call("detectLanguage", tabId, callback)
}

// CaptureVisibleTab captures the visible area of the currently active tab in the specified window.
// You must have <all_urls> permission to use this method.
func (t *Tabs) CaptureVisibleTab(windowId int, options interface{}, callback func(js.Object)) {
	t.o.Call("captureVisibleTab", windowId, options, callback)
}

// ExecuteScript injects JavaScript code into a page. For details, see the programmatic
// injection section of the content scripts doc: (https://developer.chrome.com/extensions/content_scripts#pi)
func (t *Tabs) ExecuteScript(tabId int, details interface{}, callback func(js.Object)) {
	t.o.Call("executeScript", tabId, details, callback)
}

// InsertCss injects CSS into a page. For details, see the programmatic injection
// section of the content scripts doc. (https://developer.chrome.com/extensions/content_scripts#pi)
func (t *Tabs) InsertCss(tabId int, details interface{}, callback func(js.Object)) {
	t.o.Call("insertCss", tabId, details, callback)
}

/*
* Events
 */

// OnCreated is fired when a tab is created. Note that the tab's URL may not be set at the time
// this event fired, but you can listen to onUpdated events to be notified when a URL is set.
func (t *Tabs) OnCreated(callback func(map[string]interface{})) {
	t.o.Get("onCreated").Call("addListener", callback)
}

// OnUpdated fired when a tab is updated.
func (t *Tabs) OnUpdated(callback func(tabId int, changeInfo interface{}, tab map[string]interface{})) {
	t.o.Get("onUpdated").Call("addListener", callback)
}

// OnMoved fired when a tab is moved within a window. Only one move event is fired,
// representing the tab the user directly moved. Move events are not fired for the
// other tabs that must move in response. This event is not fired when a tab is moved between windows.
func (t *Tabs) OnMoved(callback func(tabId int, movedInfo interface{})) {
	t.o.Get("onMoved").Call("addListener", callback)
}

// OnActivated fires when the active tab in a window changes. Note that the tab's URL
// may not be set at the time this event fired, but you can listen to onUpdated events
// to be notified when a URL is set.
func (t *Tabs) OnActivated(callback func(activeInfo interface{})) {
	t.o.Get("onActivated").Call("addListener", callback)
}

// OnHighlightChanged fired when the highlighted or selected tabs in a window changes.
func (t *Tabs) OnHighlightChanged(callback func(selectInfo interface{})) {
	t.o.Get("onHighlightChanged").Call("addListener", callback)
}

// OnHighlighted fired when the highlighted or selected tabs in a window changes.
func (t *Tabs) OnHighlighted(callback func(highlightInfo interface{})) {
	t.o.Get("onHighlighted").Call("addListener", callback)
}

// OnDetached fired when a tab is detached from a window, for example because it is being moved between windows.
func (t *Tabs) OnDetached(callback func(tabId int, detachInfo interface{})) {
	t.o.Get("onDetached").Call("addListener", callback)
}

// OnAttached fired when a tab is attached to a window, for example because it was moved between windows.
func (t *Tabs) OnAttached(callback func(tabId int, attachInfo interface{})) {
	t.o.Get("onAttached").Call("addListener", callback)
}

// OnRemoved fired when a tab is closed.
func (t *Tabs) OnRemoved(callback func(tabId int, removeInfo interface{})) {
	t.o.Get("onRemoved").Call("addListener", callback)
}

// OnReplaced fired when a tab is replaced with another tab due to prerendering or instant.
func (t *Tabs) OnReplaced(callback func(addedTabId int, removedTabId int)) {
	t.o.Get("OnReplaced").Call("addListener", callback)
}

// OnZoomChange fired when a tab is zoomed.
func (t *Tabs) OnZoomChange(callback func(zoomChangeInfo interface{})) {
	t.o.Get("onZoomChange").Call("addListener", callback)
}
