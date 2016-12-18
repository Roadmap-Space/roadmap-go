package roadmap_test

import (
  "testing"

  roadmap "github.com/roadmap-space/roadmap-go"
)

func Test_GetAClientWithProperData(t *testing.T) {
  t.Parallel()

  testClient := roadmap.New("test", "test")

  if testClient.BasePath != "http://localhost:8070/v1" {
    t.Fail()
  }
}

func Test_Post(t *testing.T) {
  
}