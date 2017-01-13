package roadmap

import (
	"fmt"
)

// Roadmaps is used to call actions for the Roadmaps resrouces
type Roadmaps struct {
	EndpointURL string
}

// Create adds a roadmap
func (r *Roadmaps) Create(title string) (*Roadmap, error) {
	path := fmt.Sprintf("%s", r.EndpointURL)

	roadmap := Roadmap{Title: title}
	if err := apiClient.post(path, roadmap, &roadmap); err != nil {
		return nil, err
	}
	return &roadmap, nil
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
func (r *Roadmaps) GetWidgetIdeas(roadmapID string) ([]Idea, error) {
	path := fmt.Sprintf("%s/%s/widget", r.EndpointURL, roadmapID)

	var ideas []Idea
	if err := apiClient.get(path, &ideas); err != nil {
		return nil, err
	}
	return ideas, nil
}

// Delete permanently remove a rodamap and all related resources
func (r *Roadmaps) Delete(id string) (bool, error) {
	path := fmt.Sprintf("%s/%s", r.EndpointURL, id)

	var result bool
	if err := apiClient.delete(path, &result); err != nil {
		return false, err
	}
	return result, nil
}