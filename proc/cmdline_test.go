package proc

import (
	"testing"
)

func TestCmdline_Scan(t *testing.T) {
	tests := []struct {
		name     string
		c        *Cmdline
		rootPath string
		pid      int64
	}{
		{
			name:     "OK",
			c:        &Cmdline{},
			rootPath: "/proc",
			pid:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.c.Scan(tt.rootPath, tt.pid)
			if err != nil {
				t.Errorf("Scan() error = %#v", err)
			}
			if len(tt.c.Args) == 0 {
				t.Errorf("Scan() error = %#v", tt.c)
			}
		})
	}
}
