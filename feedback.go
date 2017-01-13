package roadmap

import "fmt"

// Feedback
type Feedback struct {
	Endpoint string
}

// Create adds a new feedback
func (f *Feedback) Create(item UserFeedback) (*UserFeedback, error) {
	path := fmt.Sprintf("%s", f.Endpoint)

	if err := apiClient.post(path, item, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Get returns a feedback by its id and token
func (f *Feedback) Get(id, token string) (*UserFeedback, error) {
	path := fmt.Sprintf("%s/%s/todo-remove-this", f.Endpoint, idToURL(id, token))
	var feedback UserFeedback
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
func (f *Feedback) Convert(roadmapID, id, token string) (bool, error) {
	path := fmt.Sprintf("%s/convert", f.Endpoint)

	var data = new(struct{
		RoadmapID string `json:"roadmapId"`
		FeedbackID string `json:"feedbackId"`
		Token string `json:"token"`
	})

	data.RoadmapID = roadmapID
	data.FeedbackID = id
	data.Token = token

	var result bool
	if err := apiClient.put(path, data, &result); err != nil {
		return false, err
	}
	return result, nil
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

func (f *Feedback) Delete(id, token string) (bool, error) {
	path := fmt.Sprintf("%s/%s", f.Endpoint, idToURL(id, token))

	var result bool
	err := apiClient.delete(path, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}