package chrome

import "github.com/gopherjs/gopherjs/js"

type Tts struct {
	o *js.Object
}

/*
* Types
 */

type TtsEvent struct {
	js.Object
	Type         string `js:"type"`
	CharIndex    int64  `js:"charIndex"`
	ErrorMessage string `js:"errorMessage"`
}

type TtsVoice struct {
	js.Object
	VoiceName   string   `js:"voiceName"`
	Lang        string   `js:"lang"`
	Gender      string   `js:"gender"`
	Remote      bool     `js:"remote"`
	ExtensionId string   `js:"extensionId"`
	EventTypes  []string `js:"eventTypes"`
}

/*
* Methods
 */

// Speak speaks text using a text-to-speech engine.
func (t *Tts) Speak(utterance string, options Object, callback func()) {
	t.o.Call("speak", utterance, options, callback)
}

// Stop stops any current speech and flushes the queue of any pending utterances.
// In addition, if speech was paused, it will now be un-paused for the next call to speak.
func (t *Tts) Stop() {
	t.o.Call("stop")
}

// Pause pauses speech synthesis, potentially in the middle of an utterance.
// A call to resume or stop will un-pause speech.
func (t *Tts) Pause() {
	t.o.Call("pause")
}

// Resume if speech was paused, resumes speaking where it left off.
func (t *Tts) Resume() {
	t.o.Call("resume")
}

// IsSpeaking checks whether the engine is currently speaking. On Mac OS X, the result
// is true whenever the system speech engine is speaking, even if the speech wasn't initiated by Chrome.
func (t *Tts) IsSpeaking(callback func(speaking bool)) {
	t.o.Call("isSpeaking", callback)
}

// GetVoices gets an array of all available voices.
func (t *Tts) GetVoices(callback func(voices []TtsVoice)) {
	t.o.Call("getVoices", callback)
}
