package main

import "github.com/TSF/TFSClient/infrastructure/delivery/shell"

func main() {
	shellService := shell.NewShellCommands()
	shellService.Start()
}
