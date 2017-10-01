package tests

import (
	"testing"
)

// TestListLeaveTypes tests that we can list leave types
func TestListLeaveTypes(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List Leave Type",
		Env:      userEnv,
		Args:     []string{"list", "leave_types"},
		OutLines: 3,
		OutText: `NORMAL
NOT PAID
sick :(
`,
	}).Run(t)
}
