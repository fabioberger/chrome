package chrome

import "github.com/gopherjs/gopherjs/js"

type Idle struct {
	o *js.Object
}

/*
* Methods
 */

// QueryState returns "locked" if the system is locked, "idle" if the user has not generated any
// input for a specified number of seconds, or "active" otherwise.
func (i *Idle) QueryState(detectionIntervalInSeconds int, callback func(newState string)) {
	i.o.Call("queryState", detectionIntervalInSeconds, callback)
}

// SetDetectionInterval sets the interval, in seconds, used to determine when the system is in an idle
//  state for onStateChanged events. The default interval is 60 seconds.
func (i *Idle) SetDetectionInterval(intervalInSeconds int) {
	i.o.Call("setDetectionInterval", intervalInSeconds)
}

/*
* Events
 */

// OnStateChanged fired when the system changes to an active, idle or locked state. The event fires with
// "locked" if the screen is locked or the screensaver activates, "idle" if the system is unlocked and the
// user has not generated any input for a specified number of seconds, and "active" when the user generates
// input on an idle system.
func (i *Idle) OnStateChanged(callback func(newState string)) {
	i.o.Get("onStateChanged").Call("addListener", callback)
}
