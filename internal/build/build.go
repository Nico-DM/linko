// Package build houses build-time variables set via -ldflags.
package build

// default build-time variables
var (
	GitSHA    = "unknown"
	BuildTime = "unknown"
)
