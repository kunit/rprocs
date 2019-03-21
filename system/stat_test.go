package system

import (
	"testing"
)

func TestStat_Scan(t *testing.T) {
	tests := []struct {
		name     string
		s        *Stat
		rootPath string
	}{
		{
			name:     "OK",
			s:        &Stat{},
			rootPath: "/proc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.Scan(tt.rootPath)
			if err != nil {
				t.Errorf("Scan() error = %#v", err)
			}
			if len(tt.s.Cpu) == 0 {
				t.Errorf("Scan() error = %#v", tt.s)
			}
		})
	}
}
