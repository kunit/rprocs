package system

import (
	"testing"
)

func TestMeminfo_Scan(t *testing.T) {
	tests := []struct {
		name     string
		m        *Meminfo
		rootPath string
	}{
		{
			name:     "OK",
			m:        &Meminfo{},
			rootPath: "/proc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.m.Scan(tt.rootPath)
			if err != nil {
				t.Errorf("Scan() error = %#v", err)
			}
			if tt.m.MemTotal <= 0 {
				t.Errorf("Scan() error = %#v", tt.m)
			}
		})
	}
}
