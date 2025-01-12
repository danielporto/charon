// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

package version

import (
	"context"
	"runtime/debug"
	"strings"

	"github.com/obolnetwork/charon/app/errors"
	"github.com/obolnetwork/charon/app/log"
	"github.com/obolnetwork/charon/app/z"
)

const (
	// Version is the branch version of the codebase.
	//  - Main branch: v0.X-dev
	//  - Release branch: v0.X-rc
	// It is overridden by git tags when building official releases.
	Version = "v0.16-dev"
)

// Supported returns the supported minor versions in order of precedence.
func Supported() []string {
	return []string{
		// TODO(corver): Add v0.16 as part of https://github.com/ObolNetwork/charon/issues/2099
		"v0.15",
		"v0.14",
	}
}

// GitCommit returns the git commit hash and timestamp from build info.
func GitCommit() (hash string, timestamp string) {
	hash, timestamp = "unknown", "unknown"
	hashLen := 7

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return hash, timestamp
	}

	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			if len(s.Value) < hashLen {
				hashLen = len(s.Value)
			}
			hash = s.Value[:hashLen]
		} else if s.Key == "vcs.time" {
			timestamp = s.Value
		}
	}

	return hash, timestamp
}

// LogInfo logs charon version information along-with the provided message.
func LogInfo(ctx context.Context, msg string) {
	gitHash, gitTimestamp := GitCommit()
	log.Info(ctx, msg,
		z.Str("version", Version),
		z.Str("git_commit_hash", gitHash),
		z.Str("git_commit_time", gitTimestamp),
	)
}

// Minor returns the minor version of the provided version string.
func Minor(version string) (string, error) {
	split := strings.Split(version, ".")
	if len(split) < 2 {
		return "", errors.New("invalid version string")
	}

	return strings.Join(split[:2], "."), nil
}
