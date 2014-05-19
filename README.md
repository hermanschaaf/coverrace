CoverRace
=========

This is a simple demonstration of the perils of running `go test -cover -race`, i.e. the Go cover tool with the Go race detector at the same time. The lesson is to always also add the `-covermode=atomic` flag, like this: `go test -cover -covermode=atomic -race`.

For a full explanation, refer to the blog post on [http://herman.asia](http://herman.asia/running-the-go-race-detector-with-cover)