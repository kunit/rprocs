package proc

import (
	"fmt"
	"github.com/kunit/rprocs/system"
	"github.com/tklauser/go-sysconf"
	"os/user"
	"strconv"
)

// Proc /system/<pid>
type Proc struct {
	RootPath string
	Uid      string
	Gid      string
	UserName string
	Name     string
	Cpu      string
	Memory   string
	Start    int64
	Time     int64
	Cmdline  Cmdline
	Stat     Stat
	Status   Status
}

// GetClkTck sysconf(_SC_CLK_TCK)
func GetClkTck() (uint64, error) {
	clktck, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	if err != nil {
		return 0, err
	}

	return uint64(clktck), nil
}

// Scan /system/<pid>/xxx
func (p *Proc) Scan(stat *system.Stat, meminfo *system.Meminfo, uptime *system.Uptime, clkTck uint64, pid int64) error {
	if err := p.Cmdline.Scan(p.RootPath, pid); err != nil {
		return err
	}

	if err := p.Stat.Scan(p.RootPath, pid); err != nil {
		return err
	}

	if err := p.Status.Scan(p.RootPath, pid); err != nil {
		return err
	}

	if err := p.SetUser(); err != nil {
		return err
	}

	p.SetCpu(uptime, clkTck)
	p.SetMemory(meminfo)
	p.SetStart(stat, clkTck)
	p.SetTime(clkTck)

	return nil
}

// SetUser Set User Property
func (p *Proc) SetUser() error {
	u, err := user.LookupId(fmt.Sprintf("%v", p.Status.Uid[0]))
	if err != nil {
		return err
	}

	p.Uid = u.Uid
	p.Gid = u.Gid
	p.UserName = u.Username
	p.Name = u.Name

	return nil
}

// SetCpu Calculation CPU Rate
func (p *Proc) SetCpu(uptime *system.Uptime, clkTck uint64) {
	totalTime := p.Stat.Utime + p.Stat.Stime + uint64(p.Stat.Cutime) + uint64(p.Stat.Cstime)
	var cpu float64
	t := p.CookEtime(uptime, clkTck)
	if t != 0 {
		cpu = float64(totalTime*1000/clkTck/t) / 10
	}
	p.Cpu = strconv.FormatFloat(cpu, 'f', 1, 64)
}

// SetMemory Calculation Memory Rate
func (p *Proc) SetMemory(meminfo *system.Meminfo) {
	memory := float64(p.Status.VmRSS*1000/meminfo.MemTotal) / 10
	p.Memory = strconv.FormatFloat(memory, 'f', 1, 64)
}

// CookEtime gitlab.com/procps-ng/procps/ps/output.c cook_etime
func (p *Proc) CookEtime(uptime *system.Uptime, clkTck uint64) uint64 {
	var s uint64

	SecondsSinceBoot := uptime.Up
	if uint64(SecondsSinceBoot) >= p.Stat.Starttime/clkTck {
		s = uint64(SecondsSinceBoot) - p.Stat.Starttime/clkTck
	}

	return s
}

// SetStart Calculate Process Start time
func (p *Proc) SetStart(stat *system.Stat, clkTck uint64) {
	p.Start = int64(float64(stat.Btime) + float64(p.Stat.Starttime/clkTck))
}

// SetTime Calculate Process time
func (p *Proc) SetTime(clkTck uint64) {
	totalTime := p.Stat.Utime + p.Stat.Stime
	p.Time = int64(totalTime / clkTck)
}
