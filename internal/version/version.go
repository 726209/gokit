package version

// Version is the current version of the CLI tool.
// It is updated automatically by standard-version.
var Version = "1.0.1"

func GetVersion() string {
	return Version
}
