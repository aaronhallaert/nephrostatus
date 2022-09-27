# NephroStatus

Add NephroFlow API and Manager status in system tray.

<img width="172" alt="image" src="https://user-images.githubusercontent.com/37899722/192513810-29ac8381-75c9-4188-b6d7-8a7f120e2106.png">

## State

- API: curl `http://localhost:3000/api/version_info`
- Manager: check for open port 8080 (LISTEN)

## Actions

- Start API: `tmux send-keys -t nipro:api.1 "run_api.sh -- rails s\n"`
- Stop API: `docker ps | grep nephroflow/server | awk '{print $1}' | xargs -r docker stop`
- Stop Manager: `tmux send-keys -t nipro:manager.1 "C-c"`
- Start Manager: `tmux send-keys -t nipro:manager.1 yarn dev C-m`
- Open Manager
  - NephroFlow tab is activated with chrome-cli and chrome is focused with `open -a Google\ Chrome`

## Shortcomings

### Check open port with docker container

If the API web container is running (without an active rails server), the port is
already exposed so the status of the API seems to be `running`.

### Focus Chrome window

Opening NephroFlow manager will open the last used Google Chrome window.
If you use different profiles and multiple windows are active, it is possible
the wrong Chrome window will be opened. However, in the correct window,
the nephroflow tab will always be focused.

## Setup

### Dependencies

- [chrome-cli on Github](https://github.com/prasmussen/chrome-cli)

  - `brew install chrome-cli`

- Tmux
  - Currently, most of the commands to start/stop are based on `tmux send-keys`

### Installation

TODO

- Run with `go run main.go`

## TODO

- [ ] Make commands more generic
- [ ] Allow custom commands
- [ ] Add Github actions
