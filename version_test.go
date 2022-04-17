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
	info Info
}

func (suite *VersionSuite) BeforeTest(suiteName, testName string) {
	suite.info = Get()
}

func (suite *VersionSuite) Test_Info_String() {
	actual := suite.info.String()
	assert.NotEmpty(suite.T(), actual)
	fmt.Println(actual)
}

func (suite *VersionSuite) Test_Info_ToJSON() {
	actual := suite.info.ToJSON()
	assert.NotEmpty(suite.T(), actual)
}
