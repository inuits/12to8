package tests

import (
	"testing"
)

// TestListPerformance tests that we can list performances and that the list keeps order
func TestListPerformance(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List performances",
		Env:      userEnv,
		Args:     []string{"list", "performances", "09/2017"},
		OutLines: 24,
		OutText: `+-----+--------------+-------------+-------+
| DAY |   CONTRACT   | DESCRIPTION |   H   |
+-----+--------------+-------------+-------+
|   4 | Consult      |             |  8.00 |
|   6 | Consult      |             |  8.00 |
|   7 | Consult      |             |  8.00 |
|  10 | Webby        |             |  8.00 |
|  11 | Webby        |             |  8.00 |
|  12 | Webby        |             |  8.00 |
|  12 | keep it up!! |             |  8.00 |
|  14 | Webby        |             |  8.00 |
|  17 | Consult      |             |  8.00 |
|  18 | Consult      |             |  8.00 |
|  19 | Consult      |             |  8.00 |
|  20 | Consult      |             |  8.00 |
|  21 | Consult      |             |  3.00 |
|  25 | Consult      |             |  4.00 |
|  25 | Consult      |             | 12.00 |
|  25 | Webby        |             |  4.00 |
|  26 | Consult      |             |  4.00 |
|  27 | Consult      | YES         |  4.00 |
|  27 | Webby        |             |  4.00 |
|  28 | Webby        |             |  8.00 |
+-----+--------------+-------------+-------+
`,
	}).Run(t)
}

// TestListPerformanceWithColums tests that we can list performances
// even with specific columns
func TestListPerformanceWithColumns(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List performances",
		Env:      userEnv,
		Args:     []string{"list", "performances", "09/2017", "-C", "duration,day,contract,multiplier,type,description,id,day"},
		OutLines: 24,
		OutText: `+-------+-----+--------------+------+----------+-------------+----+-----+
|   H   | DAY |   CONTRACT   |  X   |   KIND   | DESCRIPTION | ID | DAY |
+-------+-----+--------------+------+----------+-------------+----+-----+
|  8.00 |   4 | Consult      | 1.00 | activity |             |  1 |   4 |
|  8.00 |   6 | Consult      | 1.00 | activity |             |  2 |   6 |
|  8.00 |   7 | Consult      | 1.00 | activity |             |  3 |   7 |
|  8.00 |  10 | Webby        | 1.00 | activity |             |  4 |  10 |
|  8.00 |  11 | Webby        | 1.00 | activity |             |  5 |  11 |
|  8.00 |  12 | Webby        | 1.00 | activity |             |  6 |  12 |
|  8.00 |  12 | keep it up!! | 1.00 | activity |             |  7 |  12 |
|  8.00 |  14 | Webby        | 1.00 | activity |             |  8 |  14 |
|  8.00 |  17 | Consult      | 1.00 | activity |             | 18 |  17 |
|  8.00 |  18 | Consult      | 1.00 | activity |             | 19 |  18 |
|  8.00 |  19 | Consult      | 1.00 | activity |             | 20 |  19 |
|  8.00 |  20 | Consult      | 1.00 | activity |             | 17 |  20 |
|  3.00 |  21 | Consult      | 1.00 | activity |             | 16 |  21 |
|  4.00 |  25 | Consult      | 1.00 | activity |             | 12 |  25 |
| 12.00 |  25 | Consult      | 1.00 | activity |             | 15 |  25 |
|  4.00 |  25 | Webby        | 1.00 | activity |             | 11 |  25 |
|  4.00 |  26 | Consult      | 1.00 | activity |             | 14 |  26 |
|  4.00 |  27 | Consult      | 1.00 | activity | YES         | 21 |  27 |
|  4.00 |  27 | Webby        | 1.00 | activity |             | 10 |  27 |
|  8.00 |  28 | Webby        | 1.00 | activity |             |  9 |  28 |
+-------+-----+--------------+------+----------+-------------+----+-----+
`,
	}).Run(t)
}

// TestListPerformanceWithWrongColums tests that we can not list performances
// when the columns list is not correct
func TestListPerformanceWithWrongColumns(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:          "List performances",
		Env:           userEnv,
		Args:          []string{"list", "performances", "09/2017", "-C", "duration,day,nonexisting,extra,,,"},
		ExpectFailure: true,
		OutLines:      1,
		ErrLines:      17,
		ErrText: `Error: invalid columns: nonexisting, extra
Usage:
  12to8 list performances [MM[/YYYY]] [flags]

Flags:
  -C, --columns string   comma-separated columns to be displayed (day,contract,description,duration,multiplier,type,id) (default "day,contract,description,duration")
  -h, --help             help for performances
  -P, --porcelain        porcelain (usable in scripts) output

Global Flags:
      --cache string      config file (default is $HOME/.cache/12to8) (default "~/.cache/12to8")
      --config string     config file (default is $HOME/.12to8.yaml)
  -e, --endpoint string   API endpoint (without /v1)
      --no-cache          do not use cache, fetch from 925r as needed
  -p, --password string   password
  -u, --user string       username

`,
	}).Run(t)
}

// TestListPerformancePorcelain tests that we can list performances with
// an output suitable for scripts
func TestListPerformancePorcelain(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List performances",
		Env:      userEnv,
		Args:     []string{"list", "performances", "09/2017", "-P"},
		OutLines: 20,
		OutText: `1,04/09/2017,,Consult,8.00,1.00,activity
2,06/09/2017,,Consult,8.00,1.00,activity
3,07/09/2017,,Consult,8.00,1.00,activity
4,10/09/2017,,Webby,8.00,1.00,activity
5,11/09/2017,,Webby,8.00,1.00,activity
6,12/09/2017,,Webby,8.00,1.00,activity
7,12/09/2017,,keep it up!!,8.00,1.00,activity
8,14/09/2017,,Webby,8.00,1.00,activity
18,17/09/2017,,Consult,8.00,1.00,activity
19,18/09/2017,,Consult,8.00,1.00,activity
20,19/09/2017,,Consult,8.00,1.00,activity
17,20/09/2017,,Consult,8.00,1.00,activity
16,21/09/2017,,Consult,3.00,1.00,activity
12,25/09/2017,,Consult,4.00,1.00,activity
15,25/09/2017,,Consult,12.00,1.00,activity
11,25/09/2017,,Webby,4.00,1.00,activity
14,26/09/2017,,Consult,4.00,1.00,activity
21,27/09/2017,YES,Consult,4.00,1.00,activity
10,27/09/2017,,Webby,4.00,1.00,activity
9,28/09/2017,,Webby,8.00,1.00,activity
`,
	}).Run(t)
}
