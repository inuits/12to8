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

package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

func getMonthYearFromArg(arg string) (int, int, error) {
	if arg == "" {
		return int(time.Now().Month()), time.Now().Year(), nil
	}
	v := strings.Split(arg, "/")
	if len(v) > 2 {
		return 0, 0, fmt.Errorf("Too many / in %v", v)
	}
	var y int
	var err error
	if len(v) == 1 {
		y = time.Now().Year()
	} else {
		y, err = strconv.Atoi(v[1])
		if err != nil {
			return 0, 0, err
		}
	}
	m, err := strconv.Atoi(v[0])
	if err != nil {
		return 0, 0, err
	}
	if m > 12 || m < 1 {
		return 0, 0, fmt.Errorf("bad month: %v", v[0])
	}
	return m, y, nil
}

func validTimesheetArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return errors.New("takes at most one argument")
	}
	if len(args) == 1 {
		_, _, err := getMonthYearFromArg(args[0])
		if err != nil {
			return err
		}
	}
	return nil
}

func validTimesheetArgsWithColumns(cmd *cobra.Command, args []string) error {
	if err := validTimesheetArgs(cmd, args); err != nil {
		return err
	}
	var invalidColumns []string
	for _, column := range strings.Split(columns, ",") {
		if column == "" {
			continue
		}
		found := false
		for _, validColumn := range api.PerformancesColumns {
			if validColumn == column {
				found = true
			}
		}
		if !found {
			invalidColumns = append(invalidColumns, column)
		}
	}
	if len(invalidColumns) > 0 {
		return fmt.Errorf("invalid columns: %s", strings.Join(invalidColumns, ", "))
	}
	return nil
}

func validIDArg(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("takes exactly argument")
	}
	_, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	return nil
}

func validPerfAddArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 3 {
		return errors.New("takes at most 3 argument")
	}
	if len(args) < 2 {
		return errors.New("takes at least 2 arguments")
	}
	_, _, _, err := getDaysMonthYearFromArg(args[0])
	if err != nil {
		return err
	}
	_, err = strconv.ParseFloat(args[1], 64)
	if err != nil {
		return err
	}
	return nil
}

func getDaysMonthYearFromArg(arg string) (int, int, int, error) {
	if arg == "today" {
		return time.Now().Day(), int(time.Now().Month()), time.Now().Year(), nil
	}
	v := strings.Split(arg, "/")
	if len(v) > 3 {
		return 0, 0, 0, fmt.Errorf("Too many / in %v", v)
	}
	var m int
	var y int
	var err error
	if len(v) < 3 {
		y = time.Now().Year()
	} else {
		y, err = strconv.Atoi(v[2])
		if err != nil {
			return 0, 0, 0, err
		}
	}
	if len(v) < 2 {
		m = int(time.Now().Month())
	} else {
		m, err = strconv.Atoi(v[1])
		if err != nil {
			return 0, 0, 0, err
		}
		if m > 12 || m < 1 {
			return 0, 0, 0, fmt.Errorf("bad month: %v", v[1])
		}
	}
	d, err := strconv.Atoi(v[0])
	if err != nil {
		return 0, 0, 0, err
	}
	return d, m, y, nil
}

func fetchTimesheetFromArgs(args []string) *api.Timesheet {
	monthSpec := ""
	if len(args) == 1 {
		monthSpec = args[0]
	}
	month, year, err := getMonthYearFromArg(monthSpec)
	if err != nil {
		log.Fatal(err)
	}
	timesheet := &api.Timesheet{
		Month: month,
		Year:  year,
	}
	c := NewApiClient()
	err = timesheet.Get(c)
	if err != nil {
		log.Fatal(err)
	}
	return timesheet
}
