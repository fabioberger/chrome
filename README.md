Gopherjs Chrome Bindings
------------------------------------

[![GoDoc](http://godoc.org/github.com/fabioberger/chrome?status.svg)](https://godoc.org/github.com/fabioberger/chrome)

Write Chrome extensions in Golang! This library contains the necessary bindings to interact with the Chrome Javascript API in an easy and intuitive way. 

**Warning:** This library is not yet complete. If there are any bindings missing that you would like to contribute, feel free to submit a pull request!

# Installation

Install with go get:

```bash
go get github.com/fabioberger/chrome
```

Then include the package in your imports:

```bash
import "github.com/fabioberger/chrome"
```

# Example Usage

This simple example will print out the current tab's id to the chrome extension's popup console window:

```go
package main

import (
	"fmt"
	"github.com/fabioberger/chrome"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	c := chrome.NewChrome()
	c.Tabs.CurrentSelected(chrome.WINDOW_ID_CURRENT, func(tab js.Object) {
	 	fmt.Println("Current Tab Id: ", tab.Get("id").String())
	})
}
```