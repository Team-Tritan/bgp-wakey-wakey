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

var audioRunning bool

func main() {
	for {
		cmd := exec.Command(pingCommand, pingArguments, "1", targetHost)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Cannot reach router:", err)
			if !audioRunning {
				playAudio()
				audioRunning = true
			}
		} else {
			if audioRunning {
				stopAudio()
				audioRunning = false
			}
		}

		go monitorAudio()

		time.Sleep(time.Second * 60)
	}
}

func playAudio() {
	cmd := exec.Command(audioCommand, audioFile)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error playing audio:", err)
	}
}

func stopAudio() {
	cmd := exec.Command("pkill", audioCommand)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error stopping audio:", err)
	}
}

func monitorAudio() {
	cmd := exec.Command("pgrep", audioCommand)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Audio player is not running. Restarting audio.")
		playAudio()
	}
}
