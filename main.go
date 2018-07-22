package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"bufio"
)

var csgoRes string
var originRes string

func checkForNvidia() error {
	_, nvErr := exec.LookPath("nvidia-settings")
	if nvErr != nil {
		fmt.Println("REQUIRED: 'nvidia-settings' not found. Please check that you have installed NVIDIA X Server Settings or make sure your PATH variable is correct.")
		return errors.New("gg")
	}
	return nil
}

func checkForSteam() error {
	_, sErr := exec.LookPath("steam")
	if sErr != nil {
		fmt.Println("REQUIRED: 'steam' not found. Please check that you have installed Steam or make sure your PATH variable is correct.")
		return errors.New("gg")
	}
	return nil
}

func applyCsgoRes(arg string) {
	exec.Command(`nvidia-settings`, `--assign`, `currentmetamode=`+arg).Run()
}

func applyOriginRes(arg string) {
	exec.Command(`nvidia-settings`, `--assign`, `currentmetamode=`+arg).Run() 
}

func startCsgo() {
	args := os.Args[1:]
	exec.Command(args[0], args...).Run()
}

func getConfigSettings() {
	u, _ := user.Current()
	f, _ := os.Open(u.HomeDir + "/.csbb")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	settings := []string{}
	for scanner.Scan() {
		settings = append(settings, scanner.Text())
	}
	csgoRes = settings[0]
	originRes = settings[1]
}

func main() {
	if len(os.Args) < 2 {
		nvErr := checkForNvidia()
		sErr := checkForSteam()
		if nvErr == nil && sErr == nil {
			fmt.Println("Requirements have been met.")
			fmt.Println("Make sure to write your settings in ~/.csbb")
			os.Exit(3)
		}
	}


	getConfigSettings()
	
	u, _ := user.Current()
	f, _ := os.Open(u.HomeDir + "/.csbb")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	settings := []string{}
	for scanner.Scan() {
		settings = append(settings, scanner.Text())
	}

	csgoRes = settings[0]
	originRes = settings[1]
	applyCsgoRes(csgoRes)
	startCsgo()
	applyOriginRes(originRes)
}
