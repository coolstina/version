package verflag

import (
	"fmt"
	"os"
	"strconv"

	"github.com/coolstina/version"

	flag "github.com/spf13/pflag"
)

// Version related constant definition
const (
	VersionFalse versionValue = iota
	VersionTrue
	VersionRaw
)

// private constant's related definition
const (
	strRawVersion   string = "raw"
	versionFlagName string = "version"
)

// private variable's related definition
var (
	versionFlag = Version(versionFlagName, VersionFalse,
		"Print version information and quit")
)

// versionValue version value definition
type versionValue int

// String implement Stringer interface
func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}

	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}

// Set version value
func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}

	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}

	return err
}

// Get return version value
func (v *versionValue) Get() interface{} {
	return v
}

// IsBoolFlag check version value Bool flag
func (v *versionValue) IsBoolFlag() bool {
	return true
}

// Type the type of the flag as required by the pflag.Value interface
func (v *versionValue) Type() string {
	return "version"
}

// VersionVar defines a flag with the specified name and usage string
func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	flag.Var(p, name, usage)
	// "--version" will be treated as "--version=true"
	flag.Lookup(name).NoOptDefVal = "true"
}

// Version wraps the VersionVar function
func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)
	return p
}

// AddFlags registers this package's flags on arbitrary FlagSet,
// such that they point to the same value as the global flags
func AddFlags(fs *flag.FlagSet) {
	fs.AddFlag(flag.Lookup(versionFlagName))
}

// PrintAndExitIfRequested will check if the -version flag was passed and, if so,
// print the version and exit
func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
