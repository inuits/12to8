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
	"time"
)

type Timesheet struct {
	Id           int    `json:"id"`
	Year         int    `json:"year"`
	Month        int    `json:"month"`
	DisplayLabel string `json:"display_label"`
	Status       string `json:"status"`
}

type TimesheetsList struct {
	Timesheets []Timesheet `json:"results"`
}

func (ts *TimesheetsList) Fetch(c Client) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_timesheets?page_size=9999", c.Endpoint))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(ts)
	if err != nil {
		return err
	}
	return nil
}

func (t *Timesheet) New(c Client) error {
	resp, err := c.PostRequest(fmt.Sprintf("%s/v1/my_timesheets/", c.Endpoint), t)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}
	return nil
}

// Get returns the timesheets from the server
func (t *Timesheet) Get(c Client) error {
	ts := &TimesheetsList{}
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_timesheets/?year=%d&month=%d", c.Endpoint, t.Year, t.Month))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(ts)
	if err != nil {
		return err
	}
	if len(ts.Timesheets) != 1 {
		return errors.New(fmt.Sprintf("Expected 1 timesheet, got %d", len(ts.Timesheets)))
	}
	*t = ts.Timesheets[0]
	return nil
}

func (t *Timesheet) Release(c Client) error {
	if t.Id == 0 {
		return errors.New("No ID for this timesheet")
	}
	t.Status = "PENDING"
	resp, err := c.PatchRequest(fmt.Sprintf("%s/v1/my_timesheets/%d/", c.Endpoint, t.Id), t)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TimesheetsList) PrettyPrint() {
	for _, t := range ts.Timesheets {
		t.PrettyPrint()
	}
}

func (t *Timesheet) Name() string {
	return fmt.Sprintf("%s %d", time.Month(t.Month), t.Year)
}

func (t *Timesheet) PrettyPrint() {
	fmt.Printf("%s [%s]\n", t.Name(), t.Status)
}
