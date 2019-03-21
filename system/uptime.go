package system

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Uptime /system/uptime
type Uptime struct {
	Up   float64
	Idle float64
}

// Scan /system/uptime
func (u *Uptime) Scan(rootPath string) error {
	path := fmt.Sprintf("%s/uptime", rootPath)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			sp := strings.Split(t, " ")
			u.Up, _ = strconv.ParseFloat(sp[0], 64)
			u.Idle, _ = strconv.ParseFloat(sp[1], 64)
		}
	}

	return nil
}
