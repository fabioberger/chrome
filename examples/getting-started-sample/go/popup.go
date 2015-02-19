// This example is a Go implementation of the Getting Started Sample Extension from
// this tutorial: https://developer.chrome.com/extensions/getstarted

package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"

	"github.com/fabioberger/chrome"
)

var (
	doc = dom.GetWindow().Document()
)

/**
 * getCurrentTabUrl gets the current URL.
 *
 * @param {func(string)} callback - called when the URL of the current tab
 *   is found.
 **/
func getCurrentTabUrl(c *chrome.Chrome, callback func(string)) {
	// Query filter to be passed to chrome.tabs.query - see
	// https://developer.chrome.com/extensions/tabs#method-query
	queryInfo := chrome.Object{
		"active":        true,
		"currentWindow": true,
	}

	c.Tabs.Query(queryInfo, func(tabs []chrome.Tab) {
		// c.Tabs.Query invokes the callback with a list of tabs that match the
		// query. When the popup is opened, there is certainly a window and at least
		// one tab, so we can safely assume that |tabs| is a non-empty array.
		// A window can only have one active tab at a time, so the array consists of
		// exactly one tab.
		tab := tabs[0]

		// A tab is a struct that provides information about the tab.
		// See https://developer.chrome.com/extensions/tabs#type-Tab
		url := tab.Url

		callback(url)
	})

	// Most methods of the Chrome extension APIs are asynchronous. This means that
	// you CANNOT do something like this:
	//
	// var url string
	// c.Tabs.Query(queryInfo, func(tabs []chrome.Tab) {
	//   url = tabs[0].Url
	// })
	// fmt.Println(url) // Shows "undefined", because c.Tabs.Query is async.
}

/**
 * @param {string} searchTerm - Search term for Google Image search.
 * @param {func(string,int,int)} callback - Called when an image has
 *   been found. The callback gets the URL, width and height of the image.
 * @param {func(string)} errorCallback - Called when the image is not found.
 *   The callback gets a string that describes the failure reason.
 */
func getImageUrl(searchTerm string, callback func(imageUrl, width, height string), errorCallback func(errorMessage string)) {
	// Google image search - 100 searches per day.
	// https://developers.google.com/image-search/
	searchUrl := "https://ajax.googleapis.com/ajax/services/search/images" + "?v=1.0&q=" + url.QueryEscape(searchTerm)

	req := xhr.NewRequest("GET", searchUrl)
	req.Timeout = 2000 // two second, in milliseconds
	req.ResponseType = "text"
	err := req.Send(nil)
	if err != nil {
		errorCallback("Something went wrong with GET request to " + searchUrl + "." + err.Error())
		return
	}
	returnFormat := struct {
		ResponseData struct {
			Results []struct {
				Url      string `json:"Url"`
				TbWidth  string `json:"tbWidth"`
				TbHeight string `json:"tbHeight"`
			} `json:"results"`
		} `json:"responseData"`
	}{}
	err = json.Unmarshal([]byte(req.Response.String()), &returnFormat)
	if err != nil {
		errorCallback("Failed to unmarshal response into object, with Error: " + err.Error() + ". Response was: " + req.Response.String())
		return
	}
	fmt.Println("Got from API: ", returnFormat)

	firstResult := returnFormat.ResponseData.Results[0]
	imageUrl := firstResult.Url
	width := firstResult.TbWidth
	height := firstResult.TbHeight

	callback(imageUrl, width, height)
}

func renderStatus(statusText string) {
	doc.GetElementByID("status").SetTextContent(statusText)
}

func main() {
	c := chrome.NewChrome()
	getCurrentTabUrl(c, func(url string) {
		// Put the image URL in Google search.
		renderStatus("Performing Google Image search for " + url)

		go getImageUrl(url, func(imageUrl, width, height string) {
			renderStatus("Search term: " + url + "\n" + "Google image search result: " + imageUrl)
			var imageResult = doc.GetElementByID("image-result")
			// Explicitly set the width/height to minimize the number of reflows. For
			// a single image, this does not matter, but if you're going to embed
			// multiple external images in your page, then the absence of width/height
			// attributes causes the popup to resize multiple times.
			imageResult.SetAttribute("width", width)
			imageResult.SetAttribute("height", height)
			imageResult.SetAttribute("src", imageUrl)
			imageResult.SetAttribute("display", "block")
		}, func(errorMessage string) {
			renderStatus("Cannot display image. " + errorMessage)
		})
	})
}
