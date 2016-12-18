package roadmap

import (
	"fmt"
)

// Roadmaps is used to call actions for the Roadmaps resrouces
type Roadmaps struct {
	EndpointURL string
}

// Roadmap represents a Roadmap object
type Roadmap struct {
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Name      string `json:"name"`
}

// List returns a list of Roadmap for the autneticated account
func (r *Roadmaps) List() ([]Roadmap, error) {
	path := fmt.Sprintf("%s", r.EndpointURL)

	var roadmaps []Roadmap
	if err := apiClient.get(path, &roadmaps); err != nil {
		return nil, err
	}

	return roadmaps, nil
}

// GetWidgetIdeas returns the ideas available for the widget
func (r *Roadmaps) GetWidgetIdeas(roadmapID string) ([]Item, error) {
	path := fmt.Sprintf("%s/%s/widget", r.EndpointURL, roadmapID)

	var items []Item
	if err := apiClient.get(path, &items); err != nil {
		return nil, err
	}
	return items, nil
}
