package version

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

var (
	// Version supplies the semantic version
	Version string
)

func String() string {
	ver, err := semver.NewVersion("1.0")
	if err != nil {
		Version = "Dev build (no valid version)"
	}
	return fmt.Sprintf("Version : %s\n", ver)

}
