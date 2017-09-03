package acceptance_tests

import (
	"testing"
)

// TestListCompanies tests the list rates command
func TestListRates(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List rates",
		Env:      userEnv,
		Args:     []string{"list", "rates"},
		OutLines: 2,
		OutText: `1.00 [Normal Hours]
2.00 [Double]
`,
	}).Run(t)
}
