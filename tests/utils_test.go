package tests

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
	"time"
)

var RunAsAdmin []string
var RunAsUser []string
var DefaultCmd string
var fixtures []string

type CmdTestCase struct {
	Name          string
	ExpectFailure bool
	Cmd           string
	Args          []string
	Input         string
	OutLines      int
	ErrLines      int
	OutRegex      string
	OutText       string
	ErrRegex      string
	ErrText       string
	Env           []string
}

// TestMain pulls the docker image,
// imports the fixtures, then commits them.
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	flag.Parse()
	for _, f := range fixtures {
		createContainerWithFixture(f)
	}
}

func shutdown() {
	for _, f := range fixtures {
		deleteContainerWithFixture(f)
	}
}

func (t *CmdTestCase) Run(test *testing.T) {
	test.Logf("Cmd: %s", t.Name)
	cmd := DefaultCmd
	if t.Cmd != "" {
		cmd = t.Cmd
	}
	test.Logf("Running %s with %s", cmd, t.Args)
	c := exec.Command(cmd, t.Args...)
	c.Env = t.Env

	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	if t.Input != "" {
		test.Logf("Input: %s", t.Input)
		c.Stdin = strings.NewReader(t.Input)
	}

	err := c.Run()
	result := err != nil
	test.Logf("stdout:\n%s", stdout.String())
	test.Logf("stderr:\n%s", stderr.String())

	// Do this twice for better tests logs
	if result != t.ExpectFailure {
		if t.ExpectFailure {
			test.Fatal("The command did not fail!")
		} else {
			test.Fatal("The command did fail!")
		}
	}

	outLinesCount := strings.Count(stdout.String(), "\n")
	if outLinesCount != t.OutLines {
		test.Errorf("stdout: wanted %d lines, got %d", t.OutLines, outLinesCount)
	}
	errLinesCount := strings.Count(stderr.String(), "\n")
	if errLinesCount != t.ErrLines {
		test.Errorf("stderr: wanted %d lines, got %d", t.ErrLines, errLinesCount)
	}

	if t.OutText != "" {
		if stdout.String() != t.OutText {
			test.Errorf("stdout does not match expectation:\n%s", t.OutText)
		}
	}

	if t.ErrText != "" {
		if stderr.String() != t.ErrText {
			test.Errorf("stderr does not match expectation:\n%s", t.ErrText)
		}
	}

	if t.OutRegex != "" {
		match, err := regexp.MatchString(t.OutRegex, stdout.String())
		if err != nil {
			test.Fatalf("Error while computing stdout regex: %v", err)
		}
		if !match {
			test.Errorf("No match for stdout regex: %s", t.OutRegex)
		}
	}

}
func init() {
	RunAsAdmin = []string{
		"TWELVE_TO_EIGHT_USER=admin",
		"TWELVE_TO_EIGHT_PASSWORD=pass",
	}
	RunAsUser = []string{
		"TWELVE_TO_EIGHT_USER=user",
		"TWELVE_TO_EIGHT_PASSWORD=pass",
	}
	DefaultCmd = "12to8"
	fixtures = []string{"basic_projects", "rich_timesheet"}
}

func newTimesheet(t *testing.T, c *dockerId) {
	userEnv := append(RunAsUser, c.EndpointEnv())
	currentTs := fmt.Sprintf("%s %d", time.Now().Month(), time.Now().Year())
	(&CmdTestCase{
		Name:     "Create timesheet",
		Env:      userEnv,
		Args:     []string{"new", "timesheet"},
		OutLines: 1,
		OutRegex: fmt.Sprintf("%s \\[ACTIVE\\]", currentTs),
	}).Run(t)
}
