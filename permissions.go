package chrome

import "github.com/gopherjs/gopherjs/js"

type Permissions struct {
	o *js.Object
}

/*
* Methods
 */

// GetAll gets the extension's current set of permissions.
func (p *Permissions) GetAll(callback func(permissions map[string][]string)) {
	p.o.Call("getAll", callback)
}

// Contains checks if the extension has the specified permissions.
func (p *Permissions) Contains(permissions map[string][]string, callback func(result bool)) {
	p.o.Call("contains", permissions, callback)
}

// Request requests access to the specified permissions. These permissions must be defined in the
// optional_permissions field of the manifest. If there are any problems requesting the permissions, runtime.lastError will be set.
func (p *Permissions) Request(permissions map[string][]string, callback func(granted bool)) {
	p.o.Call("request", permissions, callback)
}

// Remove removes access to the specified permissions. If there are any problems removing the permissions, runtime.lastError will be set.
func (p *Permissions) Remove(permissions map[string][]string, callback func(removed bool)) {
	p.o.Call("remove", permissions, callback)
}

/*
* Events
 */

// OnAdded fired when the extension acquires new permissions.
func (p *Permissions) OnAdded(callback func(permissions map[string][]string)) {
	p.o.Get("onAdded").Call("addListener", callback)
}

// OnRemoved fired when access to permissions has been removed from the extension.
func (p *Permissions) OnRemoved(callback func(permissions map[string][]string)) {
	p.o.Get("onRemoved").Call("addListener", callback)
}
