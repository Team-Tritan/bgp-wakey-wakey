package main

import (
	"fmt"
	"os/exec"
	"time"
)

const (
	targetHost    = "core.as393577.net"
	pingCommand   = "ping"
	pingArguments = "-c"
	audioFile     = "wakeywakey.mp3"
	audioCommand  = "mpv"
)

func main() {
	for {
		cmd := exec.Command(pingCommand, pingArguments, "1", targetHost)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Cannot reach router: ", err)
			playAudio()
		}

		time.Sleep(time.Second * 60)
	}
}

func playAudio() {
	cmd := exec.Command(audioCommand, audioFile)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing audio:", err)
	}
}
