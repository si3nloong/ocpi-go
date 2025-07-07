package ocpi

import (
	"sort"

	"golang.org/x/mod/semver"
)

// VersionsData defines model for versions_data.
type Version struct {
	URL     string        `json:"url"`
	Version VersionNumber `json:"version"`
}

type Versions []Version

func (vs Versions) LatestMutualVersion(version VersionNumber) (Version, bool) {
	if len(vs) == 0 {
		return Version{}, false
	}
	mm := semver.MajorMinor("v" + string(version))
	for _, v := range vs {
		if semver.MajorMinor("v"+string(v.Version)) == mm {
			return v, true
		}
	}
	return Version{}, false
}

func (vs Versions) LatestVersion() (Version, bool) {
	if len(vs) == 0 {
		return Version{}, false
	}
	sort.Sort(vs)
	return vs[len(vs)-1], true
}

func (vs Versions) Len() int      { return len(vs) }
func (vs Versions) Swap(i, j int) { vs[i], vs[j] = vs[j], vs[i] }
func (vs Versions) Less(i, j int) bool {
	cmp := semver.Compare("v"+string(vs[i].Version), "v"+string(vs[j].Version))
	if cmp != 0 {
		return cmp < 0
	}
	return vs[i].Version < vs[j].Version
}
