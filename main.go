package main

import (
  "os"
  "fmt"
  "time"

   "k8s.io/client-go/kubernetes"
   "k8s.io/client-go/rest"
//   "k8s.io/client-go/pkg/api/v1"
//   "github.com/nicholas_lane/ferrarin/createpod"
)

func main() {
  config, err := rest.InClusterConfig()
  if err != nil {
    panic(err.Error())
  }
  fmt.Printf(config.Host)
  client, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }
  
  test1 := os.Getenv("CREATE_POD")
  fmt.Println(test1)
  if len(test1) > 0 {
    create_pods := os.Getenv("CREATE_POD")
    fmt.Println(len(create_pods))
  }
  for {
    fmt.Println("%v\n", client)
    time.Sleep(5 * time.Second)
  }
}
