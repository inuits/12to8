package acceptance_tests

import (
	"bytes"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"testing"
)

var RunAsAdmin []string
var RunAsUser []string
var DefaultCmd string

type CmdTestCase struct {
	Name          string
	ExpectFailure bool
	Cmd           string
	Args          []string
	OutLines      int
	ErrLines      int
	OutRegex      string
	ErrRegex      string
	Env           []string
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

	err := c.Run()
	result := err != nil
	test.Logf("stdout:\n%s", stdout.String())
	test.Logf("stderr:\n%s", stderr.String())
	test.Logf("Errorful: %v", err == nil)
	if result != t.ExpectFailure {
		test.Fatalf("Expected failure: %s, but got: %s", t.ExpectFailure, result)
	}

	outLinesCount := strings.Count(stdout.String(), "\n")
	if outLinesCount != t.OutLines {
		test.Errorf("stdout: wanted %d lines, got %d", t.OutLines, outLinesCount)
	}
	errLinesCount := strings.Count(stderr.String(), "\n")
	if errLinesCount != t.ErrLines {
		test.Errorf("stderr: wanted %d lines, got %d", t.ErrLines, errLinesCount)
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
	DefaultCmd = path.Join(os.Getenv("PWD"), "..", "12to8")
}
