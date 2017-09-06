package acceptance_tests

import (
	"testing"
)

//
func TestDeletePerformancePorcelain(t *testing.T) {
	c := &dockerId{}
	c.start925r(t, "rich_timesheet")
	defer c.stop925r(t)
	userEnv := append(RunAsUser, c.EndpointEnv())
	(&CmdTestCase{
		Name:     "List performances",
		Env:      userEnv,
		Args:     []string{"list", "performances", "09/2017", "-P"},
		OutLines: 20,
		OutText: `1,04/09/2017,,Consult,8.00,1.00,activity
2,06/09/2017,,Consult,8.00,1.00,activity
3,07/09/2017,,Consult,8.00,1.00,activity
4,10/09/2017,,Webby,8.00,1.00,activity
5,11/09/2017,,Webby,8.00,1.00,activity
6,12/09/2017,,Webby,8.00,1.00,activity
7,12/09/2017,,keep it up!!,8.00,1.00,activity
8,14/09/2017,,Webby,8.00,1.00,activity
18,17/09/2017,,Consult,8.00,1.00,activity
19,18/09/2017,,Consult,8.00,1.00,activity
20,19/09/2017,,Consult,8.00,1.00,activity
17,20/09/2017,,Consult,8.00,1.00,activity
16,21/09/2017,,Consult,3.00,1.00,activity
12,25/09/2017,,Consult,4.00,1.00,activity
15,25/09/2017,,Consult,12.00,1.00,activity
11,25/09/2017,,Webby,4.00,1.00,activity
14,26/09/2017,,Consult,4.00,1.00,activity
21,27/09/2017,YES,Consult,4.00,1.00,activity
10,27/09/2017,,Webby,4.00,1.00,activity
9,28/09/2017,,Webby,8.00,1.00,activity
`,
	}).Run(t)
}
