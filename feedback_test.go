package roadmap_test

import (
	"testing"

	"github.com/roadmap-space/roadmap-go"
)

func createFeedback(roadmapID, title string) (*roadmap.UserFeedback, error) {
	return c.Feedback.Create(roadmapID, title)
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

	if ok, err := c.Feedback.Convert(testRoadmapID, f.ID, f.Token); err != nil {
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

func Test_FeedbackMerge(t *testing.T) {
	t.Parallel()

	ideas, err := c.Ideas.List(testRoadmapID, nil)
	if err != nil {
		t.Fatal(err)
	} else if len(ideas) == 0 {
		t.Fatalf("no idea to pick for target")
	}

	f, err := createFeedback(testRoadmapID, "unit test merge")
	if err != nil {
		t.Fatal(err)
	}

	p := &roadmap.FeedbackMergeParams{SourceID: f.ID, SourceToken: f.Token, TargetID: ideas[0].ID, TargetRoadmapID: testRoadmapID}
	if ok, err := c.Feedback.Merge(p); err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Errorf("the merge failed")
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