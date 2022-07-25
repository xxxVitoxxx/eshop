package version

import (
	"fmt"
	"runtime"
)

var (
	Version   = "Unknown"
	BuildTime = "Unknown"
	Commit    = "Unknown"
)

func PrintInfo() {
	fmt.Printf("  Version: %s\n  Go Version: %s\n  Build Time: %s\n  Commit: %s\n",
		Version,
		runtime.Version(),
		BuildTime,
		Commit,
	)
}
