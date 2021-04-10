## eventbus

A simple Go package that provides an event bus to be reused across your applications.

This package is thread-safe, provides a simple interface to add new listeners for an event and a way reset all events.

Lastly, `EventData` includes a `Tags` field that can be used to add additional context for the receiver. This can be useful in cases where there are multiple services and specific targets need to be specified.

### Example

```go

import (
  "fmt"
  "github.com/ntden/eventbus"

  // ...
)

eventBus := eventbus.New()

// Register a listener for the "login" event with priority 1.
eventBus.AddListener("login", func(data eventbus.EventData) {
    user := data.Content.(User)
    fmt.Printf("User %s logged in.\n", user.Name)
}, 1)

// Register a listener for the "login" event with priority 2.
eventBus.AddListener("login", func(data eventbus.EventData) {
    user := data.Content.(User)

    if user.Age > 70 {
      fmt.Printf("User %s is so old!\n", user.Name)
    }
}, 2)

// Emit a "login" event with some data.
user := User{Name: "Alice", Age: 77}
data := eventbus.EventData{
    Name:    "login",
    Content: user,
}
eventBus.Emit(data)

// Output:
// User Alice is so old!
// User Alice has logged in.
```
