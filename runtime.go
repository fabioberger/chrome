package chrome

import "github.com/gopherjs/gopherjs/js"

type Runtime struct {
	O js.Object
}

func (r *Runtime) OnMessage(callback func(message interface{}, sender js.Object, sendResponse func(interface{}))) {
	r.O.Get("onMessage").Call("addListener", callback)
}
