package pkg

import (
	"bytes"
	"fmt"

	"github.com/valyala/fasthttp"
)

// RequestWithFastHTTP http request
func RequestWithFastHTTP(uri string, reqBody []byte, headers *fasthttp.RequestHeader) ([]byte, error) {
	var body []byte
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(uri)
	if reqBody != nil {
		req.SetBody(reqBody)
	}
	if headers != nil {
		headers.CopyTo(&req.Header)
	} else {
		headers.SetMethod(fasthttp.MethodGet)
	}
	// fasthttp does not automatically request a gzipped response.
	// We must explicitly ask for it.
	if len(req.Header.Peek("Accept-Encoding")) < 2 {
		req.Header.Set("Accept-Encoding", "gzip")
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Perform the request
	err := fasthttp.Do(req, resp)
	if err != nil {
		return body, fmt.Errorf("Client get failed: %s", err)
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return body, fmt.Errorf("Expected status code %d but got %d", fasthttp.StatusOK, resp.StatusCode())
	}

	// Verify the content type
	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		return body, fmt.Errorf("Expected content type application/json but got %s", contentType)
	}

	// Do we need to decompress the response?
	contentEncoding := resp.Header.Peek("Content-Encoding")

	if bytes.EqualFold(contentEncoding, []byte("gzip")) {
		body, _ = resp.BodyGunzip()
	} else {
		body = resp.Body()
	}

	return body, nil
}
