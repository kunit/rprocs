package proc

import "testing"

func TestStatus_Scan(t *testing.T) {
	tests := []struct {
		name     string
		s        *Status
		rootPath string
		pid      int64
	}{
		{
			name:     "OK",
			s:        &Status{},
			rootPath: "/proc",
			pid:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.Scan(tt.rootPath, tt.pid)
			if err != nil {
				t.Errorf("Scan() error = %#v", err)
			}
			if tt.s.Pid != tt.pid {
				t.Errorf("Scan() error = %#v", tt.s)
			}
		})
	}
}
