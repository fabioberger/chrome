package chrome

import "github.com/gopherjs/gopherjs/js"

type History struct {
	o js.Object
}

/*
* Types
 */

type HistoryItem struct {
	js.Object
	Id            string `js:"id"`
	Url           string `js:"url"`
	Title         string `js:"title"`
	LastVisitTime int64  `js:"lastVisitTime"`
	VisitCount    int    `js:"visitCount"`
	TypedCount    int    `js:"typedCount"`
}

type VisitItem struct {
	js.Object
	Id               string `js:"id"`
	VisitId          string `js:"visitId"`
	VisitTime        int64  `js:"visitTime"`
	ReferringVisitId string `js:"referringVisitId"`
	Transition       string `js:"transition"`
}

/*
* Methods:
 */

// Search searches the history for the last visit time of each page matching the query.
func (h *History) Search(query Object, callback func(results []HistoryItem)) {
	h.o.Call("search", query, callback)
}

// GetVisits retrieves information about visits to a URL.
func (h *History) GetVisits(details Object, callback func(results []VisitItem)) {
	h.o.Call("getVisits", details, callback)
}

// AddUrl adds a URL to the history at the current time with a transition type of "link".
func (h *History) AddUrl(details Object, callback func()) {
	h.o.Call("addUrl", details, callback)
}

// DeleteUrl removes all occurrences of the given URL from the history.
func (h *History) DeleteUrl(details Object, callback func()) {
	h.o.Call("deleteUrl", details, callback)
}

// DeleteRange removes all items within the specified date range from the history.
// Pages will not be removed from the history unless all visits fall within the range.
func (h *History) DeleteRange(rang Object, callback func()) {
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
func (h *History) OnVisited(callback func(result HistoryItem)) {
	h.o.Get("onVisited").Call("addListener", callback)
}

// OnVisitedRemoved fired when one or more URLs are removed from the history service.
// When all visits have been removed the URL is purged from history.
func (h *History) OnVisitedRemoved(callback func(removed Object)) {
	h.o.Get("onVisitedRemoved").Call("addListener", callback)
}
