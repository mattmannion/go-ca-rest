package service

import (
	"_/cmd/api/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServiceLayer(t *testing.T) {
	assert.IsType(t, &ServiceLayer{}, NewServiceLayer(Deps{RepoLayer: repo.RepoLayer{}}))
}
