package proc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stat see https://linuxjm.osdn.jp/html/LDP_man-pages/man5/rproc.5.html
type Stat struct {
	Pid                 int64  `json:"pid"`
	Comm                string `json:"comm"`
	State               string `json:"state"`
	Ppid                int64  `json:"ppid"`
	Pgrp                int64  `json:"pgrp"`
	Session             int64  `json:"session"`
	TtyNr               int64  `json:"tty_nr"`
	Tpgid               int64  `json:"tpgid"`
	Flags               uint64 `json:"flags"`
	Minflt              uint64 `json:"minflt"`
	Cminflt             uint64 `json:"cminflt"`
	Majflt              uint64 `json:"majflt"`
	Cmajflt             uint64 `json:"cmajflt"`
	Utime               uint64 `json:"utime"`
	Stime               uint64 `json:"stime"`
	Cutime              int64  `json:"cutime"`
	Cstime              int64  `json:"cstime"`
	Priority            int64  `json:"priority"`
	Nice                int64  `json:"nice"`
	NumThreads          int64  `json:"num_threads"`
	Itrealvalue         int64  `json:"itrealvalue"`
	Starttime           uint64 `json:"starttime"`
	Vsize               uint64 `json:"vsize"`
	Rss                 int64  `json:"rss"`
	Rsslim              uint64 `json:"rsslim"`
	Startcode           uint64 `json:"startcode"`
	Endcode             uint64 `json:"endcode"`
	Startstack          uint64 `json:"startstack"`
	Kstkesp             uint64 `json:"kstkesp"`
	Kstkeip             uint64 `json:"kstkeip"`
	Signal              uint64 `json:"signal"`
	Blocked             uint64 `json:"blocked"`
	Sigignore           uint64 `json:"sigignore"`
	Sigcatch            uint64 `json:"sigcatch"`
	Wchan               uint64 `json:"wchan"`
	Nswap               uint64 `json:"nswap"`
	Cnswap              uint64 `json:"cnswap"`
	ExitSignal          int64  `json:"exit_signal"`
	Processor           int64  `json:"processor"`
	RtPriority          uint64 `json:"rt_priority"`
	Policy              uint64 `json:"policy"`
	DelayacctBlkioTicks uint64 `json:"delayacct_blkio_ticks"`
	GuestTime           uint64 `json:"guest_time"`
	CguestTime          int64  `json:"cguest_time"`
	StartData           uint64 `json:"start_data"`
	EndData             uint64 `json:"end_data"`
	StartBrk            uint64 `json:"start_brk"`
	ArgStart            uint64 `json:"arg_start"`
	ArgEnd              uint64 `json:"arg_end"`
	EnvStart            uint64 `json:"env_start"`
	EnvEnd              uint64 `json:"env_end"`
	ExitCode            int64  `json:"exit_code"`
}

// Scan /system/<pid>/stat
func (s *Stat) Scan(rootPath string, pid int64) error {
	path := fmt.Sprintf("%s/%d/stat", rootPath, pid)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			sp := strings.Split(t, " ")

			s.Pid, _ = strconv.ParseInt(sp[0], 10, 64)
			s.Comm = sp[1]
			s.State = sp[2]
			s.Ppid, _ = strconv.ParseInt(sp[3], 10, 64)
			s.Pgrp, _ = strconv.ParseInt(sp[4], 10, 64)
			s.Session, _ = strconv.ParseInt(sp[5], 10, 64)
			s.TtyNr, _ = strconv.ParseInt(sp[6], 10, 64)
			s.Tpgid, _ = strconv.ParseInt(sp[7], 10, 64)
			s.Flags, _ = strconv.ParseUint(sp[8], 10, 64)
			s.Minflt, _ = strconv.ParseUint(sp[9], 10, 64)
			s.Cminflt, _ = strconv.ParseUint(sp[10], 10, 64)
			s.Majflt, _ = strconv.ParseUint(sp[11], 10, 64)
			s.Cmajflt, _ = strconv.ParseUint(sp[12], 10, 64)
			s.Utime, _ = strconv.ParseUint(sp[13], 10, 64)
			s.Stime, _ = strconv.ParseUint(sp[14], 10, 64)
			s.Cutime, _ = strconv.ParseInt(sp[15], 10, 64)
			s.Cstime, _ = strconv.ParseInt(sp[16], 10, 64)
			s.Priority, _ = strconv.ParseInt(sp[17], 10, 64)
			s.Nice, _ = strconv.ParseInt(sp[18], 10, 64)
			s.NumThreads, _ = strconv.ParseInt(sp[19], 10, 64)
			s.Itrealvalue, _ = strconv.ParseInt(sp[20], 10, 64)
			s.Starttime, _ = strconv.ParseUint(sp[21], 10, 64)
			s.Vsize, _ = strconv.ParseUint(sp[22], 10, 64)
			s.Rss, _ = strconv.ParseInt(sp[23], 10, 64)
			s.Rsslim, _ = strconv.ParseUint(sp[24], 10, 64)
			s.Startcode, _ = strconv.ParseUint(sp[25], 10, 64)
			s.Endcode, _ = strconv.ParseUint(sp[26], 10, 64)
			s.Startstack, _ = strconv.ParseUint(sp[27], 10, 64)
			s.Kstkesp, _ = strconv.ParseUint(sp[28], 10, 64)
			s.Kstkeip, _ = strconv.ParseUint(sp[29], 10, 64)
			s.Signal, _ = strconv.ParseUint(sp[30], 10, 64)
			s.Blocked, _ = strconv.ParseUint(sp[31], 10, 64)
			s.Sigignore, _ = strconv.ParseUint(sp[32], 10, 64)
			s.Sigcatch, _ = strconv.ParseUint(sp[33], 10, 64)
			s.Wchan, _ = strconv.ParseUint(sp[34], 10, 64)
			s.Nswap, _ = strconv.ParseUint(sp[35], 10, 64)
			s.Cnswap, _ = strconv.ParseUint(sp[36], 10, 64)
			s.ExitSignal, _ = strconv.ParseInt(sp[37], 10, 64)
			s.Processor, _ = strconv.ParseInt(sp[38], 10, 64)
			s.RtPriority, _ = strconv.ParseUint(sp[39], 10, 64)
			s.Policy, _ = strconv.ParseUint(sp[40], 10, 64)
			s.DelayacctBlkioTicks, _ = strconv.ParseUint(sp[41], 10, 64)
			s.GuestTime, _ = strconv.ParseUint(sp[42], 10, 64)
			s.CguestTime, _ = strconv.ParseInt(sp[43], 10, 64)
			s.StartData, _ = strconv.ParseUint(sp[44], 10, 64)
			s.EndData, _ = strconv.ParseUint(sp[45], 10, 64)
			s.StartBrk, _ = strconv.ParseUint(sp[46], 10, 64)
			s.ArgStart, _ = strconv.ParseUint(sp[47], 10, 64)
			s.ArgEnd, _ = strconv.ParseUint(sp[48], 10, 64)
			s.EnvStart, _ = strconv.ParseUint(sp[49], 10, 64)
			s.EnvEnd, _ = strconv.ParseUint(sp[50], 10, 64)
			s.ExitCode, _ = strconv.ParseInt(sp[51], 10, 64)
		}
	}

	return nil
}
