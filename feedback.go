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

// Get returns a feedback by its id and token
func (f *Feedback) Get(id, token string) (*Item, error) {
	path := fmt.Sprintf("%s/%s/todo-remove-this", f.Endpoint, idToURL(id, token))
	var feedback Item
	if err := apiClient.get(path, &feedback); err != nil {
		return nil, err
	}
	return &feedback, nil
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

type FeedbackMergeParams struct {
	ID              string `json:"id"`
	SourceID        string `json:"-"`
	SourceToken     string `json:"-"`
	TargetID        string `json:"targetId"`
	TargetRoadmapID string `json:"roadmapId"`
}

func (f *Feedback) Merge(params *FeedbackMergeParams) (bool, error) {
	path := fmt.Sprintf("%s/merge", f.Endpoint)

	if len(params.SourceID) == 0 || len(params.SourceToken) == 0 {
		return false, fmt.Errorf("you need to supply the source id and source token")
	} else if len(params.TargetID) == 0 {
		return false, fmt.Errorf("you need to supply the target id")
	} else if len(params.TargetRoadmapID) == 0 {
		return false, fmt.Errorf("you need to supply the roadmap id")
	}

	params.ID = idToURL(params.SourceID, params.SourceToken)

	var status bool
	if err := apiClient.put(path, params, &status); err != nil {
		return false, err
	}
	return status, nil
}
