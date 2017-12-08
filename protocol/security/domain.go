// Code generated by cdpgen. DO NOT EDIT.

// Package security implements the Security domain. Security
package security

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Security domain. Security
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Security domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Disable invokes the Security method. Disables tracking security state changes.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Security.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Security", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the Security method. Enables tracking security state changes.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Security.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Security", Op: "Enable", Err: err}
	}
	return
}

// HandleCertificateError invokes the Security method. Handles a certificate error that fired a certificateError event.
func (d *domainClient) HandleCertificateError(ctx context.Context, args *HandleCertificateErrorArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Security.handleCertificateError", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Security.handleCertificateError", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Security", Op: "HandleCertificateError", Err: err}
	}
	return
}

// SetOverrideCertificateErrors invokes the Security method. Enable/disable overriding certificate errors. If enabled, all certificate error events need to
// be handled by the DevTools client and should be answered with handleCertificateError commands.
func (d *domainClient) SetOverrideCertificateErrors(ctx context.Context, args *SetOverrideCertificateErrorsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Security.setOverrideCertificateErrors", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Security.setOverrideCertificateErrors", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Security", Op: "SetOverrideCertificateErrors", Err: err}
	}
	return
}

func (d *domainClient) CertificateError(ctx context.Context) (CertificateErrorClient, error) {
	s, err := rpcc.NewStream(ctx, "Security.certificateError", d.conn)
	if err != nil {
		return nil, err
	}
	return &certificateErrorClient{Stream: s}, nil
}

type certificateErrorClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *certificateErrorClient) GetStream() rpcc.Stream { return c.Stream }

func (c *certificateErrorClient) Recv() (*CertificateErrorReply, error) {
	event := new(CertificateErrorReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Security", Op: "CertificateError Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) SecurityStateChanged(ctx context.Context) (StateChangedClient, error) {
	s, err := rpcc.NewStream(ctx, "Security.securityStateChanged", d.conn)
	if err != nil {
		return nil, err
	}
	return &stateChangedClient{Stream: s}, nil
}

type stateChangedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *stateChangedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *stateChangedClient) Recv() (*StateChangedReply, error) {
	event := new(StateChangedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Security", Op: "SecurityStateChanged Recv", Err: err}
	}
	return event, nil
}
