package version

// Version is the current version of the CLI tool.
// It is updated automatically by standard-version.
var Version = "0.0.5-alpha.2"

func GetVersion() string {
	return Version
}
