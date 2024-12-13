package internal

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Open the browser after the connexion
func OpenBrowser(url string) error {
	var cmd *exec.Cmd
	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("système d'exploitation non supporté : %v", os)
	}
	return cmd.Start()
}
