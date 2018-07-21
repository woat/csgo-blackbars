package main

import (
	"syscall"
	"os"
	"os/exec"
)

func setCsgoRes() {
	binary, lookErr := exec.LookPath("nvidia-settings")
	if lookErr != nil {
		panic(lookErr)
	}

	csgoArgs := []string{ `nvidia-settings`, `--assign`, `currentmetamode=dvi-i-0: nvidia-auto-select +0+0, dvi-i-3: 1024x768 +1920+0 { viewportin=1024x768, viewportout=814x768+100+0"}`} 
	// originArgs := []string{ `nvidia-settings`, `--assign`, `currentmetamode=dvi-i-0: nvidia-auto-select +0+0, dvi-i-3: 1920x1080 +1920+0`} 

	env := os.Environ()

	execErr := syscall.Exec(binary, csgoArgs, env)
	if execErr != nil {
		panic(execErr)
	}
}

func startCsgo() {
	binary, lookErr := exec.LookPath(`steam`)
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{`steam`, `steam://run/730`} 

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func main() {
	startCsgo()
}
