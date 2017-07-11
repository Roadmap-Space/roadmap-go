package roadmap

import "fmt"

// Stories used for calling all story related endpoints
type Stories struct {
	EndpointURL string
}

// Create adds a new story in a specific roadmap
func (s *Stories) Create(story Story) (*Story, error) {
	path := fmt.Sprintf("%s", s.EndpointURL)

	var result Story
	if err := apiClient.post(path, story, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AddIdea creates and inserts a new idea in
func (s *Stories) AddIdea(storyID, storyToken string, idea Idea) (*Idea, error) {
	path := fmt.Sprintf("%s/%s/ideas", s.EndpointURL, idToURL(storyID, storyToken))

	var result Idea
	if err := apiClient.post(path, idea, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// StoryAttachItem contains properties to attach an item to a story
type StoryAttachItem struct {
	RoadmapID     string `json:"roadmapId"`
	StoryID       string `json:"parentId"`
	AttachedID    string `json:"id"`
	AttachedToken string `json:"token"`
}

// Attach inserts a feedback or and idea in the attachments list of a story
func (s *Stories) Attach(p StoryAttachItem) (*Idea, error) {
	path := fmt.Sprintf("%s/ideas", s.EndpointURL)

	var result Idea
	if err := apiClient.post(path, p, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// StoryListParams contains properties to filter roadmap stories
type StoryListParams struct {
	RoadmapID string
	Completed bool
}

// List returns the list of stories for a specific roadmap
func (s *Stories) List(p *StoryListParams) ([]Story, error) {
	completed := ""
	if p.Completed {
		completed = "/done"
	}

	path := fmt.Sprintf("%s/%s%s", s.EndpointURL, p.RoadmapID, completed)

	var result []Story
	if err := apiClient.get(path, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Update saves a story
func (s *Stories) Update(story *Story) (*Story, error) {
	path := s.EndpointURL

	var result Story
	if err := apiClient.put(path, story, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
