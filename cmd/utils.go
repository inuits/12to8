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
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func getMonthYearFromArg(arg string) (int, int, error) {
	if arg == "" {
		return int(time.Now().Month()), time.Now().Year(), nil
	}
	v := strings.Split(arg, "/")
	if len(v) > 2 {
		return 0, 0, errors.New(fmt.Sprintf("Too many / in %v", v))
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
		return 0, 0, errors.New(fmt.Sprintf("bad month: %v", v[0]))
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
