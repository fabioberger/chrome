package chrome

import "github.com/gopherjs/gopherjs/js"

type Bookmarks struct {
	o js.Object
}

/*
* Methods:
 */

func (b *Bookmarks) Get(idList []string, callback func(results []map[string]interface{})) {
	b.o.Call("get", idList, callback)
}

func (b *Bookmarks) GetChildren(id string, callback func(results []map[string]interface{})) {
	b.o.Call("getChildren", id, callback)
}

func (b *Bookmarks) GetRecent(numberOfItems int, callback func(results []map[string]interface{})) {
	b.o.Call("getRecent", numberOfItems, callback)
}

func (b *Bookmarks) GetTree(callback func(results []map[string]interface{})) {
	b.o.Call("getTree", callback)
}

func (b *Bookmarks) GetSubTree(id string, callback func(results []map[string]interface{})) {
	b.o.Call("getSubTree", id, callback)
}

func (b *Bookmarks) Search(query interface{}, callback func(results []map[string]interface{})) {
	b.o.Call("search", query, callback)
}

func (b *Bookmarks) Create(bookmark interface{}, callback func(result map[string]interface{})) {
	b.o.Call("create", bookmark, callback)
}

func (b *Bookmarks) Move(id string, destination interface{}, callback func(result map[string]interface{})) {
	b.o.Call("move", id, destination, callback)
}

func (b *Bookmarks) Update(id string, changes interface{}, callback func(result map[string]interface{})) {
	b.o.Call("update", id, changes, callback)
}

func (b *Bookmarks) Remove(id string, callback func()) {
	b.o.Call("remove", id, callback)
}

func (b *Bookmarks) RemoveTree(id string, callback func()) {
	b.o.Call("removeTree", id, callback)
}

/*
* Events
 */

func (b *Bookmarks) OnCreated(callback func(id string, bookmark map[string]interface{})) {
	t.o.Get("onCreated").Call("addListener", callback)
}

func (b *Bookmarks) OnRemoved(callback func(id string, removeInfo map[string]interface{})) {
	t.o.Get("onRemoved").Call("addListener", callback)
}

func (b *Bookmarks) onChanged(callback func(id string, changeInfo map[string]interface{})) {
	t.o.Get("onChanged").Call("addListener", callback)
}

func (b *Bookmarks) OnMoved(callback func(id string, moveInfo map[string]interface{})) {
	t.o.Get("onMoved").Call("addListener", callback)
}

func (b *Bookmarks) OnChildrenReordered(callback func(id string, reorderInfo map[string]interface{})) {
	t.o.Get("onChildrenReordered").Call("addListener", callback)
}

func (b *Bookmarks) OnImportBegan(callback func()) {
	t.o.Get("onImportBegan").Call("addListener", callback)
}

func (b *Bookmarks) OnImportEnded(callback func()) {
	t.o.Get("onImportEnded").Call("addListener", callback)
}
