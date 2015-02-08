package chrome

import "github.com/gopherjs/gopherjs/js"

type Extension struct {
	o js.Object
}

/*
* Methods:
 */

// GetURL converts a relative path within an extension install directory to a fully-qualified URL.
func (e *Extension) GetURL(path string) {
	e.o.Call("getURL", path)
}

// GetViews returns an array of the JavaScript 'window' objects for each of the pages running inside the current extension.
func (e *Extension) GetViews(fetchProperties map[string]interface{}) interface{} {
	return e.o.Call("getViews", fetchProperties).Interface()
}

// GetBackgroundPage returns the JavaScript 'window' object for the background page running inside
// the current extension. Returns null if the extension has no background page.
func (e *Extension) GetBackgroundPage() interface{} {
	return e.o.Call("getBackgroundPage").Interface()
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
