package chrome

import "github.com/gopherjs/gopherjs/js"

type Power struct {
	o js.Object
}

/*
* Methods
 */

// RequestKeepAwake requests that power management be temporarily disabled. |level| describes the
//degree to which power management should be disabled. If a request previously made by the same
//app is still active, it will be replaced by the new request.
func (p *Power) RequestKeepAwake(level string) {
	p.o.Call("requestKeepAwake", level)
}

// ReleaseKeepAwake releases a request previously made via requestKeepAwake().
func (p *Power) ReleaseKeepAwake() {
	p.o.Call("releaseKeepAwake")
}
