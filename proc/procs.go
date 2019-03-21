package proc

import (
	"github.com/kunit/rprocs/system"
	"io"
	"os"
	"strconv"
)

// Procs process list
type Procs struct {
	Stat    system.Stat
	Uptime  system.Uptime
	Meminfo system.Meminfo
	ClkTck  uint64
	Procs   []Proc
}

// GetProc GET /v1/proc
func GetProc() (*Procs, error) {
	// TODO: from config
	rootPath := "/proc"

	p := &Procs{}

	if err := p.Stat.Scan(rootPath); err != nil {
		return nil, err
	}
	if err := p.Meminfo.Scan(rootPath); err != nil {
		return nil, err
	}
	if err := p.Uptime.Scan(rootPath); err != nil {
		return nil, err
	}

	clkTck, err := GetClkTck()
	if err != nil {
		return nil, err
	}
	p.ClkTck = clkTck

	pids, err := GetPids(rootPath)
	if err != nil {
		return nil, err
	}

	for _, pid := range pids {
		proc := &Proc{RootPath: rootPath}
		err := proc.Scan(&p.Stat, &p.Meminfo, &p.Uptime, p.ClkTck, pid)
		if err != nil {
			return nil, err
		}
		p.Procs = append(p.Procs, *proc)
	}

	return p, nil
}

// GetPids get pid list
func GetPids(rootPath string) ([]int64, error) {
	f, err := os.Open(rootPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	pids := make([]int64, 0, 50)
	for {
		files, err := f.Readdir(10)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if !file.IsDir() {
				continue
			}

			name := file.Name()
			if name[0] < '0' || name[0] > '9' {
				continue
			}

			pid, err := strconv.ParseInt(name, 10, 64)
			if err != nil {
				continue
			}

			pids = append(pids, pid)
		}
	}

	return pids, nil
}
