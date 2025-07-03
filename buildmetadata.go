package buildmetadata

import (
	"runtime/debug"
	"strings"
)

// GitCommitSHA will read the `vcs.revision` metadata from the build info
// symbols stored in the binary, or return "HEAD" if not found.
// For more information on the build info symbols, see:
// https://pkg.go.dev/runtime/debug#ReadBuildInfo
func GitCommitSHA() string {
	if info, ok := debug.ReadBuildInfo(); ok && info != nil {
		for _, setting := range info.Settings {
			if strings.EqualFold(setting.Key, "vcs.revision") {
				return setting.Value
			}
		}
	}
	return "HEAD"
}
