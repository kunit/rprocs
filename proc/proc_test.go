package proc

import (
	"github.com/kunit/rprocs/system"
	"testing"
)

func TestProc_Scan(t *testing.T) {
	tests := []struct {
		name     string
		p        *Proc
		rootPath string
		pid      int64
	}{
		{
			name:     "OK",
			p:        &Proc{},
			rootPath: "/proc",
			pid:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.RootPath = tt.rootPath
			stat := &system.Stat{}
			_ = stat.Scan(tt.rootPath)
			meminfo := &system.Meminfo{}
			_ = meminfo.Scan(tt.rootPath)
			uptime := &system.Uptime{}
			_ = uptime.Scan(tt.rootPath)
			clkTck, _ := GetClkTck()
			err := tt.p.Scan(stat, meminfo, uptime, clkTck, tt.pid)
			if err != nil {
				t.Errorf("Scan() error = %#v", err)
			}
			if len(tt.p.Cmdline.Args) == 0 {
				t.Errorf("Scan() error = %#v", tt.p.Cmdline)
			}
			if tt.p.Stat.Pid != tt.pid {
				t.Errorf("Scan() error = %#v", tt.p.Stat)
			}
			if tt.p.Status.Pid != tt.pid {
				t.Errorf("Scan() error = %#v", tt.p.Stat)
			}
		})
	}
}
