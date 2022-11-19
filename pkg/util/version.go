package util

import "golang.org/x/mod/semver"

type Version string

func (v Version) Newer(r Version) bool {
	return semver.Compare(string(v), string(r)) == 1
}
