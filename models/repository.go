package models

import (
	"time"

	"github.com/src-d/go-kallax"
)

func newRepository() *Repository {
	return &Repository{Status: Pending}
}

// Repository represents a remote repository found on the Internet.
type Repository struct {
	kallax.Model `table:"repositories"`
	kallax.Timestamps `kallax:",inline"`
	// Endpoints is a slice of valid git endpoints to reach this repository.
	// For example, git://host/my/repo.git and https://host/my/repo.git.
	// They are meant to be endpoints of the same exact repository, and not
	// mirrors.
	Endpoints []string
	// Status is the fetch status of tge repository in our repository storage.
	Status FetchStatus
	// FetchedAt is the timestamp of the last time this repository was
	// fetched and archived in our repository storage successfully.
	FetchedAt time.Time
	// FetchErrorAt is the timestamp of the last fetch error, if any.
	FetchErrorAt time.Time
	// LastCommitAt is the last commit time found in this repository.
	LastCommitAt time.Time
	// References is the current slice of references as present in our
	// repository storage.
	References []*Reference
}

// FetchStatus represents the fetch status of this repository.
type FetchStatus string

const (
	// NotFound means that the remote repository was not found at the given
	// endpoints.
	NotFound FetchStatus = "not_found"
	// Fetched means that the remote repository was found, fetched and
	// successfully stored.
	Fetched = "fetched"
	// Pending is the default value, meaning that the repository has not
	// been fetched yet.
	Pending = "pending"
)
