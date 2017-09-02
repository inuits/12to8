package acceptance_tests

import (
	"testing"
)

// TestListUsers tests the list users command
func TestListUsers(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List users",
		Env:      userEnv,
		Args:     []string{"list", "users"},
		OutLines: 2,
		OutRegex: `^ "user" \n "admin" \n$`,
	}).Run(t)
}
