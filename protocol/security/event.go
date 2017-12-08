// Code generated by cdpgen. DO NOT EDIT.

package security

import (
	"github.com/mafredri/cdp/rpcc"
)

// CertificateErrorClient is a client for CertificateError events. There is a certificate error. If overriding certificate errors is enabled, then it should be
// handled with the handleCertificateError command. Note: this event does not fire if the
// certificate error has been allowed internally.
type CertificateErrorClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*CertificateErrorReply, error)
	rpcc.Stream
}

// CertificateErrorReply is the reply for CertificateError events.
type CertificateErrorReply struct {
	EventID    int    `json:"eventId"`    // The ID of the event.
	ErrorType  string `json:"errorType"`  // The type of the error.
	RequestURL string `json:"requestURL"` // The url that was requested.
}

// StateChangedClient is a client for SecurityStateChanged events. The security state of the page changed.
type StateChangedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*StateChangedReply, error)
	rpcc.Stream
}

// StateChangedReply is the reply for SecurityStateChanged events.
type StateChangedReply struct {
	SecurityState         State                 `json:"securityState"`         // Security state.
	SchemeIsCryptographic bool                  `json:"schemeIsCryptographic"` // True if the page was loaded over cryptographic transport such as HTTPS.
	Explanations          []StateExplanation    `json:"explanations"`          // List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
	InsecureContentStatus InsecureContentStatus `json:"insecureContentStatus"` // Information about insecure content on the page.
	Summary               *string               `json:"summary,omitempty"`     // Overrides user-visible description of the state.
}
