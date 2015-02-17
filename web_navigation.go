package chrome

import "github.com/gopherjs/gopherjs/js"

type WebNavigation struct {
	o js.Object
}

/*
* Methods
 */

// GetFrame retrieves information about the given frame. A frame refers to an <iframe> or
// a <frame> of a web page and is identified by a tab ID and a frame ID.
func (w *WebNavigation) GetFrame(details map[string]interface{}, callback func(details map[string]interface{})) {
	w.o.Call("getFrame", details, callback)
}

// GetAllFrames retrieves information about all frames of a given tab.
func (w *WebNavigation) GetAllFrames(details map[string]interface{}, callback func(details []map[string]interface{})) {
	w.o.Call("getAllFrames", details, callback)
}

/*
* Events
 */

// OnBeforeNavigate fired when a navigation is about to occur.
func (w *WebNavigation) OnBeforeNavigate(callback func(details map[string]interface{}), filters map[string][]map[string]interface{}) {
	w.o.Get("onBeforeNavigate").Call("addListener", callback, filters)
}

// OnCommited fired when a navigation is committed. The document (and the resources it refers to,
// such as images and subframes) might still be downloading, but at least part of the document has
// been received from the server and the browser has decided to switch to the new document.
func (w *WebNavigation) OnCommited(callback func(details map[string]interface{}), filters map[string][]map[string]interface{}) {
	w.o.Get("onCommited").Call("addListener", callback, filters)
}

// OnDOMContentLoaded fired when the page's DOM is fully constructed, but the referenced resources may not finish loading.
func (w *WebNavigation) OnDOMContentLoaded(callback func(details map[string]interface{}), filters map[string][]map[string]interface{}) {
	w.o.Get("onDOMContentLoaded").Call("addListener", callback, filters)
}

// OnCompleted fired when a document, including the resources it refers to, is completely loaded and initialized.
func (w *WebNavigation) OnCompleted(callback func(details map[string]interface{}), filters map[string][]map[string]interface{}) {
	w.o.Get("onCompleted").Call("addListener", callback, filters)
}

// OnErrorOccurred fired when an error occurs and the navigation is aborted. This can happen if either a network
// error occurred, or the user aborted the navigation.
func (w *WebNavigation) OnErrorOccurred(callback func(details map[string]interface{}), filters map[string][]map[string]interface{}) {
	w.o.Get("onErrorOccurred").Call("addListener", callback, filters)
}

// OnCreatedNavigationTarget fired when a new window, or a new tab in an existing window, is created to host a navigation.
func (w *WebNavigation) OnCreatedNavigationTarget(callback func(details map[string]interface{}), filters map[string][]map[string]interface{}) {
	w.o.Get("onCreatedNavigationTarget").Call("addListener", callback, filters)
}

// OnReferenceFragmentUpdated fired when the reference fragment of a frame was updated. All future events for that frame will
// use the updated URL.
func (w *WebNavigation) OnReferenceFragmentUpdated(callback func(details map[string]interface{}), filters map[string][]map[string]interface{}) {
	w.o.Get("onReferenceFragmentUpdated").Call("addListener", callback, filters)
}

// OnTabReplaced fired when the contents of the tab is replaced by a different (usually previously pre-rendered) tab.
func (w *WebNavigation) OnTabReplaced(callback func(details map[string]interface{})) {
	w.o.Get("onTabReplaced").Call("addListener", callback)
}

// OnHistoryStateUpdated fired when the frame's history was updated to a new URL. All future events for that frame will use the updated URL.
func (w *WebNavigation) OnHistoryStateUpdated(callback func(details map[string]interface{}), filters map[string][]map[string]interface{}) {
	w.o.Get("onHistoryStateUpdated").Call("addListener", callback, filters)
}
