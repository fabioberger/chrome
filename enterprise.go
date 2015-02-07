package chrome

import "github.com/gopherjs/gopherjs/js"

type Enterprise struct {
	o            js.Object
	PlatformKeys PlatformKeys
}

type PlatformKeys struct {
	o js.Object
}

/*
* Methods
 */

/* GetTokens returns the available Tokens. In a regular user's session the list will always contain the user's token with id "user". If a system-wide TPM token is available, the returned list will also contain the system-wide token with id "system". The system-wide token will be the same for all sessions on this device (device in the sense of e.g. a Chromebook). */
func (p *PlatformKeys) GetTokens(callback func(tokens []map[string]interface{})) {
	p.o.Call("getTokens", callback)
}

// GetCertificates returns the list of all client certificates available from the given token.
// Can be used to check for the existence and expiration of client certificates that are usable for a certain authentication.
func (p *PlatformKeys) GetCertificates(tokenId string, callback func(certificates []map[string]interface{})) {
	p.o.Call("getCertificates", tokenId, callback)
}

// ImportCertificates imports certificate to the given token if the certified key is already stored in this token.
// After a successful certification request, this function should be used to store the obtained certificate and to
//  make it available to the operating system and browser for authentication.
func (p *PlatformKeys) ImportCertificates(tokenId string, certificate map[string]interface{}, callback func()) {
	p.o.Call("importCertificates", tokenId, certificate, callback)
}

// RemoveCertificate removes certificate from the given token if present. Should be used to remove obsolete
// certificates so that they are not considered during authentication and do not clutter the certificate choice.
// Should be used to free storage in the certificate store.
func (p *PlatformKeys) RemoveCertificate(tokenId string, certificate map[string]interface{}, callback func()) {
	p.o.Call("removeCertificate", tokenId, certificate, callback)
}
