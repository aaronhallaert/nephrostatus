package nephrodata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"sync"
)

type ApiVersionResponse struct {
	Version string
}

type ApiVersionInfo struct {
	Online  bool
	Version string
}

type NephroData struct {
	ApiStatus     ApiVersionInfo
	ManagerStatus bool
}

func GetNephroData() *NephroData {
	var wg sync.WaitGroup
	wg.Add(2)

	apiStatus, managerStatus := ApiVersionInfo{Online: false, Version: ""}, false

	go func() {
		defer wg.Done()
		apiStatus = fetchApiVersionInfo()
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

func fetchApiVersionInfo() ApiVersionInfo {
	resp, err := http.Get("http://localhost:3000/api/version_info")
	if err != nil {
		return ApiVersionInfo{Version: "", Online: false}
	}

	body, err := ioutil.ReadAll(resp.Body)
	var data ApiVersionResponse
	json.Unmarshal(body, &data)

	return ApiVersionInfo{Version: data.Version, Online: true}
}
func isPortOpen(port string) bool {
	nValue, err := exec.Command("lsof", "-i", "-P", "-n").Output()

	if err != nil {
		fmt.Printf("error %s", err)
	}

	result := strings.TrimSuffix(string(nValue), "\n")
	return strings.Contains(result, port+" (LISTEN)")
}
