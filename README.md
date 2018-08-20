# FreeNAS API client in golang



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

Based on the google go-github client
