package chrome

import "github.com/gopherjs/gopherjs/js"

type Input struct {
	o   js.Object
	Ime *Ime
}

type Ime struct {
	o js.Object
}

func NewInput(inputObj js.Object) *Input {
	i := new(Input)
	i.o = inputObj
	if i.o.String() == "undefined" {
		i.Ime = &Ime{o: nil}
	} else {
		i.Ime = &Ime{o: i.o.Get("ime")}
	}
	return i
}

/*
* Types
 */

type KeyboardEvent struct {
	js.Object
	Type        string `js:"type"`
	RequestId   string `js:"requestId"`
	ExtensionId string `js:"extensionId"`
	Key         string `js:"key"`
	Code        string `js:"code"`
	KeyCode     int    `js:"keyCode"`
	AltKey      bool   `js:"altKey"`
	CtrlKey     bool   `js:"ctrlKey"`
	ShiftKey    bool   `js:"shiftKey"`
	CapsLock    bool   `js:"capsLock"`
}

type InputContext struct {
	js.Object
	ContextID    int    `js:"contextID"`
	Type         string `js:"type"`
	AutoCorrect  bool   `js:"autoCorrect"`
	AutoComplete bool   `js:"autoComplete"`
	SpellCheck   bool   `js:"spellCheck"`
}

type MenuItem struct {
	js.Object
	Id      string `js:"id"`
	Label   string `js:"label"`
	Style   string `js:"style"`
	Visible bool   `js:"visible"`
	Checked bool   `js:"checked"`
	Enabled bool   `js:"enabled"`
}

/*
* Methods
 */

// SetComposition set the current composition. If this extension does not own the active IME, this fails.
func (i *Ime) SetComposition(parameters Object, callback func(success bool)) {
	i.o.Call("setComposition", parameters, callback)
}

// ClearComposition clear the current composition. If this extension does not own the active IME, this fails.
func (i *Ime) ClearComposition(parameters Object, callback func(success bool)) {
	i.o.Call("clearComposition", parameters, callback)
}

// CommitText commits the provided text to the current input.
func (i *Ime) CommitText(parameters Object, callback func(success bool)) {
	i.o.Call("commitText", parameters, callback)
}

// SendKeyEvents sends the key events. This function is expected to be used by virtual keyboards.
// When key(s) on a virtual keyboard is pressed by a user, this function is used to propagate that event to the system.
func (i *Ime) SendKeyEvents(parameters Object, callback func()) {
	i.o.Call("sendKeyEvents", parameters, callback)
}

// HideInputView hides the input view window, which is popped up automatically by system.
// If the input view window is already hidden, this function will do nothing.
func (i *Ime) HideInputView() {
	i.o.Call("hideInputView")
}

// SetCandidateWindowProperties sets the properties of the candidate window. This fails if the extension doesn't own the active IME
func (i *Ime) SetCandidateWindowProperties(parameters Object, callback func(success bool)) {
	i.o.Call("setCandidateWindowProperties", parameters, callback)
}

// SetCandidates sets the current candidate list. This fails if this extension doesn't own the active IME
func (i *Ime) SetCandidates(parameters Object, callback func(success bool)) {
	i.o.Call("setCandidates", parameters, callback)
}

// SetCursorPosition set the position of the cursor in the candidate window. This is a no-op if this extension does not own the active IME.
func (i *Ime) SetCursorPosition(parameters Object, callback func(success bool)) {
	i.o.Call("setCursorPosition", parameters, callback)
}

// SetMenuItems adds the provided menu items to the language menu when this IME is active.
func (i *Ime) SetMenuItems(parameters Object, callback func()) {
	i.o.Call("setMenuItems", parameters, callback)
}

// UpdateMenuItems updates the state of the MenuItems specified
func (i *Ime) UpdateMenuItems(parameters Object, callback func()) {
	i.o.Call("updateMenuItems", parameters, callback)
}

// DeleteSurroundingText deletes the text around the caret.
func (i *Ime) DeleteSurroundingText(parameters Object, callback func()) {
	i.o.Call("deleteSurroundingText", parameters, callback)
}

// KeyEventHandled indicates that the key event received by onKeyEvent is handled.
//This should only be called if the onKeyEvent listener is asynchronous.
func (i *Ime) KeyEventHandled(requestId string, response bool) {
	i.o.Call("keyEventHandled", requestId, response)
}

/*
* Events
 */

// OnActivate this event is sent when an IME is activated. It signals that the IME
// will be receiving onKeyPress events.
func (i *Ime) OnActivate(callback func(engineID string, screen string)) {
	i.o.Get("onActivate").Call("addListener", callback)
}

// OnDeactivated this event is sent when an IME is deactivated. It signals that the IME
// will no longer be receiving onKeyPress events.
func (i *Ime) OnDeactivated(callback func(engineID string)) {
	i.o.Get("onDeactivated").Call("addListener", callback)
}

// OnFocus this event is sent when focus enters a text box. It is sent to all extensions
//that are listening to this event, and enabled by the user.
func (i *Ime) OnFocus(callback func(context InputContext)) {
	i.o.Get("onFocus").Call("addListener", callback)
}

// OnBlur this event is sent when focus leaves a text box. It is sent to all extensions
// that are listening to this event, and enabled by the user.
func (i *Ime) OnBlur(callback func(contextId int)) {
	i.o.Get("onBlur").Call("addListener", callback)
}

// OnInputContextUpdate this event is sent when the properties of the current InputContext change,
// such as the the type. It is sent to all extensions that are listening to this event, and enabled by the user.
func (i *Ime) OnInputContextUpdate(callback func(context InputContext)) {
	i.o.Get("onInputContextUpdate").Call("addListener", callback)
}

// OnKeyEvent this event is sent if this extension owns the active IME.
func (i *Ime) OnKeyEvent(callback func(engineID string, keyData KeyboardEvent)) {
	i.o.Get("onKeyEvent").Call("addListener", callback)
}

// OnCandidateClicked this event is sent if this extension owns the active IME.
func (i *Ime) OnCandidateClicked(callback func(engineID string, candidateID int, button string)) {
	i.o.Get("onCandidateClicked").Call("addListener", callback)
}

// OnMenuItemActivated called when the user selects a menu item
func (i *Ime) OnMenuItemActivated(callback func(engineID string, name string)) {
	i.o.Get("onMenuItemActivated").Call("addListener", callback)
}

// OnSurroundingTextChanged called when the editable string around caret is changed or when the caret
// position is moved. The text length is limited to 100 characters for each back and forth direction.
func (i *Ime) OnSurroundingTextChanged(callback func(engineID string, surroundingInfo Object)) {
	i.o.Get("onSurroundingTextChanged").Call("addListener", callback)
}

// OnReset this event is sent when chrome terminates ongoing text input session.
func (i *Ime) OnReset(callback func(engineID string)) {
	i.o.Get("onReset").Call("addListener", callback)
}
