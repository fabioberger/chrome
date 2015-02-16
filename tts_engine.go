package chrome

import "github.com/gopherjs/gopherjs/js"

type TtsEngine struct {
	o js.Object
}

/*
* Events
 */

// OnSpeak called when the user makes a call to tts.speak() and one of the voices from this
// extension's manifest is the first to match the options object.
func (t *TtsEngine) OnSpeak(callback func(utterance string, options map[string]interface{}, sendItsEvent func(event TtsEvent))) {
	t.o.Get("onSpeak").Call("addListener", callback)
}

// OnStop fired when a call is made to tts.stop and this extension may be in the middle of speaking.
// If an extension receives a call to onStop and speech is already stopped, it should do nothing
// (not raise an error). If speech is in the paused state, this should cancel the paused state.
func (t *TtsEngine) OnStop(callback func()) {
	t.o.Get("onStop").Call("addListener", callback)
}

// OnPause is optional: if an engine supports the pause event, it should pause the current utterance
// being spoken, if any, until it receives a resume event or stop event. Note that a stop event should
// also clear the paused state.
func (t *TtsEngine) OnPause(callback func()) {
	t.o.Get("onPause").Call("addListener", callback)
}

// OnResume is optional: if an engine supports the pause event, it should also support the resume event,
// to continue speaking the current utterance, if any. Note that a stop event should also clear the paused state.
func (t *TtsEngine) OnResume(callback func()) {
	t.o.Get("onResume").Call("addListener", callback)
}
