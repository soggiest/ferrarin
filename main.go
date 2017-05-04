package main

import (
    "context"
    "fmt"
    "log"
//    "github.com/ericchiang/k8s"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

func main() {
//fmt.Printf("TEST1")
  config, err := rest.InClusterConfig()
  if err != nil {
    panic(err.Error())
	}
  clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		
    if err != nil {
			panic(err.Error())
		}
	
    fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
  }

//    client, err := k8s.NewInClusterClient()
//    if err != nil {
//      log.Fatal(err)
//    }
//    pods, err := client.CoreV1().ListPods(ctx, client.Namespace)
//    if err != nil {
//      log.Fatal(err)
//    }
//    for _, pod := range pods.Items {
//      fmt.Printf("%q", *pod.Metadata.Name)
//    }
}
