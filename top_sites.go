package chrome

import "github.com/gopherjs/gopherjs/js"

type TopSites struct {
	o js.Object
}

/*
* Types
 */

type MostvisitedURL struct {
	Url   string `js:"url"`
	Title string `js:"title"`
}

/*
* Methods
 */

// Get gets a list of top sites.
func (t *TopSites) Get(callback func(data []MostvisitedURL)) {
	t.o.Call("get", callback)
}
