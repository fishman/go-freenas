# go-freenas

[![Build Status](https://travis-ci.org/fishman/go-freenas.svg?branch=master)](https://travis-ci.org/fishman/go-freenas) 

go-freenas is a Go client library for accessing the [FreeNAS API](http://api.freenas.org).

## Usage ##

```go
import "github.com/fishman/go-freenas"
```

Construct a new FreeNAS client, then use the various services on the client to
access different parts of the FreeNAS API. For example:

```go
package main

import (
  "context"
  "fmt"

  freenas "github.com/fishman/go-freenas"
)

func main() {
  client := freenas.NewClient("http://freenas.local", "root", "freenas")
  // Turn on debugging
  client.Debug(true)

  shares, _, _ := client.NfsShares.List(context.Background())

  for _, element := range shares {
      fmt.Println(element.ID, element.Paths)
  }
}
```
## Ref ##
- google go-github client
