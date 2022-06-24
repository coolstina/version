package version

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestVersionSuite(t *testing.T) {
	suite.Run(t, &VersionSuite{})
}

type VersionSuite struct {
	suite.Suite
	millet *Version
}

func (suite *VersionSuite) BeforeTest(suiteName, testName string) {
	suite.millet = &Version{
		AppName:    "millet",
		Version:    "v1.2.3-beta",
		GoVersion:  "go1.18.3",
		GitCommit:  "fb557a8447f471f394cd637b6ae86c5612d7998c",
		GitBranch:  "release",
		DestOSARCH: "linux/amd64",
		FromOSARCH: "darwin/amd64",
		BuildTime:  "2022-06-23T16:57:14+08:00",
	}
}

func (suite *VersionSuite) Test_RawString() {
	expected := `    AppName: millet
    Version: v1.2.3-beta
  GoVersion: go1.18.3
  GitCommit: fb557a8447f471f394cd637b6ae86c5612d7998c
  GitBranch: release
 DestOSARCH: linux/amd64
 FromOSARCH: darwin/amd64
  BuildTime: 2022-06-23T16:57:14+08:00
`
	actual := suite.millet.RawString()
	assert.Equal(suite.T(), expected, actual)
	fmt.Println(actual)
}

func (suite *VersionSuite) Test_JSONString() {
	expected := `{
  "AppName": "millet",
  "Version": "v1.2.3-beta",
  "GoVersion": "go1.18.3",
  "GitCommit": "fb557a8447f471f394cd637b6ae86c5612d7998c",
  "GitBranch": "release",
  "DestOSARCH": "linux/amd64",
  "FromOSARCH": "darwin/amd64",
  "BuildTime": "2022-06-23T16:57:14+08:00"
}`
	actual := suite.millet.JSONString()
	assert.Equal(suite.T(), expected, actual)
	fmt.Println(actual)
}
