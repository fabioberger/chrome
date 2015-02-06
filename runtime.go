package chrome

import "github.com/gopherjs/gopherjs/js"

type Runtime struct {
	o js.Object
}

/*
* Methods
 */

func (r *Runtime) GetBackgroundPage(callback func(backgroundPage interface{})) {
	r.o.Call("getBackgroundPage", callback)
}

func (r *Runtime) GetManifest() js.Object {
	return r.o.Call("getManifest")
}

func (r *Runtime) GetURL(path string) string {
	return r.o.Call("getURL", path).String()
}

func (r *Runtime) Reload() {
	r.o.Call("reload")
}

// Maybe make status an Enum string with specific values.
func (r *Runtime) RequestUpdateCheck(callback func(status string, details interface{})) {
	r.o.Call("requestUpdateCheck", callback)
}

func (r *Runtime) Restart() {
	r.o.Call("restart")
}

func (r *Runtime) Connect(extensionId string, connectInfo interface{}) {
	r.o.Call("connect", extensionId, connectInfo)
}

func (r *Runtime) ConnectNative(application string) {
	r.o.Call("connectNative", application)
}

func (r *Runtime) SendMessage(extensionId string, message interface{}, options interface{}, responseCallback func(response interface{})) {
	r.o.Call("sendMessage", extensionId, message, options, responseCallback)
}

func (r *Runtime) SendNativeMessage(application string, message interface{}, responseCallback func(response interface{})) {
	r.o.Call("sendNativeMessage", application, message, responseCallback)
}

func (r *Runtime) GetPlatformInfo(callback func(platformInfo map[string]string)) {
	r.o.Call("getPlatformInfo", callback)
}

func (r *Runtime) GetPackageDirectoryEntry(callback func(directoryEntry interface{})) {
	r.o.Call("getPackageDirectoryEntry", directoryEntry)
}

/*
* Events
 */

func (r *Runtime) OnStartup(callback func()) {
	r.o.Get("onStartup").Call("addListener", callback)
}

func (r *Runtime) OnInstalled(callback func(details map[string]string)) {
	r.o.Get("onInstalled").Call("addListener", callback)
}

func (r *Runtime) OnSuspend(callback func()) {
	r.o.Get("onSuspend").Call("addListener", callback)
}

func (r *Runtime) OnSuspendCanceled(callback func()) {
	r.o.Get("onSuspendCanceled").Call("addListener", callback)
}

func (r *Runtime) OnUpdateAvailable(callback func(details map[string]string)) {
	r.o.Get("onUpdateAvailable").Call("addListener", callback)
}

func (r *Runtime) OnConnect(callback func(port interface{})) {
	r.o.Get("onConnect").Call("addListener", callback)
}

func (r *Runtime) OnConnectExternal(callback func(port interface{})) {
	r.o.Get("onConnectExternal").Call("addListener", callback)
}

func (r *Runtime) OnMessage(callback func(message interface{}, sender js.Object, sendResponse func(interface{}))) {
	r.o.Get("onMessage").Call("addListener", callback)
}

func (r *Runtime) OnMessageExternal(callback func(message interface{}, sender js.Object, sendResponse func(interface{}))) {
	r.o.Get("onMessageExternal").Call("addListener", callback)
}

func (r *Runtime) OnRestartRequired(callback func(reason string)) {
	r.o.Get("onRestartRequired").Call("addListener", callback)
}
