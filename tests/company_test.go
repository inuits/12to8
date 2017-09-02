package acceptance_tests

import (
	"testing"
)

// TestListCompanies tests the list companies command
func TestListCompanies(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List companies",
		Env:      userEnv,
		Args:     []string{"list", "companies"},
		OutLines: 3,
		OutRegex: "^Golang Tech \\[US\\]\nPhPhew \\[HK\\]\nPython & Co \\[LU\\]\n",
	}).Run(t)
}
