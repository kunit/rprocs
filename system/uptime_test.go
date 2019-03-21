package system

import (
	"testing"
)

func TestUptime_Scan(t *testing.T) {
	tests := []struct {
		name     string
		u        *Uptime
		rootPath string
	}{
		{
			name:     "OK",
			u:        &Uptime{},
			rootPath: "/proc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.u.Scan(tt.rootPath)
			if err != nil {
				t.Errorf("Scan() error = %#v", err)
			}
			if tt.u.Up <= 0 {
				t.Errorf("Scan() error = %#v", tt.u)
			}
			if tt.u.Idle <= 0 {
				t.Errorf("Scan() error = %#v", tt.u)
			}
		})
	}
}
