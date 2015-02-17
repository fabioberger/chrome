package chrome

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

type DeclarativeContent struct {
	o             js.Object
	OnPageChanged *OnPageChanged //Need to remove this and fix methods below...
}

type OnPageChanged struct {
	o js.Object
}

func NewDeclarativeContent(declarativeContentObj js.Object) *DeclarativeContent {
	d := new(DeclarativeContent)
	d.o = declarativeContentObj
	fmt.Println(d.o.String())
	if d.o.String() != "undefined" {
		d.OnPageChanged = &OnPageChanged{o: d.o.Get("onPageChanged")}
	} else {
		d.OnPageChanged = &OnPageChanged{o: nil}
	}
	return d
}

/*
* Types
 */

type PageStateMatcher struct {
	js.Object
	PageUrl Object   `js:"pageUrl"`
	Css     []string `js:"css"`
}

type RequestContentScript struct {
	js.Object
	Css             []string `js:"css"`
	Js              []string `js:"js"`
	AllFrames       bool     `js:"allFrames"`
	MatchAboutBlank bool     `js:"matchAboutBlank"`
}

/*
* Events
 */

// AddRules takes an array of rule instances as its first parameter and a callback function that is called on completion
func (e *OnPageChanged) AddRules(rules []map[string]interface{}, callback func(details map[string]interface{})) {
	e.o.Call("addRules", rules, callback)
}

// RemoveRules accepts an optional array of rule identifiers as its first parameter and a callback function as its second parameter.
func (e *OnPageChanged) RemoveRules(ruleIdentifiers []string, callback func()) {
	e.o.Call("removeRules", ruleIdentifiers, callback)
}

// GetRules accepts an optional array of rule identifiers with the same semantics as removeRules and a callback function.
func (e *OnPageChanged) GetRules(ruleIdentifiers []string, callback func(details map[string]interface{})) {
	e.o.Call("getRules", ruleIdentifiers, callback)
}
