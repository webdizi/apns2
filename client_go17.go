// +build go1.7

package apns2

import (
	"context"
	"net/http"
	"net/http/httputil"

	log "github.com/sirupsen/logrus"
)

// A Context carries a deadline, a cancellation signal, and other values across
// API boundaries.
//
// Context's methods may be called by multiple goroutines simultaneously.
type Context interface {
	context.Context
}

func (c *Client) requestWithContext(ctx Context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	reqDump, _ := httputil.DumpRequestOut(req, true)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	respDump, _ := httputil.DumpResponse(resp, true)
	log.WithFields(log.Fields{
		"obj": log.Fields{
			"request":  string(reqDump),
			"response": string(respDump),
		},
	}).Debugf("send request to %s", req.URL.Host)

	return resp, nil
}
