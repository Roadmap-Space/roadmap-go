package roadmap_test

import (
	"fmt"
	"testing"

	roadmap "github.com/roadmap-space/roadmap-go"
)

func createRoadmap(title string) (*roadmap.Roadmap, error) {
	return c.Roadmaps.Create(title)
}

func Test_RoadmapsGetWidgetIdeas(t *testing.T) {
	t.Parallel()

	roadmaps, err := c.Roadmaps.List()
	if err != nil {
		t.Fatal(err)
	} else if len(roadmaps) < 1 {
		t.Fatal("We don't have any roadmap")
	}

	fmt.Println(roadmaps[0].ID)

	items, err := c.Roadmaps.GetWidgetIdeas(roadmaps[0].ID)
	if err != nil {
		t.Error(err)
	} else if len(items) < 1 {
		t.Errorf("The length was %d and we were looking for > 1", len(items))
	}
}

func Test_RoadmapsGetRoadmapList(t *testing.T) {
	t.Parallel()

	if _, err := c.Roadmaps.List(); err != nil {
		t.Error(err)
	}
}
