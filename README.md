# trello ![](https://travis-ci.org/corvuscrypto/trello.svg?branch=master)
Go API wrapper for Trello that allows you to use easy chaining to get access the model you want through Trello's RESTful API.
With the trello wrapper, getting the REST API url for a model, e.g. a board's name, is as easy as 
```go
trello.GetURL(trello.Board.ID("id of board").Name)
```

