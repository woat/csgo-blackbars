package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	if len(os.Args) < 2 {
		checkRequirements()
	}

	csgoRes, originRes := readConfigSettings()

	applyCsgoRes(csgoRes).Run()
	startCsgo().Run()
	applyOriginRes(originRes).Run()
}

func checkRequirements() {
	err := false

	nvErr := checkForNvidia()
	if nvErr != nil {
		fmt.Println(nvErr)
		err = true
	}

	sErr := checkForSteam()
	if sErr != nil {
		fmt.Println(sErr)
		err = true
	}

	cErr := checkForCsbb()
	if cErr != nil {
		fmt.Println(cErr)
		err = true
	}

	if err {
		os.Exit(1)
	}

	fmt.Println("Requirements have been met.")
	fmt.Println("Make sure to write your settings in ~/.csbb")
	os.Exit(0)
}

func checkForNvidia() error {
	_, nvErr := exec.LookPath("nvidia-settings")
	if nvErr != nil {
		return errors.New("REQUIRED: 'nvidia-settings' not found. Please check " +
			"that you have installed NVIDIA X Server Settings or make sure your " +
			"PATH variable is correct.")
	}
	return nil
}

func checkForSteam() error {
	_, sErr := exec.LookPath("steam")
	if sErr != nil {
		return errors.New("REQUIRED: 'steam' not found. Please check that you " +
			"have installed Steam or make sure your PATH variable is correct.")
	}
	return nil
}

func checkForCsbb() error {
	u, _ := user.Current()
	f, err := os.Open(u.HomeDir + "/.csbb")
	if err != nil {
		return errors.New("NOT FOUND: Could not locate '~/.csbb'.")
	}

	s := bufio.NewScanner(f)
	settings := []string{}
	for s.Scan() {
		settings = append(settings, s.Text())
	}

	if len(settings) != 2 {
		return errors.New("INVALID CSBB: Parsed (" + string(len(settings)) +
			") but expected (2), make sure there are only 2 lines in your '~/.csbb'.")
	}

	return nil
}

func readConfigSettings() (string, string) {
	u, _ := user.Current()
	f, _ := os.Open(u.HomeDir + "/.csbb")
	defer f.Close()
	s := bufio.NewScanner(f)
	settings := []string{}
	for s.Scan() {
		settings = append(settings, s.Text())
	}
	return settings[0], settings[1]
}

func applyCsgoRes(arg string) *exec.Cmd {
	return exec.Command("nvidia-settings", "--assign", "currentmetamode="+arg)
}

func applyOriginRes(arg string) *exec.Cmd {
	return exec.Command("nvidia-settings", "--assign", "currentmetamode="+arg)
}

func startCsgo() *exec.Cmd {
	args := os.Args[1:]
	return exec.Command(args[0], args...)
}
