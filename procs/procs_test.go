package procs

import "testing"

func TestProc_Scan(t *testing.T) {
	p := &Proc{RootPath: "/proc"}
	p.Scan(1)
}
