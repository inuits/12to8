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
		OutText: `+-----+--------------+-------------+-------+------+----------+
| DAY |   PROJECT    | DESCRIPTION |   H   | RATE |   TYPE   |
+-----+--------------+-------------+-------+------+----------+
|   4 | Consult      |             |  8.00 | 1.00 | activity |
|   6 | Consult      |             |  8.00 | 1.00 | activity |
|   7 | Consult      |             |  8.00 | 1.00 | activity |
|  10 | Webby        |             |  8.00 | 1.00 | activity |
|  11 | Webby        |             |  8.00 | 1.00 | activity |
|  12 | Webby        |             |  8.00 | 1.00 | activity |
|  12 | keep it up!! |             |  8.00 | 1.00 | activity |
|  14 | Webby        |             |  8.00 | 1.00 | activity |
|  17 | Consult      |             |  8.00 | 1.00 | activity |
|  18 | Consult      |             |  8.00 | 1.00 | activity |
|  19 | Consult      |             |  8.00 | 1.00 | activity |
|  20 | Consult      |             |  8.00 | 1.00 | activity |
|  21 | Consult      |             |  3.00 | 1.00 | activity |
|  25 | Consult      |             |  4.00 | 1.00 | activity |
|  25 | Consult      |             | 12.00 | 1.00 | activity |
|  25 | Webby        |             |  4.00 | 1.00 | activity |
|  26 | Consult      |             |  4.00 | 1.00 | activity |
|  27 | Consult      | YES         |  4.00 | 1.00 | activity |
|  27 | Webby        |             |  4.00 | 1.00 | activity |
|  28 | Webby        |             |  8.00 | 1.00 | activity |
+-----+--------------+-------------+-------+------+----------+
`,
	}).Run(t)
}
