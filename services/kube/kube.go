package kube

import (
  "context"

  "k8s.io/client-go/rest"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  corev1 "k8s.io/api/core/v1"
  appsv1 "k8s.io/api/apps/v1"
)


func BuildConfig(kubeconfig string) (*rest.Config, error) {
  if kubeconfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// kube.GetNodes(client)
func GetNodes(client kubernetes.Interface) (*corev1.NodeList, error) {
  nodes, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
  if err != nil {
    return nil, err
  }
  return nodes, nil
}

// kube.GetDeployments(client, "pe", "component=ops-kube-synthetic")
func GetDeployments(client kubernetes.Interface, ns string, selector string) (*appsv1.DeploymentList, error) {
  deploys, err := client.AppsV1().Deployments(ns).List(context.TODO(), metav1.ListOptions{LabelSelector: selector})
  if err != nil {
    return nil, err
  }
  return deploys, nil
}

// kube.GetPods(client, "pe", "component=ops-kube-synthetic")
func GetPods(client kubernetes.Interface, ns string, selector string) (*corev1.PodList, error) {
  pods, err := client.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{LabelSelector: selector})
  if err != nil {
    return nil, err
  }
  return pods, nil
}
