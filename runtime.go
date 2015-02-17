package chrome

import "github.com/gopherjs/gopherjs/js"

type Runtime struct {
	o         js.Object
	LastError map[string]string
	Id        string
}

func NewRuntime(runtimeObj js.Object) *Runtime {
	r := new(Runtime)
	r.o = runtimeObj
	if r.o.String() != "undefined" {
		r.LastError = r.o.Get("lastError").Interface().(map[string]string)
		r.Id = r.o.Get("id").String()
	}
	return r
}

/*
* Types
 */

type Port struct {
	js.Object
	Name         string        `js:"name"`
	OnDisconnect js.Object     `js:"onDisconnect"`
	OnMessage    js.Object     `js:"onMessage"`
	Sender       MessageSender `js:"sender"`
}

type MessageSender struct {
	js.Object
	tab          Tab    `js:"tab"`
	FrameId      int    `js:"frameId"`
	Id           string `js:"id"`
	Url          string `js:"url"`
	TlsChannelId string `js:"tlsChannelId"`
}

type PlatformInfo map[string]string

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

func (r *Runtime) Connect(extensionId string, connectInfo interface{}) Port {
	return r.o.Call("connect", extensionId, connectInfo).(Port)
}

func (r *Runtime) ConnectNative(application string) Port {
	return r.o.Call("connectNative", application).(Port)
}

func (r *Runtime) SendMessage(extensionId string, message interface{}, options interface{}, responseCallback func(response interface{})) {
	r.o.Call("sendMessage", extensionId, message, options, responseCallback)
}

func (r *Runtime) SendNativeMessage(application string, message interface{}, responseCallback func(response interface{})) {
	r.o.Call("sendNativeMessage", application, message, responseCallback)
}

func (r *Runtime) GetPlatformInfo(callback func(platformInfo PlatformInfo)) {
	r.o.Call("getPlatformInfo", callback)
}

func (r *Runtime) GetPackageDirectoryEntry(callback func(directoryEntry interface{})) {
	r.o.Call("getPackageDirectoryEntry", callback)
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

func (r *Runtime) OnConnect(callback func(port Port)) {
	r.o.Get("onConnect").Call("addListener", callback)
}

func (r *Runtime) OnConnectExternal(callback func(port Port)) {
	r.o.Get("onConnectExternal").Call("addListener", callback)
}

func (r *Runtime) OnMessage(callback func(message interface{}, sender MessageSender, sendResponse func(interface{}))) {
	r.o.Get("onMessage").Call("addListener", callback)
}

func (r *Runtime) OnMessageExternal(callback func(message interface{}, sender MessageSender, sendResponse func(interface{}))) {
	r.o.Get("onMessageExternal").Call("addListener", callback)
}

func (r *Runtime) OnRestartRequired(callback func(reason string)) {
	r.o.Get("onRestartRequired").Call("addListener", callback)
}
