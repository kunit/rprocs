package proc

import (
	"bufio"
	"fmt"
	"github.com/kunit/rprocs/utils"
	"os"
	"strconv"
	"strings"
)

// Status /system/<pid>/status
type Status struct {
	Name                     string   `json:"Name"`
	Umask                    string   `json:"Umask"`
	State                    string   `json:"State"`
	Tgid                     int64    `json:"Tgid"`
	Ngid                     int64    `json:"Ngid"`
	Pid                      int64    `json:"Pid"`
	PPid                     int64    `json:"PPid"`
	TracerPid                int64    `json:"TracerPid"`
	Uid                      []int64  `json:"Uid"`
	Gid                      []int64  `json:"Gid"`
	FDSize                   int64    `json:"FDSize"`
	Groups                   []int64  `json:"Groups"`
	NStgid                   []int64  `json:"NStgid"`
	NSpid                    []int64  `json:"NSpid"`
	NSpgid                   []int64  `json:"NSpgid"`
	NSsid                    []int64  `json:"NSsid"`
	VmPeak                   uint64   `json:"VmPeak"`
	VmSize                   uint64   `json:"VmSize"`
	VmLck                    uint64   `json:"VmLck"`
	VmPin                    uint64   `json:"VmPin"`
	VmHWM                    uint64   `json:"VmHWM"`
	VmRSS                    uint64   `json:"VmRSS"`
	RssAnon                  uint64   `json:"RssAnon"`
	RssFile                  uint64   `json:"RssFile"`
	RssShmem                 uint64   `json:"RssShmem"`
	VmData                   uint64   `json:"VmData"`
	VmStk                    uint64   `json:"VmStk"`
	VmExe                    uint64   `json:"VmExe"`
	VmLib                    uint64   `json:"VmLib"`
	VmPTE                    uint64   `json:"VmPTE"`
	VmSwap                   uint64   `json:"VmSwap"`
	HugetlbPages             uint64   `json:"HugetlbPages"`
	CoreDumping              uint64   `json:"CoreDumping"`
	Threads                  uint64   `json:"Threads"`
	SigQ                     []int64  `json:"SigQ"`
	SigPnd                   string   `json:"SigPnd"`
	ShdPnd                   string   `json:"ShdPnd"`
	SigBlk                   string   `json:"SigBlk"`
	SigIgn                   string   `json:"SigIgn"`
	SigCgt                   string   `json:"SigCgt"`
	CapInh                   string   `json:"CapInh"`
	CapPrm                   string   `json:"CapPrm"`
	CapEff                   string   `json:"CapEff"`
	CapBnd                   string   `json:"CapBnd"`
	CapAmb                   string   `json:"CapAmb"`
	NoNewPrivs               uint64   `json:"NoNewPrivs"`
	Seccomp                  uint64   `json:"Seccomp"`
	SpeculationStoreBypass   string   `json:"Speculation_Store_Bypass"`
	CpusAllowed              string   `json:"Cpus_allowed"`
	CpusAllowedList          string   `json:"Cpus_allowed_list"`
	MemsAllowed              []string `json:"Mems_allowed"`
	MemsAllowedList          string   `json:"Mems_allowed_list"`
	VoluntaryCtxtSwitches    uint64   `json:"voluntary_ctxt_switches"`
	NonvoluntaryCtxtSwitches uint64   `json:"nonvoluntary_ctxt_switches"`
}

// Scan /system/<pid>/status
func (s *Status) Scan(rootPath string, pid int64) error {
	path := fmt.Sprintf("%s/%d/status", rootPath, pid)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			sp := strings.Split(t, ":")
			if len(sp) > 2 {
				sp[1] = strings.Join(sp[1:], ":")
			}
			sp[1] = strings.TrimSpace(sp[1])
			sp[1] = strings.Trim(sp[1], "\t")

			switch sp[0] {
			case "Name":
				s.Name = sp[1]
			case "Umask":
				s.Umask = sp[1]
			case "State":
				s.State = sp[1]
			case "Tgid":
				s.Tgid, _ = strconv.ParseInt(sp[1], 10, 64)
			case "Ngid":
				s.Ngid, _ = strconv.ParseInt(sp[1], 10, 64)
			case "Pid":
				s.Pid, _ = strconv.ParseInt(sp[1], 10, 64)
			case "PPid":
				s.PPid, _ = strconv.ParseInt(sp[1], 10, 64)
			case "TracerPid":
				s.TracerPid, _ = strconv.ParseInt(sp[1], 10, 64)
			case "Uid":
				s.Uid = utils.SplitAndRangeParseInt64(sp[1], "\t")
			case "Gid":
				s.Gid = utils.SplitAndRangeParseInt64(sp[1], "\t")
			case "FDSize":
				s.FDSize, _ = strconv.ParseInt(sp[1], 10, 64)
			case "Groups":
				s.Groups = utils.SplitAndRangeParseInt64(sp[1], "\t")
			case "NStgid":
				s.NStgid = utils.SplitAndRangeParseInt64(sp[1], "\t")
			case "NSpid":
				s.NSpid = utils.SplitAndRangeParseInt64(sp[1], "\t")
			case "NSpgid":
				s.NSpgid = utils.SplitAndRangeParseInt64(sp[1], "\t")
			case "NSsid":
				s.NSsid = utils.SplitAndRangeParseInt64(sp[1], "\t")
			case "VmPeak":
				s.VmPeak = utils.SplitAndParseUint64(sp[1], " ")
			case "VmSize":
				s.VmSize = utils.SplitAndParseUint64(sp[1], " ")
			case "VmLck":
				s.VmLck = utils.SplitAndParseUint64(sp[1], " ")
			case "VmPin":
				s.VmPin = utils.SplitAndParseUint64(sp[1], " ")
			case "VmHWM":
				s.VmHWM = utils.SplitAndParseUint64(sp[1], " ")
			case "VmRSS":
				s.VmRSS = utils.SplitAndParseUint64(sp[1], " ")
			case "RssAnon":
				s.RssAnon = utils.SplitAndParseUint64(sp[1], " ")
			case "RssFile":
				s.RssFile = utils.SplitAndParseUint64(sp[1], " ")
			case "RssShmem":
				s.RssShmem = utils.SplitAndParseUint64(sp[1], " ")
			case "VmData":
				s.VmData = utils.SplitAndParseUint64(sp[1], " ")
			case "VmStk":
				s.VmStk = utils.SplitAndParseUint64(sp[1], " ")
			case "VmExe":
				s.VmExe = utils.SplitAndParseUint64(sp[1], " ")
			case "VmLib":
				s.VmLib = utils.SplitAndParseUint64(sp[1], " ")
			case "VmPTE":
				s.VmPTE = utils.SplitAndParseUint64(sp[1], " ")
			case "VmSwap":
				s.VmSwap = utils.SplitAndParseUint64(sp[1], " ")
			case "HugetlbPages":
				s.HugetlbPages = utils.SplitAndParseUint64(sp[1], " ")
			case "CoreDumping":
				s.CoreDumping, _ = strconv.ParseUint(sp[1], 10, 64)
			case "Threads":
				s.Threads, _ = strconv.ParseUint(sp[1], 10, 64)
			case "SigQ":
				s.SigQ = utils.SplitAndRangeParseInt64(sp[1], "/")
			case "SigPnd":
				s.SigPnd = sp[1]
			case "ShdPnd":
				s.ShdPnd = sp[1]
			case "SigBlk":
				s.SigBlk = sp[1]
			case "SigIgn":
				s.SigIgn = sp[1]
			case "SigCgt":
				s.SigCgt = sp[1]
			case "CapInh":
				s.CapInh = sp[1]
			case "CapPrm":
				s.CapPrm = sp[1]
			case "CapEff":
				s.CapEff = sp[1]
			case "CapBnd":
				s.CapBnd = sp[1]
			case "CapAmb":
				s.CapAmb = sp[1]
			case "NoNewPrivs":
				s.NoNewPrivs, _ = strconv.ParseUint(sp[1], 10, 64)
			case "Seccomp":
				s.Seccomp, _ = strconv.ParseUint(sp[1], 10, 64)
			case "Speculation_Store_Bypass":
				s.SpeculationStoreBypass = sp[1]
			case "Cpus_allowed":
				s.CpusAllowed = sp[1]
			case "Cpus_allowed_list":
				s.CpusAllowedList = sp[1]
			case "Mems_allowed":
				s.MemsAllowed = strings.Split(sp[1], ",")
			case "Mems_allowed_list":
				s.MemsAllowedList = sp[1]
			case "voluntary_ctxt_switches":
				s.VoluntaryCtxtSwitches, _ = strconv.ParseUint(sp[1], 10, 64)
			case "nonvoluntary_ctxt_switches":
				s.NonvoluntaryCtxtSwitches, _ = strconv.ParseUint(sp[1], 10, 64)
			}
		}
	}

	return nil
}
