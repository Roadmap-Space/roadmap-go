package roadmap_test

import (
	"os"
	"testing"

	roadmap "github.com/roadmap-space/roadmap-go"
)

var c *roadmap.Client

var testRoadmapID = "5829ded880309d0420adde36"

func TestMain(m *testing.M) {
	c = roadmap.New("dominic@roadmap.space", "5829ded880309d0420adde35|307b31dd-06e8-434a-8c4c-9278f08c9276")

	retval := m.Run()
	os.Exit(retval)
}
