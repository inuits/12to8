package acceptance_tests

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTopCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", ""}, []string{"completion", "list", "new", "release"})
}

func TestCompletionCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "completion", ""}, []string{"bash"})
}

// TestAddPerformanceCompletion ensures that we can add a performance quickly (one letter + tab for each subcommand)
func TestAddPerformanceCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "n"}, []string{"new"})
	testCompletion(t, nil, []string{"12to8", "new", "p"}, []string{"performance"})
}

func TestReleaseTimesheetCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "r"}, []string{"release"})
	testCompletion(t, nil, []string{"12to8", "release", ""}, []string{"timesheet"})
}

func TestNewTimesheetCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", "n"}, []string{"new"})
	testCompletion(t, nil, []string{"12to8", "new", "t"}, []string{"timesheet"})

	now := time.Now()
	year, month, _ := now.Date()
	var prevMonth int
	var prevYear int
	var nextMonth int
	var nextYear int
	if int(month) == 1 {
		prevMonth = 12
		prevYear = year - 1
	} else {
		prevMonth = int(month) - 1
		prevYear = year
	}
	if int(month) == 12 {
		nextMonth = 1
		nextYear = year + 1
	} else {
		nextMonth = int(month) + 1
		nextYear = year
	}
	possibleTimesheets := []string{}
	possibleTimesheets = append(possibleTimesheets, fmt.Sprintf("%02d/%d", prevMonth, prevYear))
	possibleTimesheets = append(possibleTimesheets, fmt.Sprintf("%02d/%d", month, year))
	possibleTimesheets = append(possibleTimesheets, fmt.Sprintf("%02d/%d", nextMonth, nextYear))
	testCompletion(t, nil, []string{"12to8", "new", "timesheet", ""}, possibleTimesheets)
}

func TestNewPerformanceContractCompletion(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "basic_projects")
	defer c.stop925r(t)
	testCompletion(t, c, []string{"12to8", "new", "pe"}, []string{"performance"})
	testCompletion(t, c, []string{"12to8", "new", "performance", "-c", ""}, []string{`"Go Consultancy [Python & Co]"`, `"Internal Stuff (c) [Golang Tech]"`})
	testCompletion(t, c, []string{"12to8", "new", "performance", "-c", `"`}, []string{"Go Consultancy [Python & Co]", "Internal Stuff (c) [Golang Tech]"})
	testCompletion(t, c, []string{"12to8", "new", "performance", "-c", `'`}, []string{"Go Consultancy [Python & Co]", "Internal Stuff (c) [Golang Tech]"})
}

func completionBashCode(cli []string) string {
	flatcli := strings.Replace(strings.Join(cli, " "), `"`, `\"`, -1)
	flatcli = strings.Replace(flatcli, "'", `\'`, -1)
	words := len(cli) - 1
	return fmt.Sprintf(`
export PATH=..:$PATH
. /usr/share/bash-completion/bash_completion
. <(12to8 completion bash)
COMP_WORDS=(%s)
COMP_CWORD=%d
COMP_LINE="%s"
COMP_POINT=${#COMP_LINE}
_xfunc 12to8 __start_12to8
printf '%%s\n' "${COMPREPLY[@]}"
`, flatcli, words, flatcli)
}

func testCompletion(t *testing.T, c *dockerId, cli []string, expected []string) {
	var expectedOut bytes.Buffer
	for _, expectedLine := range expected {
		expectedOut.WriteString(fmt.Sprintf("%s\n", expectedLine))
	}
	tc := &CmdTestCase{
		Name:     "Autocomplete",
		Cmd:      "bash",
		Args:     []string{"-c", completionBashCode(cli)},
		OutLines: len(expected),
		OutText:  expectedOut.String(),
	}
	if c != nil {
		tc.Env = append(RunAsUser, c.EndpointEnv())
	}
	tc.Run(t)
}
