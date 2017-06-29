package roadmap_test

import (
	"testing"

	roadmap "github.com/roadmap-space/roadmap-go"
)

func createRoadmap(title string) (*roadmap.Roadmap, error) {
	return c.Roadmaps.Create(title)
}

func Test_RoadmapsGetWidgetIdeas(t *testing.T) {
	t.Parallel()

	// make sure we have one idea in widget
	idea, err := createIdea(testRoadmapID, "new in widget")
	if err != nil {
		t.Fatal("Unable to create idea", err)
	}

	if ok, err := c.Ideas.ToWidget(idea.ID, testRoadmapID, idea.Token); err != nil {
		t.Fatal("unable to tag an idea as in widget", err)
	} else if ok == false {
		t.Fatalf("api returned false for tagging idea to as in widget")
	}

	items, err := c.Roadmaps.GetWidgetIdeas(testRoadmapID)
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
