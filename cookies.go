package chrome

import "github.com/gopherjs/gopherjs/js"

type Cookies struct {
	o js.Object
}

/*
* Methods:
 */

// Get retrieves information about a single cookie. If more than one cookie of the same name
// exists for the given URL, the one with the longest path will be returned. For cookies with
// the same path length, the cookie with the earliest creation time will be returned.
func (c *Cookies) Get(details map[string]interface{}, callback func(cookie map[string]interface{})) {
	c.o.Call("get", details, callback)
}

// GetAll retrieves all cookies from a single cookie store that match the given information.
// The cookies returned will be sorted, with those with the longest path first. If multiple cookies
// have the same path length, those with the earliest creation time will be first.
func (c *Cookies) GetAll(details map[string]interface{}, callback func(cookies []map[string]interface{})) {
	c.o.Call("getAll", details, callback)
}

// Set sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
func (c *Cookies) Set(details map[string]interface{}, callback func(cookie map[string]interface{})) {
	c.o.Call("set", details, callback)
}

// Remove deletes a cookie by name.
func (c *Cookies) Remove(details map[string]interface{}, callback func(details map[string]interface{})) {
	c.o.Call("remove", details, callback)
}

// GetAllCookieStores lists all existing cookie stores.
func (c *Cookies) GetAllCookieStores(callback func(cookieStores []map[string]interface{})) {
	c.o.Call("getAllCookieStores", callback)
}

/*
* Events
 */

// OnChanged fired when a cookie is set or removed. As a special case, note that updating a cookie's properties
//  is implemented as a two step process: the cookie to be updated is first removed entirely, generating a
// notification with "cause" of "overwrite" . Afterwards, a new cookie is written with the updated values,
// generating a second notification with "cause" "explicit".
func (c *Cookies) OnChanged(callback func(changeInfo map[string]interface{})) {
	c.o.Get("onChanged").Call("addListener", callback)
}
