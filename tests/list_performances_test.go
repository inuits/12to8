package acceptance_tests

import (
	"testing"
)

// TestListPerformance tests that we can list performances and that the list keeps order
func TestListPerformance(t *testing.T) {
	c := &dockerId{}
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
	c := &dockerId{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List performances",
		Env:      userEnv,
		Args:     []string{"list", "performances", "09/2017", "-C", "duration,day,contract,multiplier,type,description,day"},
		OutLines: 24,
		OutText: `+-------+-----+--------------+------+----------+-------------+-----+
|   H   | DAY |   CONTRACT   |  X   |   KIND   | DESCRIPTION | DAY |
+-------+-----+--------------+------+----------+-------------+-----+
|  8.00 |   4 | Consult      | 1.00 | activity |             |   4 |
|  8.00 |   6 | Consult      | 1.00 | activity |             |   6 |
|  8.00 |   7 | Consult      | 1.00 | activity |             |   7 |
|  8.00 |  10 | Webby        | 1.00 | activity |             |  10 |
|  8.00 |  11 | Webby        | 1.00 | activity |             |  11 |
|  8.00 |  12 | Webby        | 1.00 | activity |             |  12 |
|  8.00 |  12 | keep it up!! | 1.00 | activity |             |  12 |
|  8.00 |  14 | Webby        | 1.00 | activity |             |  14 |
|  8.00 |  17 | Consult      | 1.00 | activity |             |  17 |
|  8.00 |  18 | Consult      | 1.00 | activity |             |  18 |
|  8.00 |  19 | Consult      | 1.00 | activity |             |  19 |
|  8.00 |  20 | Consult      | 1.00 | activity |             |  20 |
|  3.00 |  21 | Consult      | 1.00 | activity |             |  21 |
|  4.00 |  25 | Consult      | 1.00 | activity |             |  25 |
| 12.00 |  25 | Consult      | 1.00 | activity |             |  25 |
|  4.00 |  25 | Webby        | 1.00 | activity |             |  25 |
|  4.00 |  26 | Consult      | 1.00 | activity |             |  26 |
|  4.00 |  27 | Consult      | 1.00 | activity | YES         |  27 |
|  4.00 |  27 | Webby        | 1.00 | activity |             |  27 |
|  8.00 |  28 | Webby        | 1.00 | activity |             |  28 |
+-------+-----+--------------+------+----------+-------------+-----+
`,
	}).Run(t)
}

// TestListPerformanceWithWrongColums tests that we can not list performances
// when the columns list is not correct
func TestListPerformanceWithWrongColumns(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:          "List performances",
		Env:           userEnv,
		Args:          []string{"list", "performances", "09/2017", "-C", "duration,day,nonexisting,extra,,,"},
		ExpectFailure: true,
		OutLines:      1,
		ErrLines:      14,
		ErrText: `Error: invalid columns: nonexisting, extra
Usage:
  12to8 list performances [MM[/YYYY]] [flags]

Flags:
  -h, --help   help for performances

Global Flags:
  -C, --columns string    comma-separated columns to be displayed (day,contract,description,duration,multiplier,type) (default "day,contract,description,duration")
      --config string     config file (default is $HOME/.12to8.yaml)
  -e, --endpoint string   API endpoint (without /v1)
  -p, --password string   password
  -u, --user string       username

`,
	}).Run(t)
}
