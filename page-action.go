package chrome

import "github.com/gopherjs/gopherjs/js"

type PageAction struct {
	o *js.Object
}

/*
* Methods
 */

// Show shows the page action. The page action is shown whenever the tab is selected.
func (p *PageAction) Show(tabId int) {
	p.o.Call("show", tabId)
}

// Hide hides the page action.
func (p *PageAction) Hide(tabId int) {
	p.o.Call("hide", tabId)
}

// SetTitle sets the title of the page action. This is displayed in a tooltip over the page action.
func (p *PageAction) SetTitle(details Object) {
	p.o.Call("setTitle", details)
}

// GetTitle gets the title of the page action.
func (p *PageAction) GetTitle(details Object, callback func(result string)) {
	p.o.Call("getTitle", details, callback)
}

// SetIcon sets the icon for the page action. The icon can be specified either as the path to an image
// file or as the pixel data from a canvas element, or as dictionary of either one of those. Either the
// path or the imageData property must be specified.
func (p *PageAction) SetIcon(details Object, callback func()) {
	p.o.Call("setIcon", details, callback)
}

// SetPopup sets the html document to be opened as a popup when the user clicks on the page action's icon.
func (p *PageAction) SetPopup(details Object) {
	p.o.Call("setPopup", details)
}

// GetPopup gets the html document set as the popup for this page action.
func (p *PageAction) GetPopup(details Object, callback func(result string)) {
	p.o.Call("getPopup", details, callback)
}

/*
* Events
 */

// OnClicked fired when a page action icon is clicked. This event will not fire if the page action has a popup.
func (p *PageAction) OnClicked(callback func(tab Tab)) {
	p.o.Get("onClicked").Call("addListener", callback)
}
