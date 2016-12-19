package roadmap

import "fmt"

type Ideas struct {
	Endpoint string
}

// Create inserts a new idea for a specific roadmap
func (i *Ideas) Create(roadmapID, title string) (*Item, error) {
	path := fmt.Sprintf("%s", i.Endpoint)

	item := Item{RoadmpID: roadmapID, Title: title}
	if err := apiClient.post(path, item, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (i *Ideas) List(roadmapID string) ([]Item, error) {
  path := fmt.Sprintf("%s/list/%s", i.Endpoint, roadmapID)

  var ideas []Item
  if err := apiClient.get(path, &ideas); err != nil {
    return nil, err
  }
  return ideas, nil
}