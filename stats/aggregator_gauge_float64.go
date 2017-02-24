// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package stats

import (
	"bytes"
	"fmt"
	"time"
)

// GaugeFloat64Stats records a gauge of float64 sample values.
type GaugeFloat64Stats struct {
	Value     float64
	TimeStamp time.Time
}

func (gs *GaugeFloat64Stats) String() string {
	if gs == nil {
		return "nil"
	}

	var buf bytes.Buffer
	buf.WriteString("  GaugeFloat64Stats{\n")
	fmt.Fprintf(&buf, "    Value: %v,\n", gs.Value)
	fmt.Fprintf(&buf, "    TimeStamp: %v,\n", gs.TimeStamp)
	buf.WriteString("  }")
	return buf.String()
}

// newGaugeAggregatorFloat64 creates a gaugeAggregator of type float64. For a
// single GaugeAggregationDescriptor it is expected to be called multiple
// times. Once for each unique set of tags.
func newGaugeAggregatorFloat64() *gaugeAggregatorFloat64 {
	return &gaugeAggregatorFloat64{
		gs: &GaugeFloat64Stats{},
	}
}

type gaugeAggregatorFloat64 struct {
	gs *GaugeFloat64Stats
}

func (ga *gaugeAggregatorFloat64) addSample(m Measurement, now time.Time) {
	ga.gs.Value = m.float64()
	ga.gs.TimeStamp = now
}

func (ga *gaugeAggregatorFloat64) retrieveCollected() *GaugeFloat64Stats {
	return ga.gs
}
