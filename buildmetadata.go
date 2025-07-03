package buildmetadata

import (
	"runtime/debug"
)

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
