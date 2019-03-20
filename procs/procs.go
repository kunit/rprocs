package procs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Proc /proc/<pid>
type Proc struct {
	RootPath string
	Cmdline  []string
	Stat     Stat
}

// Stat see https://www.kernel.org/doc/Documentation/filesystems/proc.txt Table 1-4
type Stat struct {
	pid           int64
	tcomm         string
	state         string
	ppid          int64
	pgrp          int64
	sid           int64
	tty_nr        int64
	tty_pgrp      int64
	flags         uint64
	min_flt       uint64
	cmin_flt      uint64
	maj_flt       uint64
	cmaj_flt      uint64
	utime         uint64
	stime         uint64
	cutime        int64
	cstime        int64
	priority      int64
	nice          int64
	num_threads   int64
	it_real_value int64
	start_time    uint64
	vsize         uint64
	rss           int64
	rsslim        uint64
	start_code    uint64
	end_code      uint64
	start_stack   uint64
	esp           uint64
	eip           uint64
	pending       uint64
	blocked       uint64
	sigign        uint64
	sigcatch      uint64
	wchan         uint64
	nswap         uint64
	cnswap        uint64
	exit_signal   int64
	task_cpu      int64
	rt_priority   uint64
	policy        uint64
	blkio_ticks   uint64
	gtime         uint64
	cgtime        int64
	start_data    int64
	end_data      int64
	arg_start     int64
	arg_end       int64
	env_start     int64
	env_end       int64
	exit_code     int64
}

// ScanCmdline /proc/<pid>cmdline
func (p *Proc) ScanCmdline(pid uint) error {
	path := fmt.Sprintf("%s/%d/cmdline", p.RootPath, pid)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			p.Cmdline = strings.Split(t, string([]byte{0}))
		}
	}

	return nil
}

// ScanStat /proc/<pid>/stat
func (p *Proc) ScanStat(pid uint) error {
	path := fmt.Sprintf("%s/%d/stat", p.RootPath, pid)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			s := strings.Split(t, " ")
			fmt.Printf("t = %#v\n", t)
			fmt.Printf("s = %#v\n", s)

			p.Stat.pid, _ = strconv.ParseInt(s[0], 10, 64)
			p.Stat.tcomm = s[1]
			p.Stat.state = s[2]
			p.Stat.ppid, _ = strconv.ParseInt(s[3], 10, 64)
			p.Stat.pgrp, _ = strconv.ParseInt(s[4], 10, 64)
			p.Stat.sid, _ = strconv.ParseInt(s[5], 10, 64)
			p.Stat.tty_nr, _ = strconv.ParseInt(s[6], 10, 64)
			p.Stat.tty_pgrp, _ = strconv.ParseInt(s[7], 10, 64)
			p.Stat.flags, _ = strconv.ParseUint(s[8], 10, 64)
			p.Stat.min_flt, _ = strconv.ParseUint(s[9], 10, 64)
			p.Stat.cmin_flt, _ = strconv.ParseUint(s[10], 10, 64)
			p.Stat.maj_flt, _ = strconv.ParseUint(s[11], 10, 64)
			p.Stat.cmaj_flt, _ = strconv.ParseUint(s[12], 10, 64)
			p.Stat.utime, _ = strconv.ParseUint(s[13], 10, 64)
			p.Stat.stime, _ = strconv.ParseUint(s[14], 10, 64)
			p.Stat.cutime, _ = strconv.ParseInt(s[15], 10, 64)
			p.Stat.cstime, _ = strconv.ParseInt(s[16], 10, 64)
			p.Stat.priority, _ = strconv.ParseInt(s[17], 10, 64)
			p.Stat.nice, _ = strconv.ParseInt(s[18], 10, 64)
			p.Stat.num_threads, _ = strconv.ParseInt(s[19], 10, 64)
			p.Stat.it_real_value, _ = strconv.ParseInt(s[20], 10, 64)
			p.Stat.start_time, _ = strconv.ParseUint(s[21], 10, 64)
			p.Stat.vsize, _ = strconv.ParseUint(s[22], 10, 64)
			p.Stat.rss, _ = strconv.ParseInt(s[23], 10, 64)
			p.Stat.rsslim, _ = strconv.ParseUint(s[24], 10, 64)
			p.Stat.start_code, _ = strconv.ParseUint(s[25], 10, 64)
			p.Stat.end_code, _ = strconv.ParseUint(s[26], 10, 64)
			p.Stat.start_stack, _ = strconv.ParseUint(s[27], 10, 64)
			p.Stat.esp, _ = strconv.ParseUint(s[28], 10, 64)
			p.Stat.eip, _ = strconv.ParseUint(s[29], 10, 64)
			p.Stat.pending, _ = strconv.ParseUint(s[30], 10, 64)
			p.Stat.blocked, _ = strconv.ParseUint(s[31], 10, 64)
			p.Stat.sigign, _ = strconv.ParseUint(s[32], 10, 64)
			p.Stat.sigcatch, _ = strconv.ParseUint(s[33], 10, 64)
			p.Stat.wchan, _ = strconv.ParseUint(s[34], 10, 64)
			p.Stat.nswap, _ = strconv.ParseUint(s[35], 10, 64)
			p.Stat.cnswap, _ = strconv.ParseUint(s[36], 10, 64)
			p.Stat.exit_signal, _ = strconv.ParseInt(s[37], 10, 64)
			p.Stat.task_cpu, _ = strconv.ParseInt(s[38], 10, 64)
			p.Stat.rt_priority, _ = strconv.ParseUint(s[39], 10, 64)
			p.Stat.policy, _ = strconv.ParseUint(s[40], 10, 64)
			p.Stat.blkio_ticks, _ = strconv.ParseUint(s[41], 10, 64)
			p.Stat.gtime, _ = strconv.ParseUint(s[42], 10, 64)
			p.Stat.cgtime, _ = strconv.ParseInt(s[43], 10, 64)
			p.Stat.start_data, _ = strconv.ParseInt(s[44], 10, 64)
			p.Stat.end_data, _ = strconv.ParseInt(s[45], 10, 64)
			p.Stat.arg_start, _ = strconv.ParseInt(s[46], 10, 64)
			p.Stat.arg_end, _ = strconv.ParseInt(s[47], 10, 64)
			p.Stat.env_start, _ = strconv.ParseInt(s[48], 10, 64)
			p.Stat.env_end, _ = strconv.ParseInt(s[49], 10, 64)
			p.Stat.exit_code, _ = strconv.ParseInt(s[50], 10, 64)
		}
	}

	return nil
}

// Scan /proc/<pid>/xxx
func (p *Proc) Scan(pid uint) error {
	if err := p.ScanCmdline(pid); err != nil {
		return err
	}

	if err := p.ScanStat(pid); err != nil {
		return err
	}

	fmt.Printf("p = %#v\n", p)

	return nil
}
