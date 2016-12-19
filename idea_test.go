package roadmap_test

import (
  "testing"

  roadmap "github.com/roadmap-space/roadmap-go"
)
func createIdea(roadmapID, title string) (*roadmap.Item, error) {
  return c.Ideas.Create(roadmapID, title)
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
