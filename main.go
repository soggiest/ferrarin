package main

import (
    "context"
    "fmt"
    "log"
    "github.com/ericchiang/k8s"
)

func main() {
    fmt.Printf("TEST1")
    #    client, err := k8s.NewInClusterClient()
    #    if err != nil {
    #        log.Fatal(err)
    #    }
    #
    #    nodes, err := client.CoreV1().ListNodes(context.Background())
    #    if err != nil {
    #        log.Fatal(err)
    #    }
    #    for _, node := range nodes.Items {
    #        fmt.Printf("name=%q schedulable=%t\n", *node.Metadata.Name, !*node.Spec.Unschedulable)
    #    }
}
