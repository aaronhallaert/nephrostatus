package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"nephrostatus/nephrodata"

	"github.com/rainu/go-command-chain"
	"nephrostatus/utils"

	"github.com/getlantern/systray"
)

func main() {
	onExit := func() {
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Loading...")
	mStartApi := systray.AddMenuItem("Start API", "Start API in tmux")
	mStopApi := systray.AddMenuItem("Stop API", "Stop the docker container")
	systray.AddSeparator()

	mStartManager := systray.AddMenuItem("Start manager", "Start manager in tmux")
	mStopManager := systray.AddMenuItem("Stop manager", "Kill yarn")
	mOpenManager := systray.AddMenuItem("Open manager", "Open manager chrome tab")
	systray.AddSeparator()

	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")

	currentNephroData := &nephrodata.NephroData{}
	currentNephroData = nephrodata.GetNephroData()

	updateTray(currentNephroData)

	go func() {
		for {
			select {
			case <-mStopApi.ClickedCh:
				apiStopOutput := &bytes.Buffer{}

				err := cmdchain.Builder().
					Join("docker", "ps").
					Join("grep", "nephroflow/server").
					Join("awk", "{print $1}").
					Join("xargs", "-r", "docker", "stop").
					Finalize().WithOutput(apiStopOutput).Run()

				if err != nil {
					fmt.Printf("error %s", err.Error())
					fmt.Printf("%s", apiStopOutput)
				}

				fmt.Printf("%s", apiStopOutput)
			case <-mStartApi.ClickedCh:
				_, err := exec.Command("tmux", "send-keys", "-t", "nipro:api.1", "run_api.sh -- rails s\n").Output()

				if err != nil {
					fmt.Printf("error %s", err)
				}
			case <-mStopManager.ClickedCh:
				_, err := exec.Command("tmux", "send-keys", "-t", "nipro:manager.1", "^C").Output()

				if err != nil {
					fmt.Printf("error %s", err)
				}
			case <-mStartManager.ClickedCh:
				_, err := exec.Command("tmux", "send-keys", "-t", "nipro:manager.1", "yarn dev\n").Output()

				if err != nil {
					fmt.Printf("error %s", err)
				}
			case <-mOpenManager.ClickedCh:
				tabOutput := utils.GetNephroflowTabId()

				_, errActTab := exec.Command("chrome-cli", "activate", "-t", tabOutput).Output()
				if errActTab != nil {
					panic(errActTab)
				}

				_, errOpenChrome := exec.Command("open", "-a", "Google Chrome").Output()
				if errOpenChrome != nil {
					panic(errOpenChrome)
				}
			}
		}
	}()

	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()

	go func() {
		for {
			currentNephroData = nephrodata.GetNephroData()
			updateTray(currentNephroData)
			time.Sleep(time.Millisecond * 5000)
		}
	}()
}

func updateTray(d *nephrodata.NephroData) {
	title := ""
	if d.ApiStatus.Online {
		title = title + "API: " + d.ApiStatus.Version + " ???"
	} else {
		title = title + "API ???"
	}

	title = title + " | "

	if d.ManagerStatus {
		title = title + "Manager ???"
	} else {
		title = title + "Manager ???"
	}

	systray.SetTitle(title)
}
