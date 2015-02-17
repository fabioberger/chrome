package chrome

import "github.com/gopherjs/gopherjs/js"

type Privacy struct {
	Services js.Object
	Network  js.Object
	Websites js.Object
}

func NewPrivacy(privacyObj js.Object) *Privacy {
	p := new(Privacy)
	if privacyObj.String() != "undefined" {
		p.Services = privacyObj.Get("services")
		p.Network = privacyObj.Get("network")
		p.Websites = privacyObj.Get("websites")
	}
	return p
}

// Could add properties here and make them more easily accessable
// For now you'll need to do something like this:
/*
chrome.Privacy.Services.Get("autofillEnabled").Call("get", chrome.Object{}, func(details chrome.Object) {
	if(details["levelOfControl"] == 'controllable_by_this_extension') {
		fmt.Println("We are in control!")
	}
	})
*/
