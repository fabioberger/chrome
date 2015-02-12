package chrome

import "github.com/gopherjs/gopherjs/js"

type Omnibox struct {
	o js.Object
}

/*
* Types
 */

type SuggestResult struct {
	js.Object
	Content     string `js:"content"`
	Description string `js:"description"`
}

/*
* Methods
 */

// SetDefaultSuggestion sets the description and styling for the default suggestion.
// The default suggestion is the text that is displayed in the first suggestion row underneath the URL bar.
func (m *Omnibox) SetDefaultSuggestion(suggestion map[string]interface{}) {
	m.o.Call("setDefaultSuggestion", suggestion)
}

/*
* Events
 */

// OnInputStarted user has started a keyword input session by typing the extension's keyword.
// This is guaranteed to be sent exactly once per input session, and before any onInputChanged events.
func (m *Omnibox) OnInputStarted(callback func()) {
	m.o.Get("onInputStarted").Call("addListener", callback)
}

// OnInputChanged user has changed what is typed into the omnibox.
func (m *Omnibox) OnInputChanged(callback func(text string, suggest func(suggestResults []SuggestResult))) {
	m.o.Get("onInputChanged").Call("addListener", callback)
}

// OnInputEntered user has accepted what is typed into the omnibox.
func (m *Omnibox) OnInputEntered(callback func(text string, disposition string)) {
	m.o.Get("onInputEntered").Call("addListener", callback)
}

// OnInputCancelled user has ended the keyword input session without accepting the input.
func (m *Omnibox) OnInputCancelled(callback func()) {
	m.o.Get("onInputCancelled").Call("addListener", callback)
}
