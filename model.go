package goaudit

import "time"

// Audit represents Audit Record information. The purpose of this kind
// of records is access audits and billing.
type Audit struct {

	// Id is a random string to identify a specific Audit.
	Id string `json:"id" bson:"_id,omitempty"`

	// Version is the specification version used in this Audit.
	Version string `json:"version" bson:"version"`

	// AuthId is the authentication identifier in the sistem mapped as `auth_id`.
	AuthId string `json:"auth_id" bson:"auth_id"`

	// SessionId stores cookie session, in case the client was using a session.
	SessionId string `json:"session_id" bson:"session_id"`

	// Origin is the real client IP (v4 and v6 are supported). If the service is
	// behind a proxy, the real IP should be forwarded in the header
	// `X-Forwarded-For`.
	Origin string `json:"origin" bson:"origin"`

	// Service indicates the service/server being provided.
	Service *Service `json:"service" bson:"service"`

	// EntryDate stores a `Time` object with the starting request timestamp.
	EntryDate time.Time `json:"entry_date" bson:"entry_date"`

	// EntryTimestamp stores `EntryDate` in UNIX time format in seconds.
	EntryTimestamp float64 `json:"entry_timestamp" bson:"entry_timestamp,minsize"`

	// ElapsedSeconds is the real time the request has taken in seconds.
	ElapsedSeconds float64 `json:"elapsed_seconds" bson:"elapsed_seconds,minsize"`

	// Request is a struct with standard HTTP client request information
	// (automatically filled up).
	Request Request `json:"request" bson:"request"`

	// Response is a struct with standard HTTP response information.
	Response Response `json:"response" bson:"response"`

	// Array with all authIds allowed to read this Audit.
	ReadAccess []string `json:"read_access" bson:"read_access"`

	// Custom stores specific service information that is service-dependent in
	// any format.
	Custom interface{} `json:"custom" bson:"custom"`
}

// SetError overwrite error code and description in an Audit
func (audit *Audit) SetError(code int, desc string) {
	audit.Response.Error = &Error{
		Code:        code,
		Description: desc,
	}
}

// AddReadAccess add a authId to read access list for this Audit
func (audit *Audit) AddReadAccess(authId string) bool {
	if "" == authId {
		return false
	}

	for _, element := range audit.ReadAccess {
		if authId == element {
			return false
		}
	}
	audit.ReadAccess = append(audit.ReadAccess, authId)
	return true
}

// Error represents a response error object
type Error struct {

	// Code is a specific application error code (integer). It could be any
	// integer number: `200`, `396`, `2`, `3213254`, ...
	Code int `json:"code" bson:"code"`

	// Description is a human readable description for the error. It is a
	// specific application domain description, to developers.
	Description string `json:"description" bson:"description"`
}

// Request represents an Audit Request.
type Request struct {

	// Method is the HTTP verb
	Method string `json:"method" bson:"method"`

	// URI is the raw URL (the same one sent by the client)
	URI string `json:"uri" bson:"uri"`

	// Handler is the URL before replacing parameters
	Handler string `json:"handler" bson:"handler"`

	// Parameters is the URL parameters sent by client
	Parameters map[string]string `json:"parameters" bson:"parameters"`

	// Query is the URL query params
	Query map[string][]string `json:"query" bson:"query"`

	// Length is the client request body length
	Length int64 `json:"length" bson:"length"`
}

// Response represents an Audit response.
type Response struct {

	// StatusCode is the returned status code number.
	StatusCode int `json:"status_code" bson:"status_code"`

	// Length is the response body length.
	Length int64 `json:"length" bson:"length"`

	// Error represents an error object.
	Error *Error `json:"error" bson:"error"`
}

// Service identifies the process and other information related to the service.
type Service struct {

	// Name is the service name.
	Name string `json:"name" bson:"name"`

	// Version is the service version, for example: 1.2, 0.0.1, 7.3.
	Version string `json:"version" bson:"version"`

	// Commit is the short commit number to identify exactly the code base
	// that is being executed
	Commit string `json:"commit" bson:"commit"`
}
