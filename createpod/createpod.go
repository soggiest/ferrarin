package createpod

import (
	"fmt"
	"time"
	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/pkg/api/v1"
	v1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	//"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/kubernetes/pkg/api"
	//	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/client-go/pkg/util/intstr"
)

func CreateServer(client *kubernetes.Clientset) *v1beta1.DaemonSet {
	// Check whether the test pods daemonset already exists, and if it does clean it up before proceeding.
	dsCheck, _  := client.ExtensionsV1beta1().DaemonSets("default").Get("test-pods-server")
	if len(dsCheck.ObjectMeta.Name) > 0  {
		fmt.Println("Test Pods DaemonSet already exists, removing it")
		//deleteOptions := v1.DeleteOptions{GracePeriodSeconds: new(int64)
		err := client.ExtensionsV1beta1().DaemonSets("default").Delete("test-pods-server", &v1.DeleteOptions{GracePeriodSeconds: new(int64)})
		delList := v1.ListOptions{LabelSelector: "test-pods"}
		delPods, _ := client.CoreV1().Pods("default").List(delList)
		for _, delPod := range delPods.Items {
			client.CoreV1().Pods("default").Delete(delPod.ObjectMeta.Name, &v1.DeleteOptions{GracePeriodSeconds: new(int64)})
		}
		time.Sleep(10 * time.Second)
		if err != nil {
			panic(err.Error())
		}
	}
	daemonSet := generateServerConfig()
	daemonSetObject, err := client.ExtensionsV1beta1().DaemonSets("default").Create(daemonSet)
	if err != nil {
		//TODO: Figure out how to better handle certain errors, such as "unexpected EOF"
		panic(err.Error())
	}
	time.Sleep(1 * time.Second)
	
		dsGet, err := client.ExtensionsV1beta1().DaemonSets("default").Get("test-pods-server")
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println("%v - %v\n", dsGet.Status.NumberReady, dsGet.Status.DesiredNumberScheduled)
		//TODO: FIGURE OUT DSGET BASED FOR LOOP WHERE IT CHECKS DAEMONSTATE STATUS EACH RUN
		for dsGet.Status.NumberReady < dsGet.Status.DesiredNumberScheduled {
			fmt.Printf("%d is less than desired number of %d\n", dsGet.Status.NumberReady, dsGet.Status.DesiredNumberScheduled)
			lo := v1.ListOptions{LabelSelector: "test-pods"}
			pods, err := client.CoreV1().Pods("default").List(lo)
			if err != nil {
				panic(err.Error())
			}
			//TODO: This part didn't seem to display, anything?
			for _, pod := range pods.Items {
				//if pod.Status.Phase != "Running" {	
					for _, conditions := range pod.Status.Conditions {
						//fmt.Printf("%+v\n\n", pod)
						if conditions.Type == "Ready" && conditions.Status == "False" {
							fmt.Printf("%s is failing its readiness check\n", pod.ObjectMeta.Name)
						}
					}
				//}
			}
			//			fmt.Printf("%+v\n", dsGet)
		}
	//}

	return daemonSetObject
}

func generateServerConfig() *v1beta1.DaemonSet {
	labels := map[string]string{"test-pods": "server"}
	daemonset := &v1beta1.DaemonSet{
		ObjectMeta: v1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pods-server",
			Labels:    labels,
		},
		Spec: v1beta1.DaemonSetSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					ServiceAccountName: "default",
					ImagePullSecrets: []v1.LocalObjectReference{
						{Name: "coreos-pull-secret"},
					},
					Containers: []v1.Container{
						{
							Name:            "server",
							Image:           "docker.io/nginx:latest",
							ImagePullPolicy: "Always",
							Ports: []v1.ContainerPort{
								{ContainerPort: 443, Name: "server-https"},
								{ContainerPort: 80, Name: "server-http"},
							},
							LivenessProbe: &v1.Probe{
								Handler: v1.Handler{
									HTTPGet: &v1.HTTPGetAction{
										Path: "/index.html",
										Port: intstr.FromInt(80),
									},
								},
								InitialDelaySeconds: 1,
								TimeoutSeconds:      1,
							},
							ReadinessProbe: &v1.Probe{
								Handler: v1.Handler{
									HTTPGet: &v1.HTTPGetAction{
										Path: "/index.html",
										Port: intstr.FromInt(80),
									},
								},
								InitialDelaySeconds: 1,
								TimeoutSeconds:      1,
							},
						},
					},
				},
			},
		},
	}

	return daemonset
}