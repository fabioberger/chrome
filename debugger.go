package chrome

import "github.com/gopherjs/gopherjs/js"

type Debugger struct {
	o js.Object
}

/*
* Types
 */

type Debugee struct {
	js.Object
	TabId       int    `js:"tabId"`
	ExtensionId string `js:"extensionId"`
	TargetId    string `js:"targetId"`
}

type TargetInfo struct {
	js.Object
	Type        string `js:"type"`
	Id          string `js:"id"`
	TabId       int    `js:"tabId"`
	ExtensionId string `js:"extensionId"`
	Attached    bool   `js:"attached"`
	Title       string `js:"title"`
	Url         string `js:"url"`
	FaviconUrl  string `js:"faviconUrl"`
}

/*
* Methods:
 */

// Attach attaches debugger to the given target.
func (d *Debugger) Attach(target Debugee, requiredVersion string, callback func()) {
	d.o.Call("attach", target, requiredVersion, callback)
}

// Detach detaches debugger from the given target.
func (d *Debugger) Detach(target Debugee, callback func()) {
	d.o.Call("detach", target, callback)
}

// SendCommand sends given command to the debugging target.
func (d *Debugger) SendCommand(target Debugee, method string, commandParams Object, callback func(result Object)) {
	d.o.Call("sendCommand", target, method, commandParams, callback)
}

// GetTargets returns the list of available debug targets.
func (d *Debugger) GetTargets(callback func(result []TargetInfo)) {
	d.o.Call("getTargets", callback)
}

/*
* Events
 */

// OnEvent fired whenever debugging target issues instrumentation event.
func (d *Debugger) OnEvent(callback func(source Debugee, method string, params Object)) {
	d.o.Get("onEvent").Call("addListener", callback)
}

// OnDetach fired when browser terminates debugging session for the tab. This happens when
// either the tab is being closed or Chrome DevTools is being invoked for the attached tab.
func (d *Debugger) OnDetach(callback func(source Debugee, reason string)) {
	d.o.Get("onDetach").Call("addListener", callback)
}
