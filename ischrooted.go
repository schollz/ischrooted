package ischrooted

import (
	"os/exec"
	"runtime"
)

// IsChrooted returns whether or not the current system is chrooted
func IsChrooted() (yes bool) {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("stat", "-c", "%d:%i", "/proc/1/root")
		stdoutStderr, _ := cmd.CombinedOutput()
		if len(stdoutStderr) > 2 {
			yes = string(stdoutStderr[0:2]) == "2:"
		}
	}
	return
}
