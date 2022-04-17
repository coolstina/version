package verflag

import (
	"fmt"
	"testing"

	flag "github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
	flagSet *flag.FlagSet
}

func (suite *Suite) BeforeTest(suiteName, testName string) {
	suite.flagSet = &flag.FlagSet{}
}

func (suite *Suite) Test_AddFlags() {
	AddFlags(suite.flagSet)
	assert.Contains(suite.T(),
		fmt.Sprintf("%#+v\n", suite.flagSet), "pflag.Flag{\"version\"")
}
