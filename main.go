package main

import (
	"os"
	"strings"

	"github.com/talal/mimir/pkg/prompt"
)

func main() {
	if len(os.Args) > 1 {
		os.Stderr.Write([]byte("\x1B[1;31mPrompt error: " + "no arguing with Mímir" + "\x1B[0m\n"))
		os.Exit(1)
	}

	var line []string
	line = appendUnlessEmpty(line, prompt.GetDir())
	if showKubeInfo := os.Getenv("MIMIR_KUBE"); showKubeInfo != "false" {
		line = appendUnlessEmpty(line, prompt.GetKube())
	}
	if showOSCloudInfo := os.Getenv("MIMIR_OS_CLOUD"); showOSCloudInfo != "false" {
		line = appendUnlessEmpty(line, prompt.GetOSCloud())
	}

	os.Stdout.Write([]byte("\n" + strings.Join(line, " ") + "\n"))
}

func appendUnlessEmpty(list []string, val string) []string {
	if val == "" {
		return list
	}

	return append(list, val)
}
