# Go Roadmap

Version: v1.2.0

## Summary

Official Go API client library for [Roadmap](https://roadmap.space)'s API.

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
