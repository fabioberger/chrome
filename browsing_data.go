package chrome

import "github.com/gopherjs/gopherjs/js"

type BrowsingData struct {
	o *js.Object
}

/*
* Types
 */

type RemovalOptions struct {
	js.Object
	Since       float64         `js:"since"`
	OriginTypes map[string]bool `js:"originTypes"`
}

type DataTypeSet map[string]bool

/*
* Methods:
 */

// Settings reports which types of data are currently selected in the 'Clear browsing data'
// settings UI. Note: some of the data types included in this API are not available in the
// settings UI, and some UI settings control more than one data type listed here.
func (b *BrowsingData) Settings(callback func(result Object)) {
	b.o.Call("settings", callback)
}

// Remove clears various types of browsing data stored in a user's profile.
func (b *BrowsingData) Remove(options RemovalOptions, dataToRemove DataTypeSet, callback func()) {
	b.o.Call("remove", options, dataToRemove, callback)
}

// RemoveAppCache clears websites' appcache data.
func (b *BrowsingData) RemoveAppCache(options RemovalOptions, callback func()) {
	b.o.Call("removeAppCache", options, callback)
}

// RemoveCache clears the browser's cache.
func (b *BrowsingData) RemoveCache(options RemovalOptions, callback func()) {
	b.o.Call("removeCache", options, callback)
}

// RemoveCookies clears the browser's cookies and server-bound certificates modified within a particular timeframe.
func (b *BrowsingData) RemoveCookies(options RemovalOptions, callback func()) {
	b.o.Call("removeCookies", options, callback)
}

// RemoveDownloads clears the browser's list of downloaded files (not the downloaded files themselves).
func (b *BrowsingData) RemoveDownloads(options RemovalOptions, callback func()) {
	b.o.Call("removeDownloads", options, callback)
}

// RemoveFileSystems clears websites' file system data.
func (b *BrowsingData) RemoveFileSystems(options RemovalOptions, callback func()) {
	b.o.Call("removeFileSystems", options, callback)
}

// RemoveFormData clears the browser's stored form data (autofill).
func (b *BrowsingData) RemoveFormData(options RemovalOptions, callback func()) {
	b.o.Call("removeFormData", options, callback)
}

// RemoveHistory clears the browser's history.
func (b *BrowsingData) RemoveHistory(options RemovalOptions, callback func()) {
	b.o.Call("removeHistory", options, callback)
}

// RemoveIndexedDB clears websites' IndexedDB data.
func (b *BrowsingData) RemoveIndexedDB(options RemovalOptions, callback func()) {
	b.o.Call("removeIndexedDB", options, callback)
}

// RemoveLocalStorage clears websites' local storage data.
func (b *BrowsingData) RemoveLocalStorage(options RemovalOptions, callback func()) {
	b.o.Call("removeLocalStorage", options, callback)
}

// RemovePluginData clears plugins' data.
func (b *BrowsingData) RemovePluginData(options RemovalOptions, callback func()) {
	b.o.Call("removePluginData", options, callback)
}

// RemovePasswords clears the browser's stored passwords.
func (b *BrowsingData) RemovePasswords(options RemovalOptions, callback func()) {
	b.o.Call("removePasswords", options, callback)
}

// RemoveWebSQL clears websites' WebSQL data.
func (b *BrowsingData) RemoveWebSQL(options RemovalOptions, callback func()) {
	b.o.Call("removeWebSQL", options, callback)
}
