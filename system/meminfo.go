package system

import (
	"bufio"
	"fmt"
	"github.com/kunit/rprocs/utils"
	"os"
	"strings"
)

// Meminfo /system/meminfo
type Meminfo struct {
	MemTotal          uint64 `json:"MemTotal"`
	MemFree           uint64 `json:"MemFree"`
	MemAvailable      uint64 `json:"MemAvailable"`
	Buffers           uint64 `json:"Buffers"`
	Cached            uint64 `json:"Cached"`
	SwapCached        uint64 `json:"SwapCached"`
	Active            uint64 `json:"Active"`
	Inactive          uint64 `json:"Inactive"`
	ActiveAnon        uint64 `json:"Active(anon)"`
	InactiveAnon      uint64 `json:"Inactive(anon)"`
	ActiveFile        uint64 `json:"Active(file)"`
	InactiveFile      uint64 `json:"Inactive(file)"`
	Unevictable       uint64 `json:"Unevictable"`
	Mlocked           uint64 `json:"Mlocked"`
	SwapTotal         uint64 `json:"SwapTotal"`
	SwapFree          uint64 `json:"SwapFree"`
	Dirty             uint64 `json:"Dirty"`
	Writeback         uint64 `json:"Writeback"`
	AnonPages         uint64 `json:"AnonPages"`
	Mapped            uint64 `json:"Mapped"`
	Shmem             uint64 `json:"Shmem"`
	Slab              uint64 `json:"Slab"`
	SReclaimable      uint64 `json:"SReclaimable"`
	SUnreclaim        uint64 `json:"SUnreclaim"`
	KernelStack       uint64 `json:"KernelStack"`
	PageTables        uint64 `json:"PageTables"`
	NFSUnstable       uint64 `json:"NFS_Unstable"`
	Bounce            uint64 `json:"Bounce"`
	WritebackTmp      uint64 `json:"WritebackTmp"`
	CommitLimit       uint64 `json:"CommitLimit"`
	CommittedAS       uint64 `json:"Committed_AS"`
	VmallocTotal      uint64 `json:"VmallocTotal"`
	VmallocUsed       uint64 `json:"VmallocUsed"`
	VmallocChunk      uint64 `json:"VmallocChunk"`
	HardwareCorrupted uint64 `json:"HardwareCorrupted"`
	AnonHugePages     uint64 `json:"AnonHugePages"`
	ShmemHugePages    uint64 `json:"ShmemHugePages"`
	ShmemPmdMapped    uint64 `json:"ShmemPmdMapped"`
	CmaTotal          uint64 `json:"CmaTotal"`
	CmaFree           uint64 `json:"CmaFree"`
	HugePagesTotal    string `json:"HugePages_Total"`
	HugePagesFree     string `json:"HugePages_Free"`
	HugePagesRsvd     string `json:"HugePages_Rsvd"`
	HugePagesSurp     string `json:"HugePages_Surp"`
	Hugepagesize      uint64 `json:"Hugepagesize"`
	DirectMap4k       uint64 `json:"DirectMap4k"`
	DirectMap2M       uint64 `json:"DirectMap2M"`
	DirectMap1G       uint64 `json:"DirectMap1G"`
}

// Scan /system/meminfo
func (m *Meminfo) Scan(rootPath string) error {
	path := fmt.Sprintf("%s/meminfo", rootPath)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if t := scanner.Text(); t != "" {
			sp := strings.Split(t, ":")
			sp[1] = strings.TrimSpace(sp[1])
			sp[1] = strings.Trim(sp[1], "\t")

			switch sp[0] {
			case "MemTotal":
				m.MemTotal = utils.SplitAndParseUint64(sp[1], " ")
			case "MemFree":
				m.MemFree = utils.SplitAndParseUint64(sp[1], " ")
			case "MemAvailable":
				m.MemAvailable = utils.SplitAndParseUint64(sp[1], " ")
			case "Buffers":
				m.Buffers = utils.SplitAndParseUint64(sp[1], " ")
			case "Cached":
				m.Cached = utils.SplitAndParseUint64(sp[1], " ")
			case "SwapCached":
				m.SwapCached = utils.SplitAndParseUint64(sp[1], " ")
			case "Active":
				m.Active = utils.SplitAndParseUint64(sp[1], " ")
			case "Inactive":
				m.Inactive = utils.SplitAndParseUint64(sp[1], " ")
			case "Active(anon)":
				m.ActiveAnon = utils.SplitAndParseUint64(sp[1], " ")
			case "Inactive(anon)":
				m.InactiveAnon = utils.SplitAndParseUint64(sp[1], " ")
			case "Active(file)":
				m.ActiveFile = utils.SplitAndParseUint64(sp[1], " ")
			case "Inactive(file)":
				m.InactiveFile = utils.SplitAndParseUint64(sp[1], " ")
			case "Unevictable":
				m.Unevictable = utils.SplitAndParseUint64(sp[1], " ")
			case "Mlocked":
				m.Mlocked = utils.SplitAndParseUint64(sp[1], " ")
			case "SwapTotal":
				m.SwapTotal = utils.SplitAndParseUint64(sp[1], " ")
			case "SwapFree":
				m.SwapFree = utils.SplitAndParseUint64(sp[1], " ")
			case "Dirty":
				m.Dirty = utils.SplitAndParseUint64(sp[1], " ")
			case "Writeback":
				m.Writeback = utils.SplitAndParseUint64(sp[1], " ")
			case "AnonPages":
				m.AnonPages = utils.SplitAndParseUint64(sp[1], " ")
			case "Mapped":
				m.Mapped = utils.SplitAndParseUint64(sp[1], " ")
			case "Shmem":
				m.Shmem = utils.SplitAndParseUint64(sp[1], " ")
			case "Slab":
				m.Slab = utils.SplitAndParseUint64(sp[1], " ")
			case "SReclaimable":
				m.SReclaimable = utils.SplitAndParseUint64(sp[1], " ")
			case "SUnreclaim":
				m.SUnreclaim = utils.SplitAndParseUint64(sp[1], " ")
			case "KernelStack":
				m.KernelStack = utils.SplitAndParseUint64(sp[1], " ")
			case "PageTables":
				m.PageTables = utils.SplitAndParseUint64(sp[1], " ")
			case "NFS_Unstable":
				m.NFSUnstable = utils.SplitAndParseUint64(sp[1], " ")
			case "Bounce":
				m.Bounce = utils.SplitAndParseUint64(sp[1], " ")
			case "WritebackTmp":
				m.WritebackTmp = utils.SplitAndParseUint64(sp[1], " ")
			case "CommitLimit":
				m.CommitLimit = utils.SplitAndParseUint64(sp[1], " ")
			case "Committed_AS":
				m.CommittedAS = utils.SplitAndParseUint64(sp[1], " ")
			case "VmallocTotal":
				m.VmallocTotal = utils.SplitAndParseUint64(sp[1], " ")
			case "VmallocUsed":
				m.VmallocUsed = utils.SplitAndParseUint64(sp[1], " ")
			case "VmallocChunk":
				m.VmallocChunk = utils.SplitAndParseUint64(sp[1], " ")
			case "HardwareCorrupted":
				m.HardwareCorrupted = utils.SplitAndParseUint64(sp[1], " ")
			case "AnonHugePages":
				m.AnonHugePages = utils.SplitAndParseUint64(sp[1], " ")
			case "ShmemHugePages":
				m.ShmemHugePages = utils.SplitAndParseUint64(sp[1], " ")
			case "ShmemPmdMapped":
				m.ShmemPmdMapped = utils.SplitAndParseUint64(sp[1], " ")
			case "CmaTotal":
				m.CmaTotal = utils.SplitAndParseUint64(sp[1], " ")
			case "CmaFree":
				m.CmaFree = utils.SplitAndParseUint64(sp[1], " ")
			case "HugePages_Total":
				m.HugePagesTotal = sp[1]
			case "HugePages_Free":
				m.HugePagesFree = sp[1]
			case "HugePages_Rsvd":
				m.HugePagesRsvd = sp[1]
			case "HugePages_Surp":
				m.HugePagesSurp = sp[1]
			case "Hugepagesize":
				m.Hugepagesize = utils.SplitAndParseUint64(sp[1], " ")
			case "DirectMap4k":
				m.DirectMap4k = utils.SplitAndParseUint64(sp[1], " ")
			case "DirectMap2M":
				m.DirectMap2M = utils.SplitAndParseUint64(sp[1], " ")
			case "DirectMap1G":
				m.DirectMap1G = utils.SplitAndParseUint64(sp[1], " ")
			}
		}
	}

	return nil
}
