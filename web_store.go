package chrome

import "github.com/gopherjs/gopherjs/js"

type WebStore struct {
	o js.Object
}

/*
* Methods
 */

// Install instals an app from the chrome store
func (w *WebStore) Install(url string, successCallback func(), failureCallback func(err string, errorCode string)) {
	w.o.Call("install", url, successCallback, failureCallback)
}

/*
* Events
 */

// OnInstallStageChanged fired when an inline installation enters a new InstallStage. In order to receive
// notifications about this event, listeners must be registered before the inline installation begins.
func (w *WebStore) OnInstallStageChanged(callback func(stage string)) {
	w.o.Get("onInstallStageChanged").Call("addListener", callback)
}

// OnDownloadProgress fired periodically with the download progress of an inline install. In order to
// receive notifications about this event, listeners must be registered before the inline installation begins.
func (w *WebStore) OnDownloadProgress(callback func(percentDownloaded int64)) {
	w.o.Get("onDownloadProgress").Call("addListener", callback)
}
