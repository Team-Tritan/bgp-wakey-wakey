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
	pingInterval  = 5 * time.Second
	maxFailures   = 3
)

var (
	audioRunning     bool
	consecutiveFails int
)

func main() {
	for {
		cmd := exec.Command(pingCommand, pingArguments, "1", targetHost)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Cannot reach router:", err)
			consecutiveFails++
			if consecutiveFails >= maxFailures && !audioRunning {
				playAudio()
				audioRunning = true
				fmt.Println("Consecutive failures:", consecutiveFails)
			}
		} else {
			fmt.Println("Router heartbeat worky!")
			if audioRunning {
				stopAudio()
				audioRunning = false
			}
			consecutiveFails = 0
		}

		time.Sleep(pingInterval)
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
