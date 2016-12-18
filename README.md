# Go Roadmap

Side note: We are building this package live [here](https://www.livecoding.tv/dstpierre).

## Summary

Official Go API client library for Roadmap's API.

## Installation

```shell
go get -u github.com/roadmap-space/roadmap-go
```

## Documentation

For a comprehensive list of examples, check out the [API documentation](http://api.roadmap.space).

Here is some examples:

### Get the ideas shown on the widget

```go
import "github.com/roadmap-space/roadmap-go

client := roadmap.New("email", "token")
ideas, err := client.Roadmaps.GetWidgetIdeas("[roadmap id here]")
```