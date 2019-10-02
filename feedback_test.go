package roadmap_test

import (
	"testing"

	roadmap "github.com/roadmap-space/roadmap-go"
)

func createFeedback(roadmapID, title string) (*roadmap.UserFeedback, error) {
	fb := roadmap.UserFeedback{}
	fb.RoadmapID = roadmapID
	fb.Title = title

	return c.Feedback.Create(fb)
}

func Test_FeedbackAdd(t *testing.T) {
	t.Parallel()

	f, err := createFeedback(testRoadmapID, "Unit Test")
	if err != nil {
		t.Fatal(err)
	} else if len(f.ID) == 0 {
		t.Fatalf("Feedback insertion failed, we have no id")
	}

	if err := deleteItem(f.ID, f.Token); err != nil {
		t.Fatal(err)
	}
}

func Test_FeedbackConvert(t *testing.T) {
	t.Parallel()

	f, err := createFeedback(testRoadmapID, "Unit Test Convert")
	if err != nil {
		t.Fatal(err)
	}
	defer deleteItem(f.ID, f.Token)

	if ok, err := c.Feedback.Convert(f.RoadmapID, f.ID, f.Token); err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("feedback conversion returned false")
	}
}

func Test_FeedbackGet(t *testing.T) {
	t.Parallel()

	f, err := createFeedback(testRoadmapID, "unit test for get")
	if err != nil {
		t.Fatal(err)
	}
	defer deleteItem(f.ID, f.Token)

	if test, err := c.Feedback.Get(f.ID, f.Token); err != nil {
		t.Fatal(err)
	} else if test == nil {
		t.Errorf("unable to find the created feedback")
	}
}

func Test_FeedbackList(t *testing.T) {
	t.Parallel()

	f, err := createFeedback(testRoadmapID, "Unit Test List")
	if err != nil {
		t.Fatal(err)
	}

	list, err := c.Feedback.List(testRoadmapID)
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, i := range list {
		if i.ID == f.ID {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("The new item %s was supposed to be returned in the list", f.ID)
	}
}

func Test_FeedbackAttach(t *testing.T) {
	t.Parallel()

	f, err := createFeedback(testRoadmapID, "unit test merge")
	if err != nil {
		t.Fatal(err)
	}

	idea, err := createIdea(testRoadmapID, "for attaching feedback")
	if err != nil {
		t.Fatal(err)
	}

	p := roadmap.FeedbackAttachParams{SourceID: f.ID,
		SourceToken: f.Token,
		ParentID:    idea.ID,
		ParentToken: idea.Token,
	}

	if _, err := c.Feedback.Attach(p); err != nil {
		t.Fatal(err)
	}

	check, err := c.Items.Get(idea.ID, idea.Token)
	if err != nil {
		t.Fatal(err)
	} else if check == nil {
		t.Fatal("cannot query for the check item")
	} else if len(check.Attached) == 0 {
		t.Fatalf("expecting > 0 attached, received %d", len(check.Attached))
	}
}

func Test_FeedbackUnLink(t *testing.T) {
	t.Parallel()

	idea, err := createIdea(testRoadmapID, "parent for attachment")
	if err != nil {
		t.Fatal(err)
	}

	f, err := createFeedback(testRoadmapID, "new feedback to attach")
	if err != nil {
		t.Fatal(err)
	}

	p := roadmap.FeedbackAttachParams{ParentID: idea.ID,
		ParentToken: idea.Token,
		SourceID:    f.ID,
		SourceToken: f.Token,
	}
	if _, err := c.Feedback.Attach(p); err != nil {
		t.Fatal(err)
	}

	check, err := c.Items.Get(idea.ID, idea.Token)
	if err != nil {
		t.Fatal(err)
	} else if check == nil {
		t.Fatal("unable to find the parent idea", idea.ID)
	} else if len(check.Attached) == 0 {
		t.Fatal("feedback was not attached to parent idea")
	}

	found := false
	for _, a := range check.Attached {
		if a.ID == f.ID {
			found = true
			break
		}
	}

	if found == false {
		t.Fatal("feedback was not found in the attached parent slice")
	}

	if _, err := c.Feedback.UnLink(f.ID, f.Token, idea.ID); err != nil {
		t.Fatal(err)
	}

	check, err = c.Items.Get(idea.ID, idea.Token)
	if err != nil {
		t.Fatal(err)
	} else if check == nil {
		t.Fatal("unable to find the parent idea", idea.ID)
	} else if len(check.Attached) > 0 {
		t.Fatal("feedback was still attached to parent idea")
	}

	found = false
	for _, a := range check.Attached {
		if a.ID == f.ID {
			found = true
			break
		}
	}

	if found {
		t.Error("feedback were found and should have been removed from parent idea")
	}
}

func Test_FeedbackDelete(t *testing.T) {
	t.Parallel()

	f, err := createFeedback(testRoadmapID, "unit test for delete")
	if err != nil {
		t.Fatal(err)
	}

	if ok, err := c.Feedback.Delete(f.ID, f.Token); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fatalf("archive feedback returned false")
	}

	check, err := c.Feedback.Get(f.ID, f.Token)
	if err != nil {
		t.Fatal(err)
	} else if check.IsDeleted == false {
		t.Errorf("The feedback was not archived")
	}
}
