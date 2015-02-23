package chrome

import "github.com/gopherjs/gopherjs/js"

type ContextMenus struct {
	o                           *js.Object
	ACTION_MENU_TOP_LEVEL_LIMIT int
}

func NewContextMenus(contextMenusObj *js.Object) *ContextMenus {
	c := new(ContextMenus)
	c.o = contextMenusObj
	if c.o.String() != "undefined" {
		c.ACTION_MENU_TOP_LEVEL_LIMIT = contextMenusObj.Get("ACTION_MENU_TOP_LEVEL_LIMIT").Int()
	}
	return c
}

/*
* Methods:
 */

// Create creates a new context menu item. Note that if an error occurs during creation,
// you may not find out until the creation callback fires (the details will be in chrome.Runtime.LastError).
func (c *ContextMenus) Create(createProperties Object, callback func()) {
	c.o.Call("create", createProperties, callback)
}

// Update updates a previously created context menu item.
func (c *ContextMenus) Update(id interface{}, updateProperties Object, callback func()) {
	c.o.Call("update", id, updateProperties, callback)
}

// Remove removes a context menu item.
func (c *ContextMenus) Remove(menuItemId interface{}, callback func()) {
	c.o.Call("remove", menuItemId, callback)
}

// RemoveAll removes all context menu items added by this extension.
func (c *ContextMenus) RemoveAll(callback func()) {
	c.o.Call("removeAll", callback)
}

/*
* Events
 */

// OnClicked fired when a context menu item is clicked.
func (c *ContextMenus) OnClicked(callback func(info Object, tab Tab)) {
	c.o.Get("onClicked").Call("addListener", callback)
}
