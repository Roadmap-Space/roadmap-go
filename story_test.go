package roadmap_test

import (
	"testing"

	roadmap "github.com/roadmap-space/roadmap-go"
)

func createStory(roadmapID, title string) (*roadmap.Story, error) {
	story := roadmap.Story{}
	story.RoadmapID = roadmapID
	story.Title = title

	return c.Stories.Create(story)
}

func Test_StoriesList(t *testing.T) {
	t.Parallel()

	story, err := createStory(testRoadmapID, "testing story list")
	if err != nil {
		t.Fatal(err)
	}

	p := &roadmap.StoryListParams{RoadmapID: testRoadmapID}
	check, err := c.Stories.List(p)
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, s := range check {
		if s.ID == story.ID {
			found = true
			break
		}
	}

	if found == false {
		t.Error("Story list does not return newly created story")
	}
}

