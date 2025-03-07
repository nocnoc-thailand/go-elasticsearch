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

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// ReportingEmailAttachment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/watcher/_types/Actions.ts#L224-L232
type ReportingEmailAttachment struct {
	Inline   *bool                       `json:"inline,omitempty"`
	Interval Duration                    `json:"interval,omitempty"`
	Request  *HttpInputRequestDefinition `json:"request,omitempty"`
	Retries  *int                        `json:"retries,omitempty"`
	Url      string                      `json:"url"`
}

func (s *ReportingEmailAttachment) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "inline":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Inline = &value
			case bool:
				s.Inline = &v
			}

		case "interval":
			if err := dec.Decode(&s.Interval); err != nil {
				return err
			}

		case "request":
			if err := dec.Decode(&s.Request); err != nil {
				return err
			}

		case "retries":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Retries = &value
			case float64:
				f := int(v)
				s.Retries = &f
			}

		case "url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Url = o

		}
	}
	return nil
}

// NewReportingEmailAttachment returns a ReportingEmailAttachment.
func NewReportingEmailAttachment() *ReportingEmailAttachment {
	r := &ReportingEmailAttachment{}

	return r
}
