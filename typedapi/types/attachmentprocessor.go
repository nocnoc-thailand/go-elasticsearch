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

// AttachmentProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/ingest/_types/Processors.ts#L96-L104
type AttachmentProcessor struct {
	Description       *string              `json:"description,omitempty"`
	Field             string               `json:"field"`
	If                *string              `json:"if,omitempty"`
	IgnoreFailure     *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing     *bool                `json:"ignore_missing,omitempty"`
	IndexedChars      *int64               `json:"indexed_chars,omitempty"`
	IndexedCharsField *string              `json:"indexed_chars_field,omitempty"`
	OnFailure         []ProcessorContainer `json:"on_failure,omitempty"`
	Properties        []string             `json:"properties,omitempty"`
	ResourceName      *string              `json:"resource_name,omitempty"`
	Tag               *string              `json:"tag,omitempty"`
	TargetField       *string              `json:"target_field,omitempty"`
}

func (s *AttachmentProcessor) UnmarshalJSON(data []byte) error {

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

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Description = &o

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "if":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.If = &o

		case "ignore_failure":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreFailure = &value
			case bool:
				s.IgnoreFailure = &v
			}

		case "ignore_missing":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreMissing = &value
			case bool:
				s.IgnoreMissing = &v
			}

		case "indexed_chars":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.IndexedChars = &value
			case float64:
				f := int64(v)
				s.IndexedChars = &f
			}

		case "indexed_chars_field":
			if err := dec.Decode(&s.IndexedCharsField); err != nil {
				return err
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return err
			}

		case "properties":
			if err := dec.Decode(&s.Properties); err != nil {
				return err
			}

		case "resource_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.ResourceName = &o

		case "tag":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Tag = &o

		case "target_field":
			if err := dec.Decode(&s.TargetField); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAttachmentProcessor returns a AttachmentProcessor.
func NewAttachmentProcessor() *AttachmentProcessor {
	r := &AttachmentProcessor{}

	return r
}
