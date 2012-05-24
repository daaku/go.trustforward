// Package trustforward provides wrappers that trust X-Forwarded-*
// headers when looking up certain values.
package trustforward

import (
	"flag"
	"net/http"
)

var trust = flag.Bool(
	"trustforward", false, "Control trust of x-forwarded headers.")

// Get the Host.
func Host(r *http.Request) string {
	if *trust {
		if fwdHost := r.Header.Get("x-forwarded-host"); fwdHost != "" {
			return fwdHost
		}
	}
	return r.Host
}

// Get the Scheme.
func Scheme(r *http.Request) string {
	if *trust {
		if fwdScheme := r.Header.Get("x-forwarded-proto"); fwdScheme != "" {
			return fwdScheme
		}
	}
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

// Get the Remote Address.
func Remote(r *http.Request) string {
	if *trust {
		if fwdRemote := r.Header.Get("x-forwarded-for"); fwdRemote != "" {
			return fwdRemote
		}
	}
	return r.RemoteAddr
}
