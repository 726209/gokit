package version

// Version is the current version of the CLI tool.
// It is updated automatically by standard-version.
var Version = "1.0.3-alpha.0"

func GetVersion() string {
	return Version
}
