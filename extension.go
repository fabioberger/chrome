package chrome

import "github.com/gopherjs/gopherjs/js"

type Extension struct {
	o js.Object
}

/*
* Methods:
 */

func (e *Extension) SendRequest(extensionId string, request interface{}, responseCallback func(response interface{})) {
	e.o.Call("sendRequest", extensionId, request, responseCallback)
}

func (e *Extension) GetURL(path string) {
	e.o.Call("getURL", path)
}

func (e *Extension) GetViews(fetchProperties map[string]interface{}) interface{} {
	return e.o.Call("getViews", fetchProperties).Interface()
}
