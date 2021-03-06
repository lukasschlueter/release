package release

import (
	"github.com/Masterminds/semver"
	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

type Options struct {
	GitHubToken string
	Version     string
	Branch      string
	VersionFile string
	semVersion  *semver.Version
	// add your stuff here
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.GitHubToken) == 0 {
		opts.GitHubToken = envy.Get("GITHUB_TOKEN", "")
		if len(opts.GitHubToken) == 0 {
			return errors.New("you must set a GITHUB_TOKEN")
		}
	}
	if len(opts.Version) == 0 {
		opts.Version = "v0.0.1"
	}
	v, err := semver.NewVersion(opts.Version)
	if err != nil {
		return errors.WithStack(err)
	}
	opts.semVersion = v
	if len(opts.Branch) == 0 {
		opts.Branch = "master"
	}
	return nil
}
