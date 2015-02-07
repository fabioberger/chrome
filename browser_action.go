package chrome

import "github.com/gopherjs/gopherjs/js"

type BrowserAction struct {
	o js.Object
}

/*
* Methods:
 */

// SetTitle sets the title of the browser action. This shows up in the tooltip.
func (b *BrowserAction) SetTitle(details map[string]interface{}) {
	b.o.Call("setTitle", details)
}

// GetTitle gets the title of the browser action.
func (b *BrowserAction) GetTitle(details map[string]interface{}, callback func(result string)) {
	b.o.Call("getTitle", details, callback)
}

// SetIcon sets the icon for the browser action. The icon can be specified either as the path to an image
// file or as the pixel data from a canvas element, or as map of either one of those. Either the path
// or the imageData property must be specified.
func (b *BrowserAction) SetIcon(details map[string]interface{}, callback func()) {
	b.o.Call("setIcon", details, callback)
}

// SetPopup sets the html document to be opened as a popup when the user clicks on the browser action's icon.
func (b *BrowserAction) SetPopup(details map[string]interface{}) {
	b.o.Call("setPopup", details)
}

// GetPopup gets the html document set as the popup for this browser action.
func (b *BrowserAction) GetPopup(details map[string]interface{}, callback func(result string)) {
	b.o.Call("getPopup", details, callback)
}

// SetBadgeText sets the badge text for the browser action. The badge is displayed on top of the icon.
func (b *BrowserAction) SetBadgeText(details map[string]interface{}) {
	b.o.Call("setBadgeText", details)
}

// getBadgeText gets the badge text of the browser action. If no tab is specified, the non-tab-specific badge text is returned.
func (b *BrowserAction) getBadgeText(details map[string]interface{}, callback func(result string)) {
	b.o.Call("getBadgeText", details, callback)
}

// SetBadgeBackgroundColor sets the background color for the badge.
func (b *BrowserAction) SetBadgeBackgroundColor(details map[string]interface{}) {
	b.o.Call("setBadgeBackgroundColor", details)
}

// GetBadgeBackgroundColor gets the background color of the browser action.
func (b *BrowserAction) GetBadgeBackgroundColor(details map[string]interface{}, callback func(result []int)) {
	b.o.Call("getBadgeBackgroundColor", details, callback)
}

// Enable enables the browser action for a tab. By default, browser actions are enabled.
func (b *BrowserAction) Enable(tabId int) {
	b.o.Call("enable", tabId)
}

// Disable disables the browser action for a tab.
func (b *BrowserAction) Disable(tabId int) {
	b.o.Call("disable", tabId)
}

/*
* Events
 */

// OnClicked fired when a browser action icon is clicked. This event will not fire if the browser action has a popup.
func (b *BrowserAction) OnClicked(callback func(tab map[string]interface{})) {
	b.o.Get("onClicked").Call("addListener", callback)
}
