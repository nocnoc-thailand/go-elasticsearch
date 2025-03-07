// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/899364a63e7415b60033ddd49d50a30369da26d7

// Allows to manually change the allocation of individual shards in the cluster.
package reroute

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Reroute struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int
}

// NewReroute type alias for index.
type NewReroute func() *Reroute

// NewRerouteFunc returns a new instance of Reroute with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRerouteFunc(tp elastictransport.Interface) NewReroute {
	return func() *Reroute {
		n := New(tp)

		return n
	}
}

// Allows to manually change the allocation of individual shards in the cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cluster-reroute.html
func New(tp elastictransport.Interface) *Reroute {
	r := &Reroute{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Reroute) Raw(raw io.Reader) *Reroute {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Reroute) Request(req *Request) *Reroute {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Reroute) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Reroute: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("reroute")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Reroute) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Reroute query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a reroute.Response
func (r Reroute) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	return nil, errorResponse
}

// Header set a key, value pair in the Reroute headers map.
func (r *Reroute) Header(key, value string) *Reroute {
	r.headers.Set(key, value)

	return r
}

// DryRun If true, then the request simulates the operation only and returns the
// resulting state.
// API name: dry_run
func (r *Reroute) DryRun(b bool) *Reroute {
	r.values.Set("dry_run", strconv.FormatBool(b))

	return r
}

// Explain If true, then the response contains an explanation of why the commands can or
// cannot be executed.
// API name: explain
func (r *Reroute) Explain(b bool) *Reroute {
	r.values.Set("explain", strconv.FormatBool(b))

	return r
}

// Metric Limits the information returned to the specified metrics.
// API name: metric
func (r *Reroute) Metric(v string) *Reroute {
	r.values.Set("metric", v)

	return r
}

// RetryFailed If true, then retries allocation of shards that are blocked due to too many
// subsequent allocation failures.
// API name: retry_failed
func (r *Reroute) RetryFailed(b bool) *Reroute {
	r.values.Set("retry_failed", strconv.FormatBool(b))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *Reroute) MasterTimeout(v string) *Reroute {
	r.values.Set("master_timeout", v)

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *Reroute) Timeout(v string) *Reroute {
	r.values.Set("timeout", v)

	return r
}
