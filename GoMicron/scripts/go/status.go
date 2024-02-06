package scripts

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"encoding/json"

	"github.com/AtahanPoyraz/cmd"
)

type Stats struct {
	CPU_Percent  float64 `json:"CPU_Percent"`
	RAM_Percent  float64 `json:"RAM_Percent"`
	Disk_Percent float64 `json:"Disk_Percent"`
}

func ServerStatus() error {
	var stats Stats
	l := log.New(os.Stdout, fmt.Sprintf("%sGOMICRON-API >> %s", cmd.TCYAN, cmd.TRESET), log.LstdFlags)
	command := exec.Command("python", "./scripts/python/status.py")

	output, err := command.Output()
	if err != nil {
		l.Printf("[ERROR]: %e", err)
		return fmt.Errorf("%e", err)
	}

	err = json.Unmarshal(output, &stats)
	if err != nil {
		l.Printf("[ERROR]: %v", err)
		return err
	}

	l.Print("<--- Server Stats --->")
	l.Printf("%sCPU_Percent: %.2f %s", cmd.TYELLOW, stats.CPU_Percent, cmd.TRESET)
	l.Printf("%sRAM_Percent: %.2f %s", cmd.TYELLOW, stats.RAM_Percent, cmd.TRESET)
	l.Printf("%sDisk_Percent: %.2f %s", cmd.TYELLOW, stats.Disk_Percent, cmd.TRESET)	
	l.Print("<-------------------->")

	return nil
}
