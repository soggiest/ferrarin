package main

import (
  "os"
  "fmt"
//  "time"

   "k8s.io/client-go/kubernetes"
   "k8s.io/client-go/rest"
   "k8s.io/client-go/pkg/api/v1"
   "github.com/nicholas_lane/ferrarin/createpod"
)

func get_client() *Clientset {
  config, err := rest.InClusterConfig()
  if err != nil {
    panic(err.Error())
  }
  fmt.Printf(config.Host)
  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }

  return clientset
}

func main() {

  client := get_client()
  test1 := os.Getenv("CREATE_POD")
  fmt.Println(test1)
  if test1 != nil {
    create_pods := os.Getenv("CREATE_POD")
  }
  for {
    fmt.Println("DOES IT WORK?!")
    fmt.Println(client.Settings())
  }
}
