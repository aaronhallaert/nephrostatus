package nephrodata

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

type NephroData struct {
	ApiStatus     bool
	ManagerStatus bool
}

func GetNephroData() *NephroData {
	var wg sync.WaitGroup
	wg.Add(2)

	apiStatus, managerStatus := false, false

	go func() {
		defer wg.Done()
		apiStatus = isPortOpen("3000")
	}()
	go func() {
		defer wg.Done()
		managerStatus = isPortOpen("8080")
	}()

	wg.Wait()

	return &NephroData{
		apiStatus,
		managerStatus,
	}
}

func isPortOpen(port string) bool {
	nValue, err := exec.Command("lsof", "-i", "-P", "-n").Output()

	if err != nil {
		fmt.Printf("error %s", err)
	}
	

	result := strings.TrimSuffix(string(nValue), "\n")
	return strings.Contains(result, port + " (LISTEN)")
}
