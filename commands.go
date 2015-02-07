package chrome

import "github.com/gopherjs/gopherjs/js"

type Commands struct {
	o js.Object
}

/*
* Methods:
 */

// GetAll returns all the registered extension commands for this extension and their shortcut (if active).
func (c *Commands) GetAll(callback func(commands []map[string]interface{})) {
	c.o.Call("getAll", callback)
}

/*
* Events
 */

// OnCommand fired when a registered command is activated using a keyboard shortcut.
func (c *Commands) OnCommand(callback func(command string)) {
	c.o.Get("onCommand").Call("addListener", callback)
}
