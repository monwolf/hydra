// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOauth2TokenParams creates a new Oauth2TokenParams object
// with the default values initialized.
func NewOauth2TokenParams() *Oauth2TokenParams {
	var ()
	return &Oauth2TokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOauth2TokenParamsWithTimeout creates a new Oauth2TokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOauth2TokenParamsWithTimeout(timeout time.Duration) *Oauth2TokenParams {
	var ()
	return &Oauth2TokenParams{

		timeout: timeout,
	}
}

// NewOauth2TokenParamsWithContext creates a new Oauth2TokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewOauth2TokenParamsWithContext(ctx context.Context) *Oauth2TokenParams {
	var ()
	return &Oauth2TokenParams{

		Context: ctx,
	}
}

// NewOauth2TokenParamsWithHTTPClient creates a new Oauth2TokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOauth2TokenParamsWithHTTPClient(client *http.Client) *Oauth2TokenParams {
	var ()
	return &Oauth2TokenParams{
		HTTPClient: client,
	}
}

/*Oauth2TokenParams contains all the parameters to send to the API endpoint
for the oauth2 token operation typically these are written to a http.Request
*/
type Oauth2TokenParams struct {

	/*ClientID*/
	ClientID *string
	/*Code*/
	Code *string
	/*GrantType*/
	GrantType string
	/*RedirectURI*/
	RedirectURI *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the oauth2 token params
func (o *Oauth2TokenParams) WithTimeout(timeout time.Duration) *Oauth2TokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oauth2 token params
func (o *Oauth2TokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oauth2 token params
func (o *Oauth2TokenParams) WithContext(ctx context.Context) *Oauth2TokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oauth2 token params
func (o *Oauth2TokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oauth2 token params
func (o *Oauth2TokenParams) WithHTTPClient(client *http.Client) *Oauth2TokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oauth2 token params
func (o *Oauth2TokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the oauth2 token params
func (o *Oauth2TokenParams) WithClientID(clientID *string) *Oauth2TokenParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the oauth2 token params
func (o *Oauth2TokenParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithCode adds the code to the oauth2 token params
func (o *Oauth2TokenParams) WithCode(code *string) *Oauth2TokenParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the oauth2 token params
func (o *Oauth2TokenParams) SetCode(code *string) {
	o.Code = code
}

// WithGrantType adds the grantType to the oauth2 token params
func (o *Oauth2TokenParams) WithGrantType(grantType string) *Oauth2TokenParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the oauth2 token params
func (o *Oauth2TokenParams) SetGrantType(grantType string) {
	o.GrantType = grantType
}

// WithRedirectURI adds the redirectURI to the oauth2 token params
func (o *Oauth2TokenParams) WithRedirectURI(redirectURI *string) *Oauth2TokenParams {
	o.SetRedirectURI(redirectURI)
	return o
}

// SetRedirectURI adds the redirectUri to the oauth2 token params
func (o *Oauth2TokenParams) SetRedirectURI(redirectURI *string) {
	o.RedirectURI = redirectURI
}

// WriteToRequest writes these params to a swagger request
func (o *Oauth2TokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// form param client_id
		var frClientID string
		if o.ClientID != nil {
			frClientID = *o.ClientID
		}
		fClientID := frClientID
		if fClientID != "" {
			if err := r.SetFormParam("client_id", fClientID); err != nil {
				return err
			}
		}

	}

	if o.Code != nil {

		// form param code
		var frCode string
		if o.Code != nil {
			frCode = *o.Code
		}
		fCode := frCode
		if fCode != "" {
			if err := r.SetFormParam("code", fCode); err != nil {
				return err
			}
		}

	}

	// form param grant_type
	frGrantType := o.GrantType
	fGrantType := frGrantType
	if fGrantType != "" {
		if err := r.SetFormParam("grant_type", fGrantType); err != nil {
			return err
		}
	}

	if o.RedirectURI != nil {

		// form param redirect_uri
		var frRedirectURI string
		if o.RedirectURI != nil {
			frRedirectURI = *o.RedirectURI
		}
		fRedirectURI := frRedirectURI
		if fRedirectURI != "" {
			if err := r.SetFormParam("redirect_uri", fRedirectURI); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}