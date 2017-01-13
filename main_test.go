package roadmap_test

import (
	"fmt"
	"os"
	"testing"

	roadmap "github.com/roadmap-space/roadmap-go"
)

var c *roadmap.Client

var testRoadmapID = ""

func TestMain(m *testing.M) {
	c = roadmap.New("dominic@roadmap.space", "5829ded880309d0420adde35|307b31dd-06e8-434a-8c4c-9278f08c9276")

	r, err := createRoadmap("unit test")
	if err != nil {
		fmt.Println("Unable to create the test roadmap", err)
		os.Exit(0)
	}
	testRoadmapID = r.ID

	retval := m.Run()

	// cannot be in a defer
	c.Roadmaps.Delete(testRoadmapID)

	os.Exit(retval)
}
