package main

import (
	"fmt"
	"time"

	"github.com/fabioberger/chrome"
	QUnit "github.com/fabioberger/qunit"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

var (
	doc = dom.GetWindow().Document()
)

func main() {
	c := chrome.NewChrome()

	QUnit.Module("Chrome-Popup")

	/*
	* Alarm Method Tests
	 */

	// Create New Alarm
	alarmOps := chrome.Object{
		"when": time.Now().UnixNano() + 1000000,
	}
	c.Alarms.Create("test_alarm", alarmOps)

	// Get the Alarm created above
	c.Alarms.Get("test_alarm", func(alarm chrome.Alarm) {
		QUnit.Test("Alarm.Get()", func(assert QUnit.QUnitAssert) {
			assert.Equal(alarm.Name, "test_alarm", "Get")
		})

		// Clear the Alarm retrieved above
		c.Alarms.Clear("test_alarm", func(wasCleared bool) {
			QUnit.Test("Alarm.Clear()", func(assert QUnit.QUnitAssert) {
				assert.Equal(wasCleared, true, "Clear")
			})

			// Create two more Alarms
			c.Alarms.Create("test_alarm2", alarmOps)
			c.Alarms.Create("test_alarm3", alarmOps)

			// Get both Alarms created above
			c.Alarms.GetAll(func(alarms []chrome.Alarm) {
				QUnit.Test("Alarm.GetAll()", func(assert QUnit.QUnitAssert) {
					assert.Equal(alarms[0].Name, "test_alarm2", "GetAll")
					assert.Equal(alarms[1].Name, "test_alarm3", "GetAll")
				})
				// Clear both Alarms above
				c.Alarms.ClearAll(func(wasCleared bool) {
					QUnit.Test("Alarm.ClearAll()", func(assert QUnit.QUnitAssert) {
						assert.Equal(wasCleared, true, "Clear")
					})
				})
			})
		})
	})

	/*
	* Bookmarks Method Test
	 */
	bookmark := chrome.Object{
		"title": "Testing",
		"url":   "http://www.testing.com/",
	}
	// Test Create New Bookmarks
	c.Bookmarks.Create(bookmark, func(result chrome.BookmarkTreeNode) {
		QUnit.Test("Bookmarks.Create()", func(assert QUnit.QUnitAssert) {
			assert.Equal(result.Title, "Testing", "Create")
		})

		// Test Get Bookmark by List of Id's
		c.Bookmarks.Get([]string{result.Id}, func(results []chrome.BookmarkTreeNode) {
			QUnit.Test("Bookmarks.Get()", func(assert QUnit.QUnitAssert) {
				assert.Equal(results[0].Url, "http://www.testing.com/", "Get")
			})
		})
	})

	/*
	* BrowserAction Method Tests
	 */

	change := chrome.Object{
		"title": "Testing",
	}
	c.BrowserAction.SetTitle(change)
	c.BrowserAction.GetTitle(chrome.Object{}, func(result string) {
		QUnit.Test("BrowserAction.GetTitle()", func(assert QUnit.QUnitAssert) {
			assert.Equal(result, "Testing", "GetTitle")
		})
	})

	/*
	* BrowsingData Method Tests
	 */

	// Test Retrieving BrowserData Settings
	c.BrowsingData.Settings(func(result chrome.Object) {
		for key, _ := range result {
			QUnit.Test("BrowsingData.Settings()", func(assert QUnit.QUnitAssert) {
				assert.Equal(key, "dataRemovalPermitted", "Settings")
			})
			break
		}
	})

	/*
	* Cookies Method Tests
	 */

	// Set a new Cookie
	cookieInfo := chrome.Object{
		"url":   "http://www.google.com",
		"name":  "testing",
		"value": "testvalue",
	}
	c.Cookies.Set(cookieInfo, func(cookie chrome.Cookie) {
		QUnit.Test("Cookies.Set()", func(assert QUnit.QUnitAssert) {
			assert.Equal(cookie.Name, "testing", "Set")
			assert.Equal(cookie.Value, "testvalue", "Set")
		})

		// Get the cookie set previously
		cookieInfo = chrome.Object{
			"url":  "http://www.google.com",
			"name": "testing",
		}
		c.Cookies.Get(cookieInfo, func(cookie chrome.Cookie) {
			QUnit.Test("Cookies.Get()", func(assert QUnit.QUnitAssert) {
				assert.Equal(cookie.Name, "testing", "Get")
			})
		})
	})

	/*
	* Extension Method Tests
	 */

	// Get the popup views for the Extension
	fetchProperties := chrome.Object{
		"type": "popup",
	}
	windows := c.Extension.GetViews(fetchProperties)
	QUnit.Test("Extension.GetViews()", func(assert QUnit.QUnitAssert) {
		assert.Equal(windows[0].Incognito, false, "GetViews")
	})

	/*
	* FontSettings Method Tests
	 */

	// Get Details for a Generic Font Family
	fontDetails := chrome.Object{
		"genericFamily": "standard",
		"script":        "Arab",
	}
	c.FontSettings.GetFont(fontDetails, func(details chrome.Object) {
		for key, _ := range details {
			QUnit.Test("FontSettings.GetFont()", func(assert QUnit.QUnitAssert) {
				assert.Equal(key, "fontId", "GetFont")
			})
			break
		}
	})

	/*
	* History Method Tests
	 */

	// Add URL to History
	urlDetails := chrome.Object{
		"url": "http://www.testing.com/",
	}
	c.History.AddUrl(urlDetails, func() {

		// Search for created URL in History
		s := chrome.Object{
			"text": "www.testing.com",
		}
		c.History.Search(s, func(results []chrome.HistoryItem) {
			QUnit.Test("History.Search()", func(assert QUnit.QUnitAssert) {
				assert.Equal(results[0].Url, "http://www.testing.com/", "Search")
			})
		})
	})

	/*
	* Tab Method Test
	 */

	// Find the current Tab
	queryInfo := chrome.Object{
		"currentWindow": true,
		"active":        true,
	}
	c.Tabs.Query(queryInfo, func(tabs []chrome.Tab) {
		id := tabs[0].Id

		// Send Message to the given Tab & have Event Listener on Tab Respond
		msg := chrome.Object{"greeting": "hello"}
		c.Tabs.SendMessage(id, msg, func(response chrome.Object) {
			err := js.Global.Get("chrome").Get("runtime").Get("lastError")
			if err.String() != "undefined" {
				fmt.Println("Tabs.SendMessage Error: ", err.Get("message").String())
			}
			QUnit.Test("Tabs.SendMessage() & Runtime.OnMessage() Event", func(assert QUnit.QUnitAssert) {
				assert.Equal(response["farewell"], "goodbye", "SendMessage")
			})
		})
	})

}
