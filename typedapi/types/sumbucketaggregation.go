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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// SumBucketAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/_types/aggregations/pipeline.ts#L286-L286
type SumBucketAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath BucketsPath          `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        Metadata             `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
}

func (s *SumBucketAggregation) UnmarshalJSON(data []byte) error {

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

		case "buckets_path":
			if err := dec.Decode(&s.BucketsPath); err != nil {
				return err
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Format = &o

		case "gap_policy":
			if err := dec.Decode(&s.GapPolicy); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Name = &o

		}
	}
	return nil
}

// NewSumBucketAggregation returns a SumBucketAggregation.
func NewSumBucketAggregation() *SumBucketAggregation {
	r := &SumBucketAggregation{}

	return r
}
