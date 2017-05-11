package createpod

import (
	"fmt"
	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/pkg/api/v1"
	v1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	//"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/kubernetes/pkg/api"
	//	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/client-go/pkg/util/intstr"
)

func createServer(client *kubernetes.Clientset) *v1beta1.DaemonSet {
	daemonSet := generateServerConfig()
	daemonSetObject, err := client.ExtensionsV1beta1().DaemonSets().Create(daemonSet)
	if err != nil {
		panic(err.Error())
	}
	return daemonSet
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
					Containers: []v1.Container{
						{
							Name:            "server",
							Image:           "docker/nginx",
							ImagePullPolicy: "Always",
							Ports: []v1.ContainerPort{
								{ContainerPort: 443, Name: "server-https"},
								{ContainerPort: 80, Name: "server-http"},
							},
							LivenessProbe: &v1.Probe{
								Handler: v1.Handler{
									HTTPGet: &v1.HTTPGetAction{
										Path: "/liveness",
										Port: intstr.FromInt(80),
									},
								},
								InitialDelaySeconds: 1,
								TimeoutSeconds:      1,
							},
							ReadinessProbe: &v1.Probe{
								Handler: v1.Handler{
									HTTPGet: &v1.HTTPGetAction{
										Path: "/readiness",
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
