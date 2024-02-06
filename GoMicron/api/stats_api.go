package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/AtahanPoyraz/cmd"
)

type Stats struct {
	CPU_Percent  float64 `json:"CPU_Percent"`
	RAM_Percent  float64 `json:"RAM_Percent"`
	Disk_Percent float64 `json:"Disk_Percent"`
}

var l = log.New(os.Stdout, fmt.Sprintf("%sGOMICRON-API >> %s", cmd.TCYAN, cmd.TRESET), log.LstdFlags)

func Return_Server_Stats() (Stats, error) {
	var stats Stats

	stat_cmd := exec.Command("python", "./scripts/python/status.py")
	output, err := stat_cmd.Output()
	if err != nil {
		l.Printf("[ERROR]: %v", err)
		return stats, err
	}

	err = json.Unmarshal(output, &stats)
	if err != nil {
		l.Printf("[ERROR]: %v", err)
		return stats, err
	}

	return stats, nil
}

func ServerAPI(w http.ResponseWriter, r *http.Request) {
	s, err := Return_Server_Stats()
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR]: %v", err), http.StatusInternalServerError)
		l.Printf("[ERROR] %v", err)
		return
	}

	res, err := json.Marshal(s)
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR]: %v", err), http.StatusInternalServerError)
		l.Printf("[ERROR] %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}