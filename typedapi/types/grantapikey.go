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

	"encoding/json"
)

// GrantApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/security/grant_api_key/types.ts#L25-L32
type GrantApiKey struct {
	Expiration      *string                     `json:"expiration,omitempty"`
	Metadata        Metadata                    `json:"metadata,omitempty"`
	Name            string                      `json:"name"`
	RoleDescriptors []map[string]RoleDescriptor `json:"role_descriptors,omitempty"`
}

func (s *GrantApiKey) UnmarshalJSON(data []byte) error {

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

		case "expiration":
			if err := dec.Decode(&s.Expiration); err != nil {
				return err
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "role_descriptors":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]RoleDescriptor, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.RoleDescriptors = append(s.RoleDescriptors, o)
			case '[':
				o := make([]map[string]RoleDescriptor, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.RoleDescriptors = o
			}

		}
	}
	return nil
}

// NewGrantApiKey returns a GrantApiKey.
func NewGrantApiKey() *GrantApiKey {
	r := &GrantApiKey{}

	return r
}
