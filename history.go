package chrome

import "github.com/gopherjs/gopherjs/js"

type History struct {
	o js.Object
}

/*
* Methods:
 */

// Search searches the history for the last visit time of each page matching the query.
func (h *History) Search(query map[string]interface{}, callback func(results []map[string]interface{})) {
	h.o.Call("search", query, callback)
}

// GetVisits retrieves information about visits to a URL.
func (h *History) GetVisits(details map[string]interface{}, callback func(results []map[string]interface{})) {
	h.o.Call("getVisits", details, callback)
}

// AddUrl adds a URL to the history at the current time with a transition type of "link".
func (h *History) AddUrl(details map[string]interface{}, callback func()) {
	h.o.Call("addUrl", details, callback)
}

// DeleteUrl removes all occurrences of the given URL from the history.
func (h *History) DeleteUrl(details map[string]interface{}, callback func()) {
	h.o.Call("deleteUrl", details, callback)
}

// DeleteRange removes all items within the specified date range from the history.
// Pages will not be removed from the history unless all visits fall within the range.
func (h *History) DeleteRange(rang map[string]interface{}, callback func()) {
	h.o.Call("deleteRange", rang, callback)
}

// DeleteAll deletes all items from the history.
func (h *History) DeleteAll(callback func()) {
	h.o.Call("deleteAll", callback)
}

/*
* Events
 */

// OnVisited fired when a URL is visited, providing the HistoryItem data for that URL.
// This event fires before the page has loaded.
func (h *History) OnVisited(callback func(result map[string]interface{})) {
	h.o.Get("onVisited").Call("addListener", callback)
}

// OnVisitedRemoved fired when one or more URLs are removed from the history service.
// When all visits have been removed the URL is purged from history.
func (h *History) OnVisitedRemoved(callback func(removed map[string]interface{})) {
	h.o.Get("onVisitedRemoved").Call("addListener", callback)
}
