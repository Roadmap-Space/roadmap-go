package roadmap

import "fmt"

type Ideas struct {
	Endpoint string
}

// Create inserts a new idea for a specific roadmap
func (i *Ideas) Create(item Idea) (*Idea, error) {
	path := fmt.Sprintf("%s", i.Endpoint)

	if err := apiClient.post(path, item, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// IdeaListParams
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

  path := fmt.Sprintf("%s/list/%s%s", i.Endpoint, roadmapID, archived)

  var ideas []Idea
  if err := apiClient.get(path, &ideas); err != nil {
    return nil, err
  }
  return ideas, nil
}

// Delete archives an idea
func (i *Ideas) Delete(id, token string) (bool, error) {
	path := fmt.Sprintf("%s/%s", i.Endpoint, idToURL(id, token))

	var result bool
	if err := apiClient.delete(path, &result); err != nil {
		return false, err
	}
	return result, nil
}

func (i *Ideas) Get(id, token string) (*Idea, error) {
	path := fmt.Sprintf("%s/%s", i.Endpoint, idToURL(id, token))

	var idea Idea
	if err := apiClient.get(path, &idea); err != nil {
		return nil, err
	}
	return &idea, nil
}

// ToWidget
func (i *Ideas) ToWidget(id, roadmapID, token string) (bool, error) {
	return i.move(id, roadmapID, token, "widget")
}

func (i *Ideas) ToIdea(id, roadmapID, token string) (bool, error) {
	return i.move(id, roadmapID, token, "idea")
}

func (i *Ideas) SetActive(id, roadmapID, token string) (bool, error) {
	return i.move(id, roadmapID, token, "active")
}

func (i *Ideas) move(id, roadmapID, token, state string) (bool, error) {
	path := fmt.Sprintf("%s/move/%s", i.Endpoint, state)

	idea := &BaseItem{ID: id, RoadmapID: roadmapID, Token: token}
	var result bool
	if err := apiClient.post(path, idea, &result); err != nil {
		return false, err
	}
	return result, nil
}

func (i *Ideas) Update(idea Idea) (bool, error) {
	path := fmt.Sprintf("%s", i.Endpoint)

	var result bool
	if err := apiClient.put(path, idea, &result); err != nil {
		return false, err
	}
	return result, nil
}