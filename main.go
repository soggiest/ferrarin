package main

import (
	"fmt"
	"os"
	//  "time"
	"flag"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	//   "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	//   "k8s.io/client-go/pkg/api/v1"
	"github.com/soggiest/ferrarin/createpod"
	//v1 "k8s.io/client-go/pkg/api/v1"
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
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	createPod := os.Getenv("CREATE_POD")
	if len(createPod) != 0 {
		if createPod == "true" {
			//fmt.Println(len(create_pod))
			serverDaemonSet := createpod.CreateServer(client)
			fmt.Printf("%s\n", serverDaemonSet.ObjectMeta.Name)
			//	fmt.Printf("%+v\n", serverDaemonSet.xConditions)
		} else {
			fmt.Println("CREATE_POD environment variable set to false, skipping test")
			glog.Info("CREATE_POD environment variable set to false, skipping test")
		}
	} else {
		fmt.Println("CREATE_POD environment variable missing, skipping test")
		glog.V(2).Infof("CREATE_POD environment variable missing, skipping test")
	}

}
