// Package version supplies version information collected at build time to
// apimachinery components.
package version

import (
	"fmt"
	"runtime"

	"github.com/coolstina/json"
	"github.com/gosuri/uitable"
)

var (
	// GitVersion is semantic version.
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ').
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit sha1 from git, output of $(git rev-parse HEAD).
	GitCommit = "$Format:%H$"
	// GitTreeState state of git tree, either "clean" or "dirty".
	GitTreeState = ""
)

// Info contains versioning information.
type Info struct {
	GitVersion   string `json:"GitVersion"`
	GitCommit    string `json:"GitCommit"`
	GitTreeState string `json:"GitTreeState"`
	BuildDate    string `json:"BuildDate"`
	GoVersion    string `json:"GoVersion"`
	Compiler     string `json:"Compiler"`
	Platform     string `json:"Platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	if s, err := info.Text(); err == nil {
		return string(s)
	}

	return info.GitVersion
}

// ToJSON returns the JSON string of version information.
func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

// Text encodes the version information into UTF-8-encoded text and
// returns the result.
func (info Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("GitVersion:", info.GitVersion)
	table.AddRow("GitCommit:", info.GitCommit)
	table.AddRow("GitTreeState:", info.GitTreeState)
	table.AddRow("BuildDate:", info.BuildDate)
	table.AddRow("GoVersion:", info.GoVersion)
	table.AddRow("Compiler:", info.Compiler)
	table.AddRow("Platform:", info.Platform)

	return table.Bytes(), nil
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
