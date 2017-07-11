package roadmap_test

import (
	"fmt"
	"testing"

	roadmap "github.com/roadmap-space/roadmap-go"
)

func createIdea(roadmapID, title string) (*roadmap.Idea, error) {
	idea := roadmap.Idea{}
	idea.RoadmapID = roadmapID
	idea.Title = title

	return c.Ideas.Create(idea)
}

func Test_IdeaCreate(t *testing.T) {
	i, err := createIdea(testRoadmapID, "unit test idea")
	if err != nil {
		t.Fatal(err)
	} else if i == nil {
		t.Errorf("create idea returned a nil pointer")
	}

	deleteItem(i.ID, i.Token)
}

func Test_IdeaList(t *testing.T) {
	t.Parallel()

	i, err := createIdea(testRoadmapID, "unit test list")
	if err != nil {
		t.Fatal(err)
	}
	defer deleteItem(i.ID, i.Token)

	list, err := c.Ideas.List(testRoadmapID, nil)
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, idea := range list {
		if idea.ID == i.ID {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("List is not returning our newly created idea")
	}
}

func Test_IdeaListArchived(t *testing.T) {
	t.Parallel()

	i, err := createIdea(testRoadmapID, "unit test list archived")
	if err != nil {
		t.Fatal(err)
	}
	defer deleteItem(i.ID, i.Token)

	if ok, err := c.Ideas.Delete(i.ID, i.Token); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Errorf("the delete idea return false")
	}

	list, err := c.Ideas.List(testRoadmapID, &roadmap.IdeaListParams{IsArchived: true})
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, idea := range list {
		if idea.ID == i.ID {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Unable to find the archived idea in archived list")
	}
}

func Test_IdeaGet(t *testing.T) {
	t.Parallel()

	i, err := createIdea(testRoadmapID, "unit test get idea")
	if err != nil {
		t.Fatal(err)
	}
	defer deleteItem(i.ID, i.Token)

	idea, err := c.Ideas.Get(i.ID, i.Token)
	if err != nil {
		t.Fatal(err)
	} else if idea == nil {
		t.Errorf("unable to get the idea")
	}
}

func Test_IdeaToWidget(t *testing.T) {
	t.Parallel()

	i, err := createIdea(testRoadmapID, "unit test to widget")
	if err != nil {
		t.Fatal(err)
	}
	defer deleteItem(i.ID, i.Token)

	if ok, err := c.Ideas.ToWidget(i.ID, testRoadmapID, i.Token); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Error("unable to set the idea to be on widget")
	}

	//TODO: this kind of gave some issue during live stream
	// not obvious that this field was needed to be true
	i.Published = true

	if _, err := c.Ideas.Update(*i); err != nil {
		t.Fatal(err)
	}

	widget, err := c.Roadmaps.GetWidgetIdeas(testRoadmapID)
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, w := range widget {
		fmt.Println(w.ID, i.ID)
		if w.ID == i.ID {
			found = true
			break
		}
	}

	if !found {
		t.Error("new idea not found on widget")
	}
}

func Test_IdeaUpdate(t *testing.T) {
	t.Parallel()

	i, err := createIdea(testRoadmapID, "test unit update")
	if err != nil {
		t.Fatal(err)
	}
	defer deleteItem(i.ID, i.Token)

	i.Description = "## this is markdown"

	if u, err := c.Ideas.Update(*i); err != nil {
		t.Fatal(err)
	} else if u == nil {
		t.Errorf("unable to update the idea")
	}
}

func Test_IdeaConvert(t *testing.T) {
	t.Parallel()

	idea, err := createIdea(testRoadmapID, "new idea to convert")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.Ideas.Convert(testRoadmapID, idea.ID, idea.Token); err != nil {
		t.Fatal(err)
	}

	p := &roadmap.StoryListParams{RoadmapID: testRoadmapID}
	check, err := c.Stories.List(p)
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, s := range check {
		if s.ID == idea.ID {
			found = true
			break
		}
	}

	if found == false {
		t.Error("the idea converted to story were not found in the roadmap stories")
	}
}
