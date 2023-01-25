package steamid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDToID3(t *testing.T) {
	s := IDToID3("STEAM_0:0:59492810")
	assert.Equal(t, "[U:1:118985620]", s)
}

func TestIDToID64(t *testing.T) {
	s := IDToID64("STEAM_0:0:59492810")
	assert.Equal(t, "76561198079251348", s)
}

func TestID3ToID(t *testing.T) {
	s := ID3ToID("[U:1:118985620]")
	assert.Equal(t, "STEAM_0:0:59492810", s)
}

func TestID3ToID64(t *testing.T) {
	s := ID3ToID64("[U:1:118985620]")
	assert.Equal(t, "76561198079251348", s)
}

func TestID64ToID(t *testing.T) {
	s := ID64ToID("76561198079251348")
	assert.Equal(t, "STEAM_0:0:59492810", s)
}

func TestID64ToID3(t *testing.T) {
	s := ID64ToID3("76561198079251348")
	assert.Equal(t, "[U:1:118985620]", s)
}
