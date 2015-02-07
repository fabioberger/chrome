package chrome

import "github.com/gopherjs/gopherjs/js"

type Debugger struct {
	o js.Object
}

/*
* Methods:
 */

// Attach attaches debugger to the given target.
func (d *Debugger) Attach(target map[string]interface{}, requiredVersion string, callback func()) {
	d.o.Call("attach", target, requiredVersion, callback)
}

// Detach detaches debugger from the given target.
func (d *Debugger) Detach(target map[string]interface{}, callback func()) {
	d.o.Call("detach", target, callback)
}

// SendCommand sends given command to the debugging target.
func (d *Debugger) SendCommand(target map[string]interface{}, method string, commandParams map[string]interface{}, callback func(result map[string]interface{})) {
	d.o.Call("sendCommand", target, method, commandParams, callback)
}

// GetTargets returns the list of available debug targets.
func (d *Debugger) GetTargets(callback func(result []map[string]interface{})) {
	d.o.Call("getTargets", callback)
}

/*
* Events
 */

// OnEvent fired whenever debugging target issues instrumentation event.
func (d *Debugger) OnEvent(callback func(source map[string]interface{}, method string, params map[string]interface{})) {
	d.o.Get("onEvent").Call("addListener", callback)
}

// OnDetach fired when browser terminates debugging session for the tab. This happens when
// either the tab is being closed or Chrome DevTools is being invoked for the attached tab.
func (d *Debugger) OnDetach(callback func(source map[string]interface{}, reason string)) {
	d.o.Get("onDetach").Call("addListener", callback)
}
