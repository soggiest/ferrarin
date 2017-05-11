package main

import (
	"os"
	//  "time"
	"flag"
	"k8s.io/client-go/kubernetes"
	//   "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	//   "k8s.io/client-go/pkg/api/v1"
	"github.com/soggiest/ferrarin/createpod"
)

var (
	kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
)

func main() {
	//TODO: GET RID OF THE OUT OF CLUSTER CONFIG WHEN PUSHING INTO THE CLUSTER
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	//  config, err := rest.InClusterConfig()
	//  if err != nil {
	//    panic(err.Error())
	//  }
	fmt.Println(config.Host)
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	createPod := os.Getenv("CREATE_POD")
	if len(createPod) > 0 {
		if createPod == "true" {
			//fmt.Println(len(create_pod))
			serverDaemonSet := createpod.createServer(client)
		}
	}
}
