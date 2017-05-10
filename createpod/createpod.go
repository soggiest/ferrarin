package createpod

import (
  v1beta1 "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
  v1 "k8s.io/client-go/pkg/api/v1"
)


func create_server(client *kubernetes.ClientSet) (v1beta1.DaemonSet) {
  daemonManifest := generate_server_config()
  
  daemonSet, err := client.ExtensionsV1beta1().DaemonSet().Create(daemonManifest)

  return daemonSet
}

func generate_server_config() *v1beta1.DaemonSet {
  	labels := map[string]string{"test-pods": "server"} 
	daemonset := &extensions.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: v1.NameSpaceDefault,
			Name:      "test-pods-server",
			Labels:    labels,
		},
		Spec: extensions.DaemonSetSpec{
			Template: api.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: api.PodSpec{
					ServiceAccountName: opts.ServiceAccount,
					Containers: []api.Container{
						{
							Name:            "server",
							Image:           "docker/nginx",
							ImagePullPolicy: "Always",
							Ports: []api.ContainerPort{
								{ContainerPort: 443, Name: "server-https"},
								{ContainerPort: 80, Name: "server-http"}
							},
							LivenessProbe: &api.Probe{
								Handler: api.Handler{
									HTTPGet: &api.HTTPGetAction{
										Path: "/liveness",
										Port: intstr.FromInt(80),
									},
								},
								InitialDelaySeconds: 1,
								TimeoutSeconds:      1,
							},
							ReadinessProbe: &api.Probe{
								Handler: api.Handler{
									HTTPGet: &api.HTTPGetAction{
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

