package roadmap

import (
	"encoding/base64"
	"fmt"
	"log"
)

// Items is used for commong endpoints to all item types
type Items struct {
	EndpointURL string
}

func idToURL(itemID, token string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s|%s", itemID, token)))
}

// Add creates a full item with no alteration
func (i *Items) Add(item *Item) (*Item, error) {
	path := fmt.Sprintf("%s", i.EndpointURL)
	var result Item
	if err := apiClient.post(path, item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a full item via unique id and token
func (i *Items) Get(id, token string) (*Item, error) {
	path := fmt.Sprintf("%s/%s", i.EndpointURL, idToURL(id, token))

	var result Item
	if err := apiClient.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an item
func (i *Items) Delete(itemID, token string) error {
	path := fmt.Sprintf("%s/force/%s", i.EndpointURL, idToURL(itemID, token))
	var result bool
	if err := apiClient.delete(path, &result); err != nil {
		return err
	} else if !result {
		log.Println("Unable to delete this item")
	}

	return nil
}
