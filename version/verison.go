package version

import (
	"fmt"
	"runtime/debug"
)

func buildVersionInfoString() string {
	buildInfoMap := make(map[string]string)

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range buildInfo.Settings {
			buildInfoMap[setting.Key] = setting.Value
		}
	}

	vcsRevision := buildInfoMap["vcs.revision"]
	if len(vcsRevision) > 5 {
		vcsRevision = vcsRevision[0:5]
	}

	return fmt.Sprintf("GOOS=%v GOARCH=%v rev=%v", buildInfoMap["GOOS"], buildInfoMap["GOARCH"], vcsRevision)
}

var versionInfoStringValue = buildVersionInfoString()

func VersionInfoString() string {
	return versionInfoStringValue
}
