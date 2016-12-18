package roadmap

import "fmt"

// Feedback
type Feedback struct {
	Endpoint string
}

// Create adds a new feedback
func (f *Feedback) Create(roadmapID, title string) (*Item, error) {
	path := fmt.Sprintf("%s", f.Endpoint)

	item := Item{RoadmpID: roadmapID, Title: title}
	if err := apiClient.post(path, item, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// List
func (f *Feedback) List(roadmapID string) ([]Item, error) {
	path := fmt.Sprintf("%s/list/%s", f.Endpoint, roadmapID)
	var result []Item
	if err := apiClient.get(path, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Convert
func (f *Feedback) Convert(id, token string) error {
	apiID := idToURL(id, token)
	var result bool
	if err := apiClient.put(fmt.Sprintf("%s/convert/%s", f.Endpoint, apiID), nil, &result); err != nil {
		return err
	}
	return nil
}
