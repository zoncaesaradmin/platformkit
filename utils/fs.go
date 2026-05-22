package utils

import (
	"corekit/fsutil"
	"os"
)

// Deprecated: use corekit/fsutil.EnsureDir.
// EnsureDir creates a directory path if it does not already exist.
// Empty paths are ignored so callers can safely pass optional env-derived paths.
func EnsureDir(path string, perm os.FileMode) error {
	return fsutil.EnsureDir(path, perm)
}
