package services

import (
	"fmt"
	"os/exec"
)

func CollectMetrics() (map[string]string, error) {
	cpuUsage, err := exec.Command("sh", "-c", "top -bn1 | grep 'Cpu(s)' | sed 's/.*, *\([0-9.]*\)%* id.*/\1/'").Output()
	if err != nil {
		return nil, err
	}

	metrics := map[string]string{
		"CPU": string(cpuUsage),
	}
	return metrics, nil
}