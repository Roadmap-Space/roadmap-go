package roadmap_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	roadmap "github.com/roadmap-space/roadmap-go"
)

var c *roadmap.Client

var testRoadmapID = ""

func TestMain(m *testing.M) {
	os.Setenv("RM_DEBUG", "1")
	c = roadmap.New("dominic@roadmap.space", "57b8488ac7899e5cb4337ea1|ebd5361a-95d9-466f-8d94-26d157f9afea")
	if strings.Index(c.BasePath, "localhost") == -1 {
		fmt.Printf("base path is %s and should be localhost", c.BasePath)
		os.Exit(0)
	}

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
