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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// NodeInfoXpackSecurityAuthcRealms type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/nodes/info/types.ts#L250-L254
type NodeInfoXpackSecurityAuthcRealms struct {
	File   map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"file,omitempty"`
	Native map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"native,omitempty"`
	Pki    map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"pki,omitempty"`
}

// NodeInfoXpackSecurityAuthcRealmsBuilder holds NodeInfoXpackSecurityAuthcRealms struct and provides a builder API.
type NodeInfoXpackSecurityAuthcRealmsBuilder struct {
	v *NodeInfoXpackSecurityAuthcRealms
}

// NewNodeInfoXpackSecurityAuthcRealms provides a builder for the NodeInfoXpackSecurityAuthcRealms struct.
func NewNodeInfoXpackSecurityAuthcRealmsBuilder() *NodeInfoXpackSecurityAuthcRealmsBuilder {
	r := NodeInfoXpackSecurityAuthcRealmsBuilder{
		&NodeInfoXpackSecurityAuthcRealms{
			File:   make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, 0),
			Native: make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, 0),
			Pki:    make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoXpackSecurityAuthcRealms struct
func (rb *NodeInfoXpackSecurityAuthcRealmsBuilder) Build() NodeInfoXpackSecurityAuthcRealms {
	return *rb.v
}

func (rb *NodeInfoXpackSecurityAuthcRealmsBuilder) File(values map[string]*NodeInfoXpackSecurityAuthcRealmsStatusBuilder) *NodeInfoXpackSecurityAuthcRealmsBuilder {
	tmp := make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.File = tmp
	return rb
}

func (rb *NodeInfoXpackSecurityAuthcRealmsBuilder) Native(values map[string]*NodeInfoXpackSecurityAuthcRealmsStatusBuilder) *NodeInfoXpackSecurityAuthcRealmsBuilder {
	tmp := make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Native = tmp
	return rb
}

func (rb *NodeInfoXpackSecurityAuthcRealmsBuilder) Pki(values map[string]*NodeInfoXpackSecurityAuthcRealmsStatusBuilder) *NodeInfoXpackSecurityAuthcRealmsBuilder {
	tmp := make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Pki = tmp
	return rb
}