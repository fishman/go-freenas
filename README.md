# go-freenas

[![Build Status](https://travis-ci.org/google/go-github.svg?branch=master)](https://travis-ci.org/google/go-github) 

go-freenas is a Go client library for accessing the [FreeNAS API][].

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
      client := freenas.NewClient(
        &freenas.Config{
          Address:  "http://freenas.local",
          User:     "root",
          Password: "freenas",
        },
      )
      // Turn on debugging
      client.Debug(true)

      shares, _, _ := client.NfsShares.List(context.Background())

      for _, element := range shares {
          fmt.Println(element.ID, element.Paths)
      }

      share := freenas.NfsShare{
          Comment: "Kube",
          Paths: []string{"/mnt/kubernetes/persistentvol1"},
      }

      res, _, _ := client.NfsShares.Edit(context.Background(), 18, share)
      fmt.Println(res)
    }
```
## Ref ##
- google go-github client
