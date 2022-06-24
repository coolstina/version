package version

import (
	"bytes"
	"encoding/json"

	formatalign "github.com/coolstina/format-align"
)

// Version builder information structure.
type Version struct {
	AppName    string
	Version    string
	GoVersion  string
	GitCommit  string
	GitBranch  string
	DestOSARCH string // Destination OS linux/amd64
	FromOSARCH string // Building from darwin/amd64
	BuildTime  string
}

func (millet *Version) JSONString() string {
	data, _ := json.MarshalIndent(&millet, "", "  ")
	return string(data)
}

func (millet *Version) RawString() string {
	var buffer bytes.Buffer
	var format = formatalign.NewFormatAlign(
		formatalign.WithAlignment(formatalign.AlignmentOfRight),
		formatalign.WithSeparator(""),
		formatalign.WithPlaceholder(12),
	)

	buffer.Write(append([]byte(format.Format("AppName:", millet.AppName)), '\n'))
	buffer.Write(append([]byte(format.Format("Version:", millet.Version)), '\n'))
	buffer.Write(append([]byte(format.Format("GoVersion:", millet.GoVersion)), '\n'))
	buffer.Write(append([]byte(format.Format("GitCommit:", millet.GitCommit)), '\n'))
	buffer.Write(append([]byte(format.Format("GitBranch:", millet.GitBranch)), '\n'))
	buffer.Write(append([]byte(format.Format("DestOSARCH:", millet.DestOSARCH)), '\n'))
	buffer.Write(append([]byte(format.Format("FromOSARCH:", millet.FromOSARCH)), '\n'))
	buffer.Write(append([]byte(format.Format("BuildTime:", millet.BuildTime)), '\n'))

	return buffer.String()
}
