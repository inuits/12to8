package acceptance_tests

import (
	"fmt"
	"testing"
	"time"
)

// TestTimesheetCurrentMonth Tests the lifecycle of a timesheet
// when we do not give any argument
func TestTimesheetCurrentMonth(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name: "List timesheets",
		Env:  userEnv,
		Args: []string{"list", "timesheets"},
	}).Run(t)
	(&CmdTestCase{
		Name:     "Create timesheet",
		Env:      userEnv,
		Args:     []string{"new", "timesheet"},
		OutLines: 1,
	}).Run(t)
	currentTs := fmt.Sprintf("%s %d", time.Now().Month(), time.Now().Year())
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[ACTIVE\\]", currentTs),
	}).Run(t)
	(&CmdTestCase{
		Name:     "Release Timesheet",
		Env:      userEnv,
		Args:     []string{"release", "timesheet", "-f"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
}

// TestTimesheetCurrentYear Tests the lifecycle of a timesheet
// when we give only the current month
func TestTimesheetCurrentYear(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name: "List timesheets",
		Env:  userEnv,
		Args: []string{"list", "timesheets"},
	}).Run(t)
	(&CmdTestCase{
		Name:     "Create timesheet",
		Env:      userEnv,
		Args:     []string{"new", "timesheet", "3"},
		OutLines: 1,
	}).Run(t)
	currentTs := fmt.Sprintf("March %d", time.Now().Year())
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[ACTIVE\\]", currentTs),
	}).Run(t)
	(&CmdTestCase{
		Name:     "Release Timesheet",
		Env:      userEnv,
		Args:     []string{"release", "timesheet", "3", "-f"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
}

// TestTimesheetFixedMonthYear Tests the lifecycle of a timesheet
// when we give only the month and the year
func TestTimesheetFixedMonthYear(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name: "List timesheets",
		Env:  userEnv,
		Args: []string{"list", "timesheets"},
	}).Run(t)
	(&CmdTestCase{
		Name:     "Create timesheet",
		Env:      userEnv,
		Args:     []string{"new", "timesheet", "8/2008"},
		OutLines: 1,
	}).Run(t)
	currentTs := "August 2008"
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[ACTIVE\\]", currentTs),
	}).Run(t)
	(&CmdTestCase{
		Name:     "Release Timesheet",
		Env:      userEnv,
		Args:     []string{"release", "timesheet", "8/2008", "-f"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
}

func TestReleaseTimesheetY(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "Create timesheet",
		Env:      userEnv,
		Args:     []string{"new", "timesheet"},
		OutLines: 1,
	}).Run(t)
	currentTs := fmt.Sprintf("%s %d", time.Now().Month(), time.Now().Year())
	(&CmdTestCase{
		Name:     "Release Timesheet",
		Env:      userEnv,
		Input:    "y",
		Args:     []string{"release", "timesheet"},
		OutLines: 2,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
}

func TestReleaseTimesheetYes(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "Create timesheet",
		Env:      userEnv,
		Args:     []string{"new", "timesheet"},
		OutLines: 1,
	}).Run(t)
	currentTs := fmt.Sprintf("%s %d", time.Now().Month(), time.Now().Year())
	(&CmdTestCase{
		Name:     "Release Timesheet",
		Env:      userEnv,
		Input:    "yes",
		Args:     []string{"release", "timesheet"},
		OutLines: 2,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[PENDING\\]", currentTs),
	}).Run(t)
}

func TestReleaseTimesheetNo(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "Create timesheet",
		Env:      userEnv,
		Args:     []string{"new", "timesheet"},
		OutLines: 1,
	}).Run(t)
	currentTs := fmt.Sprintf("%s %d", time.Now().Month(), time.Now().Year())
	(&CmdTestCase{
		Name:          "Release Timesheet",
		Env:           userEnv,
		Input:         "no",
		Args:          []string{"release", "timesheet"},
		OutLines:      2,
		ExpectFailure: true,
		OutRegex:      fmt.Sprintf("Aborted by user"),
	}).Run(t)
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[ACTIVE\\]", currentTs),
	}).Run(t)
}

func TestReleaseTimesheetNoInput(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "Create timesheet",
		Env:      userEnv,
		Args:     []string{"new", "timesheet"},
		OutLines: 1,
	}).Run(t)
	currentTs := fmt.Sprintf("%s %d", time.Now().Month(), time.Now().Year())
	(&CmdTestCase{
		Name:          "Release Timesheet",
		Env:           userEnv,
		Args:          []string{"release", "timesheet"},
		OutLines:      2,
		ExpectFailure: true,
		OutRegex:      fmt.Sprintf("Aborted by user"),
	}).Run(t)
	(&CmdTestCase{
		Name:     "List timesheets",
		Env:      userEnv,
		Args:     []string{"list", "timesheets"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[ACTIVE\\]", currentTs),
	}).Run(t)
}
