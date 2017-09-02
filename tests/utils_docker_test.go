package acceptance_tests

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
)

var skip_docker bool

type dockerId struct {
	Id   string
	Port int
}

func (d *dockerId) start925r(t *testing.T, fixture string) {
	if skip_docker {
		t.Log("Not using docker.")
		d.Port = 8000
		return
	}
	t.Parallel()
	// Start a Docker container
	c := exec.Command("docker", "run", "-d", "-p", "8000", "-e", fmt.Sprintf("FIXTURE=%s", fixture), "-v", fmt.Sprintf("%s:/tests", os.Getenv("PWD")), "--rm", "roidelapluie/925r", "/tests/run-925r.sh")
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	err := c.Run()
	if err != nil {
		t.Logf("stdout:\n%s", stdout.String())
		t.Logf("stderr:\n%s", stderr.String())
		t.Fatal(err)
	}
	d.Id = strings.TrimSpace(stdout.String())

	time.Sleep(time.Second * 3)
	var stdoutInspect, stderrInspect bytes.Buffer
	c = exec.Command("docker", "inspect", d.Id, "-f", `{{index (index (index .NetworkSettings.Ports "8000/tcp") 0) "HostPort"}}`)
	c.Stdout = &stdoutInspect
	c.Stderr = &stderrInspect
	err = c.Run()
	if err != nil {
		t.Logf("Can't run docker inspect: %s\n%s", stdoutInspect.String(), stderrInspect.String())
		t.Fatal(err)
	}
	d.Port, err = strconv.Atoi(strings.TrimSpace(stdoutInspect.String()))
	if err != nil {
		t.Logf("Can't figure out port number: %s\n%s", stdoutInspect.String(), stderrInspect.String())
		t.Fatal(err)
	}
	d.waitFor925r(t)
}

func (d *dockerId) waitFor925r(t *testing.T) {
	tries := 1
	maxTries := 60
	addr := fmt.Sprintf("http://127.0.0.1:%d", d.Port)
	t.Logf("Waiting for 925r on %s", addr)
	for {
		time.Sleep(time.Second)
		_, err := http.Get(addr)
		if err == nil {
			t.Log("925r ready!")
			break
		}
		tries++
		if tries >= maxTries {
			t.Fatal("Too many attempts")
		}

	}
}

func (d *dockerId) stop925r(t *testing.T) {
	if d == nil || d.Id == "" {
		return
	}
	c := exec.Command("docker", "kill", d.Id)
	c.Run()
}

func (d *dockerId) EndpointEnv() string {
	return fmt.Sprintf("TWELVE_TO_EIGHT_ENDPOINT=http://127.0.0.1:%d/api", d.Port)
}

func init() {
	flag.BoolVar(&skip_docker, "skip-docker", false, "Do not manage the 925r instances using docker. Make tests not parallel. Exects 925r on http://127.0.0.1:8000.")
}
