package chrome

import "github.com/gopherjs/gopherjs/js"

type Identity struct {
	o js.Object
}

/*
* Types
 */

type AccountInfo struct {
	js.Object
	Id string `js:"id"`
}

/*
* Methods
 */

// GetAccounts retrieves a list of AccountInfo objects describing the accounts present on the profile.
func (i *Identity) GetAccounts(callback func(accounts []AccountInfo)) {
	i.o.Call("getAccounts", callback)
}

// GetAuthToken gets an OAuth2 access token using the client ID and scopes specified in the oauth2 section of manifest.json.
/* The Identity API caches access tokens in memory, so it's ok to call getAuthToken non-interactively any time a token is required. The token cache automatically handles expiration.
For a good user experience it is important interactive token requests are initiated by UI in your app explaining what the authorization is for. Failing to do this will cause your users to get authorization requests, or Chrome sign in screens if they are not signed in, with with no context. In particular, do not use getAuthToken interactively when your app is first launched. */
func (i *Identity) GetAuthToken(details Object, callback func(token string)) {
	i.o.Call("getAuthToken", details, callback)
}

// GetProfileUserInfo retrieves email address and obfuscated gaia id of the user signed into a profile.
// This API is different from identity.getAccounts in two ways. The information returned is available offline,
// and it only applies to the primary account for the profile.
func (i *Identity) GetProfileUserInfo(callback func(userInfo Object)) {
	i.o.Call("getProfileUserInfo", callback)
}

// RemoveCacheAuthToken removes an OAuth2 access token from the Identity API's token cache.
// If an access token is discovered to be invalid, it should be passed to removeCachedAuthToken
// to remove it from the cache. The app may then retrieve a fresh token with getAuthToken.
func (i *Identity) RemoveCacheAuthToken(details Object, callback func()) {
	i.o.Call("removeCacheAuthToken", details, callback)
}

// LaunchWebAuthFrom starts an auth flow at the specified URL.
/* This method enables auth flows with non-Google identity providers by launching a web view and navigating it to the first URL in the provider's auth flow. When the provider redirects to a URL matching the pattern https://<app-id>.chromiumapp.org/*, the window will close, and the final redirect URL will be passed to the callback function.
For a good user experience it is important interactive auth flows are initiated by UI in your app explaining what the authorization is for. Failing to do this will cause your users to get authorization requests with no context. In particular, do not launch an interactive auth flow when your app is first launched. */
func (i *Identity) LaunchWebAuthFrom(details Object, callback func(response string)) {
	i.o.Call("launchWebAuthFrom", details, callback)
}

// GetRedirectURL generates a redirect URL to be used in |launchWebAuthFlow|.
// The generated URLs match the pattern https://<app-id>.chromiumapp.org/*.
func (i *Identity) GetRedirectURL(path string) {
	i.o.Call("getRedirectURL", path)
}

/*
* Events
 */

// OnSignInChanged fired when signin state changes for an account on the user's profile.
func (i *Identity) OnSignInChanged(callback func(account AccountInfo, signedIn bool)) {
	i.o.Get("onSignInChanged").Call("addListener", callback)
}
