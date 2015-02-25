package main

import (
	"fmt"

	"github.com/fabioberger/chrome"
)

func main() {
	c := chrome.NewChrome()

	/*
	* Runtime Events
	 */

	// Listen for OnMessage Event
	c.Runtime.OnMessage(func(message interface{}, sender chrome.MessageSender, sendResponse func(interface{})) {
		fmt.Println("Runtime.OnMessage received: ", message.(map[string]interface{}))
		resp := chrome.Object{
			"farewell": "goodbye",
		}
		sendResponse(resp)
	})
}
