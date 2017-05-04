package main

import (
    "context"
    "fmt"
    "log"
    "github.com/ericchiang/k8s"
)

func main() {
    //fmt.Printf("TEST1")

    client, err := k8s.NewInClusterClient()
    if err != nil {
      log.Fatal(err)
    }
    pods, err := client.CoreV1().ListPods(ctx, client.Namespace)
    if err != nil {
      log.Fatal(err)
    }
    for _, pods := range pods.Items {
      fmt.Printf("%q", *pod.Metadata.Name)
    }
}
