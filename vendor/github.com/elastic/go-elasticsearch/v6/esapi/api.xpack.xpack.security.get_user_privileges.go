// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 6.8.8: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newXPackSecurityGetUserPrivilegesFunc(t Transport) XPackSecurityGetUserPrivileges {
	return func(o ...func(*XPackSecurityGetUserPrivilegesRequest)) (*Response, error) {
		var r = XPackSecurityGetUserPrivilegesRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// XPackSecurityGetUserPrivileges - https://www.elastic.co/guide/en/elasticsearch/reference/6.7/security-api-get-privileges.html
//
type XPackSecurityGetUserPrivileges func(o ...func(*XPackSecurityGetUserPrivilegesRequest)) (*Response, error)

// XPackSecurityGetUserPrivilegesRequest configures the X Pack Security Get User Privileges API request.
//
type XPackSecurityGetUserPrivilegesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XPackSecurityGetUserPrivilegesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_xpack/security/user/_privileges"))
	path.WriteString("/_xpack/security/user/_privileges")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f XPackSecurityGetUserPrivileges) WithContext(v context.Context) func(*XPackSecurityGetUserPrivilegesRequest) {
	return func(r *XPackSecurityGetUserPrivilegesRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XPackSecurityGetUserPrivileges) WithPretty() func(*XPackSecurityGetUserPrivilegesRequest) {
	return func(r *XPackSecurityGetUserPrivilegesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XPackSecurityGetUserPrivileges) WithHuman() func(*XPackSecurityGetUserPrivilegesRequest) {
	return func(r *XPackSecurityGetUserPrivilegesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XPackSecurityGetUserPrivileges) WithErrorTrace() func(*XPackSecurityGetUserPrivilegesRequest) {
	return func(r *XPackSecurityGetUserPrivilegesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XPackSecurityGetUserPrivileges) WithFilterPath(v ...string) func(*XPackSecurityGetUserPrivilegesRequest) {
	return func(r *XPackSecurityGetUserPrivilegesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f XPackSecurityGetUserPrivileges) WithHeader(h map[string]string) func(*XPackSecurityGetUserPrivilegesRequest) {
	return func(r *XPackSecurityGetUserPrivilegesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f XPackSecurityGetUserPrivileges) WithOpaqueID(s string) func(*XPackSecurityGetUserPrivilegesRequest) {
	return func(r *XPackSecurityGetUserPrivilegesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
