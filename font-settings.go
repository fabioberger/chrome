package chrome

import "github.com/gopherjs/gopherjs/js"

type FontSettings struct {
	o js.Object
}

/*
* Methods:
 */

// ClearFont clears the font set by this extension, if any.
func (f *FontSettings) ClearFont(details map[string]interface{}, callback func()) {
	f.o.Call("clearFont", details, callback)
}

// GetFont gets the font for a given script and generic font family.
func (f *FontSettings) GetFont(details map[string]interface{}, callback func(details map[string]interface{})) {
	f.o.Call("getFont", details, callback)
}

// SetFont sets the font for a given script and generic font family.
func (f *FontSettings) SetFont(details map[string]interface{}, callback func()) {
	f.o.Call("setFont", details, callback)
}

// GetFontList gets a list of fonts on the system.
func (f *FontSettings) GetFontList(callback func(results []map[string]interface{})) {
	f.o.Call("getFontList", callback)
}

// ClearDefaultFontSize clears the default font size set by this extension, if any.
func (f *FontSettings) ClearDefaultFontSize(details map[string]interface{}, callback func()) {
	f.o.Call("clearDefaultFontSize", details, callback)
}

// GetDefaultFontSize gets the default font size.
func (f *FontSettings) GetDefaultFontSize(details map[string]interface{}, callback func(details map[string]interface{})) {
	f.o.Call("getDefaultFontSize", details, callback)
}

// SetDefaultFontSize sets the default font size.
func (f *FontSettings) SetDefaultFontSize(details map[string]interface{}, callback func()) {
	f.o.Call("setDefaultFontSize", details, callback)
}

// ClearDefaultFixedFontSize clears the default fixed font size set by this extension, if any.
func (f *FontSettings) ClearDefaultFixedFontSize(details map[string]interface{}, callback func()) {
	f.o.Call("clearDefaultFixedFontSize", details, callback)
}

// GetDefaultFixedFontSize gets the default size for fixed width fonts.
func (f *FontSettings) GetDefaultFixedFontSize(details map[string]interface{}, callback func(details map[string]interface{})) {
	f.o.Call("getDefaultFixedFontSize", details, callback)
}

// SetDefaultFixedFontSize sets the default size for fixed width fonts.
func (f *FontSettings) SetDefaultFixedFontSize(details map[string]interface{}, callback func()) {
	f.o.Call("setDefaultFixedFontSize", details, callback)
}

// ClearMinimumFontSize lears the minimum font size set by this extension, if any.
func (f *FontSettings) ClearMinimumFontSize(details map[string]interface{}, callback func()) {
	f.o.Call("clearMinimumFontSize", details, callback)
}

// GetMinimumFontSize gets the minimum font size.
func (f *FontSettings) GetMinimumFontSize(details map[string]interface{}, callback func(details map[string]interface{})) {
	f.o.Call("getMinimumFontSize", details, callback)
}

// SetMinimumFontSize sets the minimum font size.
func (f *FontSettings) SetMinimumFontSize(details map[string]interface{}, callback func()) {
	f.o.Call("setMinimumFontSize", details, callback)
}

/*
* Events
 */

// OnFontChanged fired when a font setting changes.
func (f *FontSettings) OnFontChanged(callback func(details map[string]interface{})) {
	f.o.Get("onFontChanged").Call("addListener", callback)
}

// OnDefaultFontSizeChanged fired when the default font size setting changes.
func (f *FontSettings) OnDefaultFontSizeChanged(callback func(details map[string]interface{})) {
	f.o.Get("onDefaultFontSizeChanged").Call("addListener", callback)
}

// OnDefaultFixedFontSizeChanged fired when the default fixed font size setting changes.
func (f *FontSettings) OnDefaultFixedFontSizeChanged(callback func(details map[string]interface{})) {
	f.o.Get("onDefaultFixedFontSizeChanged").Call("addListener", callback)
}

// OnMinimumFontSizeChanged fired when the minimum font size setting changes.
func (f *FontSettings) OnMinimumFontSizeChanged(callback func(details map[string]interface{})) {
	f.o.Get("onMinimumFontSizeChanged").Call("addListener", callback)
}
