// Copyright © 2017 Julien Pivotto <roidelapluie@inuits.eu>
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

package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PerformanceType int

const (
	ActivityPerformance PerformanceType = iota
	StandbyPerformance
)

func (p PerformanceType) String() string {
	switch p {
	case ActivityPerformance:
		return "activity"
	case StandbyPerformance:
		return "standby"
	}
	return fmt.Sprintf("PerformanceType(%d)", p)
}

func (p *PerformanceType) UnmarshalJSON(b []byte) error {
	var textValue string
	err := json.Unmarshal(b, &textValue)
	if err != nil {
		return err
	}
	if textValue == "ActivityPerformance" {
		*p = ActivityPerformance
	} else if textValue == "StandbyPerformance" {
		*p = StandbyPerformance
	} else {
		return errors.New("Can't UnmarshalJSON")
	}
	return nil
}
