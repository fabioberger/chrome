package chrome

import "github.com/gopherjs/gopherjs/js"

type WebRequest struct {
	o                                                 js.Object
	MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES int
}

func NewWebRequest(webRequestObj js.Object) *WebRequest {
	w := new(WebRequest)
	w.o = webRequestObj
	if w.o.String() != "undefined" {
		w.MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES = w.o.Get("MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES").Int()
	}
	return w
}

/*
* Types
 */

type RequestFilter struct {
	js.Object
	Urls     []string `js:"urls"`
	Types    string   `js:"types"`
	TabId    int      `js:"tabId"`
	WindowId int      `js:"windowId"`
}

type BlockingResponse struct {
	js.Object
	Cancel          bool              `js:"cancel"`
	RedirectUrl     string            `js:"redirectUrl"`
	RequestHeaders  Object            `js:"requestHeaders"`
	ResponseHeaders Object            `js:"responseHeaders"`
	AuthCredentials map[string]string `js:"authCredentials"`
}

type UploadData struct {
	js.Object
	Bytes interface{} `js:"bytes"`
	File  string      `js:"file"`
}

/*
* Methods
 */

// HandlerBehaviorChanged needs to be called when the behavior of the webRequest handlers has
// changed to prevent incorrect handling due to caching. This function call is expensive. Don't call it often.
func (w *WebRequest) HandlerBehaviorChanged(callback func()) {
	w.o.Call("handlerBehaviorChanged", callback)
}

/*
* Events
 */

// OnBeforeRequest fired when a request is about to occur.
func (w *WebRequest) OnBeforeRequest(callback func(details Object)) {
	w.o.Get("onBeforeRequest").Call("addListener", callback)
}

// OnBeforeSendHeaders fired before sending an HTTP request, once the request headers are available.
// This may occur after a TCP connection is made to the server, but before any HTTP data is sent.
func (w *WebRequest) OnBeforeSendHeaders(callback func(details Object)) {
	w.o.Get("onBeforeSendHeaders").Call("addListener", callback)
}

// OnSendHeaders fired just before a request is going to be sent to the server (modifications of
// previous onBeforeSendHeaders callbacks are visible by the time onSendHeaders is fired).
func (w *WebRequest) OnSendHeaders(callback func(details Object)) {
	w.o.Get("onSendHeaders").Call("addListener", callback)
}

// OnHeadersReceived fired when HTTP response headers of a request have been received.
func (w *WebRequest) OnHeadersReceived(callback func(details Object)) {
	w.o.Get("onHeadersReceived").Call("addListener", callback)
}

// OnAuthRequired fired when an authentication failure is received. The listener has three options:
// it can provide authentication credentials, it can cancel the request and display the error page,
// or it can take no action on the challenge. If bad user credentials are provided, this may be
// called multiple times for the same request.
func (w *WebRequest) OnAuthRequired(callback func(details Object)) {
	w.o.Get("onAuthRequired").Call("addListener", callback)
}

// OnResponseStarted fired when the first byte of the response body is received. For HTTP requests,
// this means that the status line and response headers are available.
func (w *WebRequest) OnResponseStarted(callback func(details Object)) {
	w.o.Get("onResponseStarted").Call("addListener", callback)
}

// OnBeforeRedirect fired when a server-initiated redirect is about to occur.
func (w *WebRequest) OnBeforeRedirect(callback func(details Object)) {
	w.o.Get("onBeforeRedirect").Call("addListener", callback)
}

// OnCompleted fired when a request is completed.
func (w *WebRequest) OnCompleted(callback func(details Object)) {
	w.o.Get("onCompleted").Call("addListener", callback)
}

// OnErrorOccured fired when an error occurs.
func (w *WebRequest) OnErrorOccured(callback func(details Object)) {
	w.o.Get("onErrorOccured").Call("addListener", callback)
}
