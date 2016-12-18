package roadmap_test

import (
  "testing"

  "github.com/roadmap-space/roadmap-go"
)

func createFeedback(roadmapID, title string) (*roadmap.Item, error) {
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

  if err = c.Feedback.Convert(f.ID, f.Token); err != nil {
    t.Error(err)
  }

  deleteItem(f.ID, f.Token)
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