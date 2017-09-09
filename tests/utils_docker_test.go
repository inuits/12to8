package tests

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
)

var skipDocker bool
var skipDockerImages bool
var skipDockerImagesDeletion bool

type dockerID struct {
	ID   string
	Port int
}

func createContainerWithFixture(fixture string) {
	if skipDocker || skipDockerImages {
		return
	}
	// Create a docker container
	c := exec.Command("docker", "run", "-d", "-e", "CFG_FILE_PATH=/code/925r.yml", "-e", fmt.Sprintf("FIXTURE=%s", fixture), "-v", fmt.Sprintf("%s:/tests", os.Getenv("PWD")), "925r:upstream", "/tests/run-925r.sh")
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	err := c.Run()
	if err != nil {
		log.Fatalf("stdout:\n%s", stdout.String())
		log.Fatalf("stderr:\n%s", stderr.String())
		log.Fatal(err)
	}
	id := strings.TrimSpace(stdout.String())

	// remove the container at the end
	defer exec.Command("docker", "rm", id).Run()

	// wait for the docker command to finish
	err = exec.Command("docker", "attach", id).Run()
	if err != nil {
		log.Fatalf("container %s can't attach: %v", fixture, err)
	}

	// commit the image with the data
	c = exec.Command("docker", "commit", id)
	var commitStdout, commitStderr bytes.Buffer
	c.Stdout = &commitStdout
	c.Stderr = &commitStderr
	err = c.Run()
	if err != nil {
		log.Fatalf("container %s can't commit: %v\n%s", fixture, err, commitStderr.String())
	}
	commitID := strings.TrimSpace(commitStdout.String())
	err = exec.Command("docker", "tag", commitID, fmt.Sprintf("925r:%s", fixture)).Run()
	if err != nil {
		log.Fatalf("container %s can't tag: %v", fixture, err)
	}
}

func deleteContainerWithFixture(fixture string) {
	if skipDocker || skipDockerImages || skipDockerImagesDeletion {
		return
	}
	err := exec.Command("docker", "rmi", fmt.Sprintf("925r:%s", fixture)).Run()
	if err != nil {
		log.Print(err)
	}
}

func (d *dockerID) start925r(t *testing.T, fixture string) {
	if skipDocker {
		t.Log("Not using docker.")
		d.Port = 8000
		return
	}
	t.Parallel()
	// Start a Docker container
	c := exec.Command("docker", "run", "-d", "-e", "FIXTURE=", "-p", "8000", "-v", fmt.Sprintf("%s:/tests", os.Getenv("PWD")), "--rm", fmt.Sprintf("925r:%s", fixture), "/tests/run-925r.sh")
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	err := c.Run()
	if err != nil {
		t.Logf("stdout:\n%s", stdout.String())
		t.Logf("stderr:\n%s", stderr.String())
		t.Fatal(err)
	}
	d.ID = strings.TrimSpace(stdout.String())

	time.Sleep(time.Second * 3)
	var stdoutInspect, stderrInspect bytes.Buffer
	c = exec.Command("docker", "inspect", d.ID, "-f", `{{index (index (index .NetworkSettings.Ports "8000/tcp") 0) "HostPort"}}`)
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

func (d *dockerID) waitFor925r(t *testing.T) {
	tries := 1
	maxTries := 10
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

func (d *dockerID) stop925r(t *testing.T) {
	if d == nil || d.ID == "" {
		return
	}
	c := exec.Command("docker", "kill", d.ID)
	c.Run()
}

func (d *dockerID) EndpointEnv() string {
	return fmt.Sprintf("TWELVE_TO_EIGHT_ENDPOINT=http://127.0.0.1:%d/api", d.Port)
}

func init() {
	flag.BoolVar(&skipDocker, "skip-docker", false, "Do not manage the 925r instances using docker. Make tests not parallel. Exects 925r on http://127.0.0.1:8000.")
	flag.BoolVar(&skipDockerImages, "skip-docker-images-creation", false, "Do not manage the 925r docker images with fixtures.")
	flag.BoolVar(&skipDockerImagesDeletion, "skip-docker-images-deletion", false, "Do not delete the 925r docker images with fixtures after the tests.")
}
