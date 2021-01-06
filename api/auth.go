package main

import (
	"net/http"
	"github.com/eggtart26/VideoStreamingInGO/api/session"
	// "github.com/eggtart26/VideoStreamingInGO/api"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-USER-Name"

func ValidateUserSession(r *http.Request) bool  {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		// sendErrorResponse()
		return false
	}

	return true
 }