package roadmap

import "fmt"

// Stories used for calling all story related endpoints
type Stories struct {
	EndpointURL string
}

// List returns the list of stories for a specific roadmap
func (s *Stories) List(roadmapID string) ([]Story, error) {
	path := fmt.Sprintf("%s/%s", s.EndpointURL, roadmapID)

	var result []Story
	if err := apiClient.get(path, &result); err != nil {
		return nil, err
	}
	return result, nil
}
