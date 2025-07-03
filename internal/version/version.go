package version

// Version is the current version of the CLI tool.
// It is updated automatically by standard-version.
var Version = "0.0.5-alpha.7"

func GetVersion() string {
	return Version
}
