package proc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Cmdline /system/<pid>/cmdline
type Cmdline struct {
	Args []string `json:"args"`
}

// Scan /system/<pid>/cmdline
func (c *Cmdline) Scan(rootPath string, pid int64) error {
	path := fmt.Sprintf("%s/%d/cmdline", rootPath, pid)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			sp := strings.Split(t, string([]byte{0}))
			c.Args = sp[0 : len(sp)-1]
		}
	}

	return nil
}
