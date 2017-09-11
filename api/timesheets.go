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

// Timesheet represents a ninetofiver timesheet
type Timesheet struct {
	ID           int    `json:"id"`
	Year         int    `json:"year"`
	Month        int    `json:"month"`
	DisplayLabel string `json:"display_label"`
	Status       string `json:"status"`
}

// TimesheetsList represents a list of ninetofiver timesheet
type TimesheetsList struct {
	Timesheets []Timesheet `json:"results"`
}

// New creates a new timesheet on the server
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

func (ts *TimesheetsList) apiURL() string {
	return "v1/my_timesheets"
}

func (ts *TimesheetsList) slug() string {
	return "timesheets"
}

func (ts *TimesheetsList) augment() error {
	return nil
}

func (ts *TimesheetsList) isEmpty() bool {
	return len(ts.Timesheets) == 0
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
		return fmt.Errorf("Expected 1 timesheet, got %d", len(ts.Timesheets))
	}
	*t = ts.Timesheets[0]
	return nil
}

// GetByID returns the timesheet from the server by its id
func (t *Timesheet) GetByID(c Client) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_timesheets/%d/", c.Endpoint, t.ID))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}
	return nil
}

// Release sends a request to ninetofiver to mark the timesheet as PENDING.
func (t *Timesheet) Release(c Client) error {
	if t.ID == 0 {
		return errors.New("No ID for this timesheet")
	}
	t.Status = "PENDING"
	resp, err := c.PatchRequest(fmt.Sprintf("%s/v1/my_timesheets/%d/", c.Endpoint, t.ID), t)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}
	return nil
}

// PrettyPrint prints timesheets in a nice way to the console
func (ts *TimesheetsList) PrettyPrint() {
	for _, t := range ts.Timesheets {
		t.PrettyPrint()
	}
}

// Name returns the timesheet name (English month name + year)
func (t *Timesheet) Name() string {
	return fmt.Sprintf("%s %d", time.Month(t.Month), t.Year)
}

// PrettyPrint prints timesheet in a nice way to the console
func (t *Timesheet) PrettyPrint() {
	fmt.Printf("%s [%s]\n", t.Name(), t.Status)
}
