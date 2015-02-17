package chrome

import "github.com/gopherjs/gopherjs/js"

type TabCapture struct {
	o js.Object
}

/*
* Types
 */

type CaptureInfo struct {
	js.Object
	TabId      int    `js:"tabId"`
	Status     string `js:"status"`
	FullScreen bool   `js:"fullScreen"`
}

type MediaStreamConstraint struct {
	Mandatory js.Object `js:"mandatory"`
	Optional  js.Object `js:"optional"`
}

/*
* Methods
 */

// Capture captures the visible area of the currently active tab. Capture can only be started on
// the currently active tab after the extension has been invoked. Capture is maintained across page
// navigations within the tab, and stops when the tab is closed, or the media stream is closed by the extension.
func (t *TabCapture) Capture(options Object, callback func(stream interface{})) {
	t.o.Call("capture", options, callback)
}

// GetCapturedTabs returns a list of tabs that have requested capture or are being captured, i.e. status
// != stopped and status != error. This allows extensions to inform the user that there is an existing tab
// capture that would prevent a new tab capture from succeeding (or to prevent redundant requests for the same tab).
func (t *TabCapture) GetCapturedTabs(callback func(result []CaptureInfo)) {
	t.o.Call("getCapturedTabs", callback)
}

/*
* Events
 */

// OnStatusChanged event fired when the capture status of a tab changes. This allows extension authors to keep track
// of the capture status of tabs to keep UI elements like page actions and infobars in sync.
func (t *TabCapture) OnStatusChanged(callback func(info CaptureInfo)) {
	t.o.Get("onStatusChanged").Call("addListener", callback)
}
