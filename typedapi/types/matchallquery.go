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

// MatchAllQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/query_dsl/MatchAllQuery.ts#L22-L22
type MatchAllQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
}

// MatchAllQueryBuilder holds MatchAllQuery struct and provides a builder API.
type MatchAllQueryBuilder struct {
	v *MatchAllQuery
}

// NewMatchAllQuery provides a builder for the MatchAllQuery struct.
func NewMatchAllQueryBuilder() *MatchAllQueryBuilder {
	r := MatchAllQueryBuilder{
		&MatchAllQuery{},
	}

	return &r
}

// Build finalize the chain and returns the MatchAllQuery struct
func (rb *MatchAllQueryBuilder) Build() MatchAllQuery {
	return *rb.v
}

func (rb *MatchAllQueryBuilder) Boost(boost float32) *MatchAllQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *MatchAllQueryBuilder) QueryName_(queryname_ string) *MatchAllQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}