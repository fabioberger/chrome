package chrome

import "github.com/gopherjs/gopherjs/js"

type Bookmarks struct {
	o js.Object
}

/*
* Types
 */

type BookmarkTreeNode struct {
	js.Object
	Id                string             `js:"id"`
	ParentId          string             `js:"parentId,omitempty"`
	Index             int                `js:"index,omitempty"`
	Url               string             `js:"url,omitempty"`
	Title             string             `js:"title"`
	DateAdded         float64            `js:"dateAdded,omitempty"`
	DateGroupModified float64            `js:"dateGroupModified,omitempty"`
	Unmodifiable      string             `js:"unmodifiable,omitempty"`
	Children          []BookmarkTreeNode `js:"children,omitempty"`
}

/*
* Methods:
 */

// Get retrieves the specified BookmarkTreeNode(s).
func (b *Bookmarks) Get(idList []string, callback func(results []BookmarkTreeNode)) {
	b.o.Call("get", idList, callback)
}

// GetChildren retrieves the children of the specified BookmarkTreeNode id.
func (b *Bookmarks) GetChildren(id string, callback func(results []BookmarkTreeNode)) {
	b.o.Call("getChildren", id, callback)
}

// GetRecent retrieves the recently added bookmarks.
func (b *Bookmarks) GetRecent(numberOfItems int, callback func(results []BookmarkTreeNode)) {
	b.o.Call("getRecent", numberOfItems, callback)
}

// GetTree retrieves the entire Bookmarks hierarchy.
func (b *Bookmarks) GetTree(callback func(results []BookmarkTreeNode)) {
	b.o.Call("getTree", callback)
}

// GetSubTree retrieves part of the Bookmarks hierarchy, starting at the specified node.
func (b *Bookmarks) GetSubTree(id string, callback func(results []BookmarkTreeNode)) {
	b.o.Call("getSubTree", id, callback)
}

// Search searches for BookmarkTreeNodes matching the given query. Queries specified
// with an object produce BookmarkTreeNodes matching all specified properties.
func (b *Bookmarks) Search(query interface{}, callback func(results []BookmarkTreeNode)) {
	b.o.Call("search", query, callback)
}

// Create creates a bookmark or folder under the specified parentId.
// If url is nil or missing, it will be a folder.
func (b *Bookmarks) Create(bookmark interface{}, callback func(result BookmarkTreeNode)) {
	b.o.Call("create", bookmark, callback)
}

// Move moves the specified BookmarkTreeNode to the provided location.
func (b *Bookmarks) Move(id string, destination interface{}, callback func(result BookmarkTreeNode)) {
	b.o.Call("move", id, destination, callback)
}

// Update updates the properties of a bookmark or folder. Specify only the properties that you want
// to change; unspecified properties will be left unchanged. Note: Currently, only 'title' and 'url' are supported.
func (b *Bookmarks) Update(id string, changes interface{}, callback func(result BookmarkTreeNode)) {
	b.o.Call("update", id, changes, callback)
}

// Remove removes a bookmark or an empty bookmark folder.
func (b *Bookmarks) Remove(id string, callback func()) {
	b.o.Call("remove", id, callback)
}

// RemoveTree recursively removes a bookmark folder.
func (b *Bookmarks) RemoveTree(id string, callback func()) {
	b.o.Call("removeTree", id, callback)
}

/*
* Events
 */

// OnCreated fired when a bookmark or folder is created.
func (b *Bookmarks) OnCreated(callback func(id string, bookmark BookmarkTreeNode)) {
	b.o.Get("onCreated").Call("addListener", callback)
}

// OnRemoved fired when a bookmark or folder is removed. When a folder is removed recursively,
// a single notification is fired for the folder, and none for its contents.
func (b *Bookmarks) OnRemoved(callback func(id string, removeInfo map[string]interface{})) {
	b.o.Get("onRemoved").Call("addListener", callback)
}

// onChanged fired when a bookmark or folder changes. Note: Currently, only title and url changes trigger this.
func (b *Bookmarks) onChanged(callback func(id string, changeInfo map[string]interface{})) {
	b.o.Get("onChanged").Call("addListener", callback)
}

// OnMoved fired when a bookmark or folder is moved to a different parent folder.
func (b *Bookmarks) OnMoved(callback func(id string, moveInfo map[string]interface{})) {
	b.o.Get("onMoved").Call("addListener", callback)
}

// OnChildrenReordered fired when the children of a folder have changed their order due to
// the order being sorted in the UI. This is not called as a result of a move().
func (b *Bookmarks) OnChildrenReordered(callback func(id string, reorderInfo map[string]interface{})) {
	b.o.Get("onChildrenReordered").Call("addListener", callback)
}

// OnImportBegan fired when a bookmark import session is begun. Expensive observers should ignore
// onCreated updates until onImportEnded is fired. Observers should still handle other notifications immediately.
func (b *Bookmarks) OnImportBegan(callback func()) {
	b.o.Get("onImportBegan").Call("addListener", callback)
}

// OnImportEnded fired when a bookmark import session is ended.
func (b *Bookmarks) OnImportEnded(callback func()) {
	b.o.Get("onImportEnded").Call("addListener", callback)
}
