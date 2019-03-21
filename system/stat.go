package system

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stat /system/stat
type Stat struct {
	Cpu          []uint64            `json:"cpu"`
	Cpus         map[string][]uint64 `json:"cpus"`
	Intr         []uint64            `json:"intr"`
	Ctxt         uint64              `json:"ctxt"`
	Btime        uint64              `json:"btime"`
	Processes    uint64              `json:"processes"`
	ProcsRunning uint64              `json:"procs_running"`
	ProcsBlocked uint64              `json:"procs_blocked"`
	Softirq      []uint64            `json:"softirq"`
}

// Scan /system/stat
func (s *Stat) Scan(rootPath string) error {
	path := fmt.Sprintf("%s/stat", rootPath)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	s.Cpus = make(map[string][]uint64)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			sp := strings.Split(t, " ")

			switch {
			case sp[0] == "cpu":
				s.Cpu = ParseUint(sp[1:])
			case sp[0][:3] == "cpu":
				s.Cpus[sp[0]] = ParseUint(sp[1:])
			case sp[0] == "intr":
				s.Intr = ParseUint(sp[1:])
			case sp[0] == "ctxt":
				s.Ctxt, _ = strconv.ParseUint(sp[1], 10, 64)
			case sp[0] == "btime":
				s.Btime, _ = strconv.ParseUint(sp[1], 10, 64)
			case sp[0] == "processes":
				s.Processes, _ = strconv.ParseUint(sp[1], 10, 64)
			case sp[0] == "procs_running":
				s.ProcsRunning, _ = strconv.ParseUint(sp[1], 10, 64)
			case sp[0] == "proc_blocked":
				s.ProcsBlocked, _ = strconv.ParseUint(sp[1], 10, 64)
			case sp[0] == "softirq":
				s.Softirq = ParseUint(sp[1:])
			}
		}
	}

	return nil
}

// ParseUint Split use sep param And ParseInt
func ParseUint(sp []string) []uint64 {
	var r []uint64

	for _, v := range sp {
		vv, _ := strconv.ParseUint(v, 10, 64)
		r = append(r, vv)
	}

	return r
}
