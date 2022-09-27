package utils

import ( 
    "bytes"
	"github.com/rainu/go-command-chain"
 )

func GetNephroflowTabId() string {
    tabOutput := &bytes.Buffer{}

    err := cmdchain.Builder().
        Join("chrome-cli", "list", "tabs").
        Join("grep", "NephroFlow").
        Join("sed", "-n",  "s/\\[.*:\\(.*\\)\\] .*/\\1/p").
        Finalize().WithOutput(tabOutput).Run()

    if err != nil {
        panic(err)
    } else {
        return tabOutput.String()
    }
}

func GetNephroflowWindowId() string {
    windowOutput := &bytes.Buffer{}

    err := cmdchain.Builder().
        Join("chrome-cli", "list", "tabs").
        Join("grep", "NephroFlow").
        Join("sed", "-n",  "s/\\[\\(.*\\):.*\\] .*/\\1/p").
        Finalize().WithOutput(windowOutput).Run()

    if err != nil {
        panic(err)
    } else {
        return windowOutput.String()
    }
}
