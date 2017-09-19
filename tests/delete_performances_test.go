package tests

import (
	"testing"
)

func TestDeletePerformance(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:          "Delete Performance",
		Env:           userEnv,
		Args:          []string{"delete", "performance", "2"},
		OutLines:      2,
		ExpectFailure: true,
		OutText: `2 - 06/09/2017 8.00h  [Consult]
Are you sure you want to delete that performance? [y/N] Aborted by user
`,
	}).Run(t)
}

func TestDeletePerformanceY(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "Delete Performance",
		Env:      userEnv,
		Input:    "y",
		Args:     []string{"delete", "performance", "2"},
		OutLines: 1,
		OutText: `2 - 06/09/2017 8.00h  [Consult]
Are you sure you want to delete that performance? [y/N] `,
	}).Run(t)
	(&CmdTestCase{
		Name:          "Delete Performance (already deleted)",
		Env:           userEnv,
		Args:          []string{"delete", "performance", "2"},
		OutLines:      2,
		ExpectFailure: true,
		OutRegex:      `^Received 404, expecting 200 status code while fetching .*\n{"detail":"Not found."}`,
	}).Run(t)
}

func TestDeletePerformanceYes(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "Delete Performance",
		Env:      userEnv,
		Input:    "yes",
		Args:     []string{"delete", "performance", "2"},
		OutLines: 1,
		OutText: `2 - 06/09/2017 8.00h  [Consult]
Are you sure you want to delete that performance? [y/N] `,
	}).Run(t)
	(&CmdTestCase{
		Name:          "Delete Performance (already deleted)",
		Env:           userEnv,
		Args:          []string{"delete", "performance", "2"},
		OutLines:      2,
		ExpectFailure: true,
		OutRegex:      `^Received 404, expecting 200 status code while fetching .*\n{"detail":"Not found."}`,
	}).Run(t)
}

func TestDeletePerformanceNo(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	testCase := &CmdTestCase{
		Name:          "Delete Performance",
		Env:           userEnv,
		Input:         "n",
		Args:          []string{"delete", "performance", "2"},
		OutLines:      2,
		ExpectFailure: true,
		OutText: `2 - 06/09/2017 8.00h  [Consult]
Are you sure you want to delete that performance? [y/N] Aborted by user
`,
	}

	testCase.Run(t)
	testCase.Run(t)
}

func TestDeletePerformanceForce(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:    "Delete Performance",
		Env:     userEnv,
		Args:    []string{"delete", "performance", "2", "-f"},
		OutText: "",
	}).Run(t)
	(&CmdTestCase{
		Name:          "Delete Performance (already deleted)",
		Env:           userEnv,
		Args:          []string{"delete", "performance", "2"},
		OutLines:      2,
		ExpectFailure: true,
		OutRegex:      `^Received 404, expecting 200 status code while fetching .*\n{"detail":"Not found."}`,
	}).Run(t)
}
