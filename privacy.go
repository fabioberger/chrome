package chrome

import "github.com/gopherjs/gopherjs/js"

type Privacy struct {
	o js.Object
}

// Could add properties here and make them more easily accessable
// For now you'll need to do something like this:
/*
chrome.Privacy.Get("services").Get("autofillEnabled").Call("get", map[string]interface{}{}, func(details map[string]interface{}) {
	if(details["levelOfControl"] == 'controllable_by_this_extension') {
		fmt.Println("We are in control!")
	}
	})
*/
