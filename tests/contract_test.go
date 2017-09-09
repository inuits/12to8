package tests

import (
	"testing"
)

// TestListContracts tests the list contracts command
func TestListContracts(t *testing.T) {
	c := &dockerID{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List contracts",
		Env:      userEnv,
		Args:     []string{"list", "contracts"},
		OutLines: 2,
		OutRegex: "^Go Consultancy \\[Python & Co\\]\nInternal Stuff \\(c\\) \\[Golang Tech\\]\n$",
	}).Run(t)
}
