package roadmap

import "fmt"

// Ideas used for calling all idea related endpoints
type Ideas struct {
	EndpointURL string
}

// Create inserts a new idea for a specific roadmap
func (i *Ideas) Create(item Idea) (*Idea, error) {
	path := fmt.Sprintf("%s", i.EndpointURL)

	if err := apiClient.post(path, item, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// IdeaListParams used when calling list
type IdeaListParams struct {
	IsArchived bool
}

// List returns the idea for a specific roadmai
func (i *Ideas) List(roadmapID string, params *IdeaListParams) ([]Idea, error) {
	if params == nil {
		params = &IdeaListParams{}
	}

	archived := ""
	if params.IsArchived {
		archived = "/archive"
	}

	path := fmt.Sprintf("%s/list/%s%s", i.EndpointURL, roadmapID, archived)

	var ideas []Idea
	if err := apiClient.get(path, &ideas); err != nil {
		return nil, err
	}
	return ideas, nil
}

// Get returns an idea by its ids
func (i *Ideas) Get(id, token string) (*Idea, error) {
	path := fmt.Sprintf("%s/%s", i.EndpointURL, idToURL(id, token))

	var idea Idea
	if err := apiClient.get(path, &idea); err != nil {
		return nil, err
	}
	return &idea, nil
}

// ToWidget marks this idea as visible in the widget
func (i *Ideas) ToWidget(id, roadmapID, token string) (bool, error) {
	return i.move(id, roadmapID, token, "widget")
}

// ToIdea returns a widget idea to the backlog
func (i *Ideas) ToIdea(id, roadmapID, token string) (bool, error) {
	return i.move(id, roadmapID, token, "idea")
}

// SetActive set an idea as active
func (i *Ideas) SetActive(id, roadmapID, token string) (bool, error) {
	return i.move(id, roadmapID, token, "active")
}

func (i *Ideas) move(id, roadmapID, token, state string) (bool, error) {
	path := fmt.Sprintf("%s/move/%s", i.EndpointURL, state)

	idea := &BaseItem{ID: id, RoadmapID: roadmapID, Token: token}
	var result bool
	if err := apiClient.post(path, idea, &result); err != nil {
		return false, err
	}
	return result, nil
}

// Update saves an idea
func (i *Ideas) Update(idea Idea) (*Idea, error) {
	path := fmt.Sprintf("%s", i.EndpointURL)

	var result Idea
	if err := apiClient.put(path, idea, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Convert transforms an idea into a story
func (i *Ideas) Convert(roadmapID, id, token string) (ok bool, err error) {
	path := fmt.Sprintf("%s/convert", i.EndpointURL)

	data := Idea{}
	data.RoadmapID = roadmapID
	data.ID = id
	data.Token = token

	err = apiClient.put(path, data, &ok)
	return
}

// Delete archives an idea
func (i *Ideas) Delete(id, token string) (bool, error) {
	path := fmt.Sprintf("%s/%s", i.EndpointURL, idToURL(id, token))

	var result bool
	if err := apiClient.delete(path, &result); err != nil {
		return false, err
	}
	return result, nil
}
