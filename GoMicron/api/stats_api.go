package api

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"os"
	"log"
	"github.com/AtahanPoyraz/cmd"

)

type Stats struct {
	CPU_Percent  float64  `json:"CPU_Percent"`
	RAM_Percent  float64  `json:"RAM_Percent"`
	Disk_Percent float64  `json:"Disk_Percent"`
}

func Return_Server_Stats() (Stats, error) {
	var stats Stats
	l := log.New(os.Stdout, fmt.Sprintf("%sGOMICRON >> %s", cmd.TCYAN, cmd.TRESET), log.LstdFlags)
	stat_cmd := exec.Command("python", "./scripts/python/status.py")

	output, err := stat_cmd.Output()
	if err != nil {
		l.Print("[ERROR]: ", err)
		return stats, err
	}

	err = json.Unmarshal(output, &stats)
	if err != nil {
		l.Print("[ERROR]: ", err)
		return stats, err
	}
	l.Print("<--- Server Stats --->")
	l.Printf("%sCPU_Percent: %.2f %s", cmd.TYELLOW, stats.CPU_Percent, cmd.TRESET)
	l.Printf("%sRAM_Percent: %.2f %s", cmd.TYELLOW, stats.RAM_Percent, cmd.TRESET)
	l.Printf("%sDisk_Percent: %.2f %s", cmd.TYELLOW, stats.Disk_Percent, cmd.TRESET)	
	l.Print("<-------------------->")

	return Stats{
		CPU_Percent: stats.CPU_Percent,
		RAM_Percent: stats.RAM_Percent,
		Disk_Percent: stats.RAM_Percent,
	}, nil
}
