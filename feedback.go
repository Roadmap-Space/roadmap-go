package roadmap

import "fmt"

// Feedback used to call all feedback related endpoints
type Feedback struct {
	EndpointURL string
}

// Create adds a new feedback
func (f *Feedback) Create(item UserFeedback) (*UserFeedback, error) {
	path := fmt.Sprintf("%s", f.EndpointURL)

	if err := apiClient.post(path, item, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Get returns a feedback by its id and token
func (f *Feedback) Get(id, token string) (*UserFeedback, error) {
	path := fmt.Sprintf("%s/%s", f.EndpointURL, idToURL(id, token))
	var feedback UserFeedback
	if err := apiClient.get(path, &feedback); err != nil {
		return nil, err
	}
	return &feedback, nil
}

// List returns the feedback list for a specific roadmap
func (f *Feedback) List(roadmapID string) ([]Item, error) {
	path := fmt.Sprintf("%s/list/%s", f.EndpointURL, roadmapID)
	var result []Item
	if err := apiClient.get(path, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Convert converts a feedback into an idea
func (f *Feedback) Convert(roadmapID, feedbackID, token string) (bool, error) {
	path := fmt.Sprintf("%s/convert", f.EndpointURL)

	data := UserFeedback{}
	data.RoadmapID = roadmapID
	data.ID = feedbackID
	data.Token = token

	var result bool
	if err := apiClient.put(path, data, &result); err != nil {
		return false, err
	}
	return result, nil
}

// FeedbackAttachParams is used to attach a feedback to an idea or story
type FeedbackAttachParams struct {
	SourceID    string `json:"sourceId"`
	SourceToken string `json:"sourceToken"`
	ParentID    string `json:"parentId"`
	ParentToken string `json:"parentToken"`
}

// Attach adds a feedback to an existing idea or story
func (f *Feedback) Attach(p FeedbackAttachParams) (*UserFeedback, error) {
	path := fmt.Sprintf("%s/attach", f.EndpointURL)

	var result UserFeedback
	if err := apiClient.post(path, p, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UnLink removes the link created when feedback were attached to an idea or a story
func (f *Feedback) UnLink(id, token, parentID string) (ok bool, err error) {
	path := fmt.Sprintf("%s/%s/unlink/%s", f.EndpointURL, idToURL(id, token), parentID)

	var result bool
	if err = apiClient.delete(path, &result); err != nil {
		return
	}
	ok = result
	return
}

// Delete deletes a feedback
func (f *Feedback) Delete(id, token string) (bool, error) {
	path := fmt.Sprintf("%s/%s", f.EndpointURL, idToURL(id, token))

	var result bool
	err := apiClient.delete(path, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}
