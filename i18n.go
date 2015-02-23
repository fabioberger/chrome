package chrome

import "github.com/gopherjs/gopherjs/js"

type I18n struct {
	o *js.Object
}

/*
* Methods:
 */

// GetAcceptLanguages gets the accept-languages of the browser. This is different from
// the locale used by the browser; to get the locale, use i18n.getUILanguage.
func (i *I18n) GetAcceptLanguages(callback func(languages []string)) {
	i.o.Call("getAcceptLanguages", callback)
}

// GetMessage gets the localized string for the specified message. If the message is missing,
// this method returns an empty string (''). If the format of the getMessage() call is wrong —
// for example, messageName is not a string or the substitutions array has more than 9 elements —
// this method returns undefined.
func (i *I18n) GetMessage(messageName string, substitutions interface{}) string {
	return i.o.Call("getMessage", messageName, substitutions).String()
}

// GetUILanguage gets the browser UI language of the browser. This is different from
// i18n.getAcceptLanguages which returns the preferred user languages.
func (i *I18n) GetUILanguage() string {
	return i.o.Call("getUILanguage").String()
}
