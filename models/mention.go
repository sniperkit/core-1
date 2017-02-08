package models

import "github.com/src-d/go-kallax"

// Mention is the sighting of a remote repository online.
type Mention struct {
	kallax.Model      `table:"mentions"`
	kallax.Timestamps `kallax:",inline"`
	// Endpoint is the repository URL as found.
	Endpoint string
	// Provider is the repository provider (e.g. github).
	Provider string
	// VCS contains the version control system of this Mention.
	VCS VCS
	// Context is arbitrary provider-dependent context of the sighting.
	Context map[string]string
}

type VCS string

const (
	// Git version control system
	GIT VCS = "git"
)
