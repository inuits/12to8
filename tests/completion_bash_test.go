package acceptance_tests

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTopCompletion(t *testing.T) {
	testCompletion(t, nil, []string{"12to8", ""}, []string{"completion", "list", "new", "release"})
}

func completionBashCode(cli []string) string {
	flatcli := strings.Join(cli, " ")
	words := len(cli)
	return fmt.Sprintf(`
. /usr/share/bash-completion/bash_completion
. <(../12to8 completion bash)
COMP_WORDS=(%s)
COMP_CWORD=%d
COMP_LINE='%s'
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
		OutLines: 4,
		OutText:  expectedOut.String(),
	}
	if c != nil {
		tc.Env = append(RunAsUser, c.EndpointEnv())
	}
	tc.Run(t)
}
