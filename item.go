package roadmap

import (
  "fmt"
  "log"
	"time"
	"encoding/base64"
)

type Items struct {
  EndpointURL string
}

// Item represents
type Item struct {
	ID          string    `json:"id"`
	RoadmpID    string    `json:"roadmapId"`
	Category    string    `json:"category"`
	ItemType    int       `json:"type"`
	ColumnIndex int       `json:"column"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Hidden      bool      `json:"hidden"`
	Order       int       `json:"order"`
	Completed   bool      `json:"completed"`
	CompletedOn time.Time `json:"completedOn"`
	Token       string    `json:"token"`
}

func idToURL(itemID, token string) string {
  return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s|%s", itemID, token)))
}

// Delete
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