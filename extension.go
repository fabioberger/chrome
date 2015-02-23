package chrome

import "github.com/gopherjs/gopherjs/js"

type Extension struct {
	o                  *js.Object
	LastError          *js.Object
	InIncognitoContext bool
}

func NewExtension(extensionObj *js.Object) *Extension {
	e := new(Extension)
	e.o = extensionObj
	if extensionObj.String() != "undefined" {
		e.LastError = e.o.Get("lastError")
		e.InIncognitoContext = e.o.Get("inIncognitoContext").Bool()
	}
	return e
}

/*
* Methods:
 */

// GetURL converts a relative path within an extension install directory to a fully-qualified URL.
func (e *Extension) GetURL(path string) {
	e.o.Call("getURL", path)
}

// GetViews returns an array of the JavaScript 'window' objects for each of the pages running inside the current extension.
// Fix this and the other functions to return Window objects instead of js.Object or whatever else
func (e *Extension) GetViews(fetchProperties Object) []Window {
	windows := []Window{}
	windowObjs := e.o.Call("getViews", fetchProperties)
	for i := 0; i < windowObjs.Length(); i++ {
		window := windowObjs.Index(i)
		windows = append(windows, Window{Object: *window})
	}
	return windows
}

// GetBackgroundPage returns the JavaScript 'window' object for the background page running inside
// the current extension. Returns null if the extension has no background page.
func (e *Extension) GetBackgroundPage() Window {
	return Window{Object: *e.o.Call("getBackgroundPage")}
}

// GetExtensionTabs returns an array of the JavaScript 'window' objects for each of the tabs running inside
// the current extension. If windowId is specified, returns only the 'window' objects of tabs attached to the specified window.
func (e *Extension) GetExtensionTabs(windowId int) []Window {
	windows := []Window{}
	windowObjs := e.o.Call("getExtensionTabs", windowId)
	for i := 0; i < windowObjs.Length(); i++ {
		window := windowObjs.Index(i)
		windows = append(windows, Window{Object: *window})
	}
	return windows
}

// IsAllowedIncognitoAccess retrieves the state of the extension's access to Incognito-mode
// (as determined by the user-controlled 'Allowed in Incognito' checkbox.
func (e *Extension) IsAllowedIncognitoAccess(callback func(isAllowedAccess bool)) {
	e.o.Call("isAllowedIncognitoAccess", callback)
}

// IsAllowedFileSchemeAccess retrieves the state of the extension's access to the 'file://'
// scheme (as determined by the user-controlled 'Allow access to File URLs' checkbox.
func (e *Extension) IsAllowedFileSchemeAccess(callback func(isAllowedAccess bool)) {
	e.o.Call("isAllowedFileSchemeAccess", callback)
}

// SetUpdateUrlData sets the value of the ap CGI parameter used in the extension's update URL.
// This value is ignored for extensions that are hosted in the Chrome Extension Gallery.
func (e *Extension) SetUpdateUrlData(data string) {
	e.o.Call("setUpdateUrlData", data)
}
