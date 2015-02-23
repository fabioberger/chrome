package main

import (
	"fmt"

	"github.com/fabioberger/chrome"
)

func main() {
	c := chrome.NewChrome()

	// Listen for OnMessage Event
	fmt.Println("We are running")
	c.Runtime.OnMessage(func(message interface{}, sender chrome.MessageSender, sendResponse func(interface{})) {
		fmt.Println("Got in here")
		resp := chrome.Object{
			"farewell": "goodbye",
		}
		sendResponse(resp)
	})
}
