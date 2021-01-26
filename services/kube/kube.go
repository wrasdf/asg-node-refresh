package kube

import (
  "context"
  "time"

  "k8s.io/client-go/rest"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
  "github.com/golang/protobuf/ptypes"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  corev1 "k8s.io/api/core/v1"
  appsv1 "k8s.io/api/apps/v1"

  utils "github.com/wrasdf/asg-node-roller/services/utils"
)


func buildConfig(kubeconfig string) (*rest.Config, error) {
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

func KubeClient(kubeconfig string) (*kubernetes.Clientset, error){
  config, err:= buildConfig(kubeconfig)
  if err != nil {
    return nil, err
  }

  kubeClient := kubernetes.NewForConfigOrDie(config)
  return kubeClient, nil
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

// kube.IsLongerThanTTL(node, "48")
func IsLongerThanTTL(node corev1.Node, ttlHours string) bool {
  nodeTimestamp, _ := ptypes.TimestampProto(node.CreationTimestamp.Time)
  nowTimestamp := time.Now().Unix()
  ttl, _ := utils.StringToInt64(ttlHours)
  if ((nowTimestamp - nodeTimestamp.GetSeconds())/3600 - ttl) > 0 {
    return true
  }
  return false
}
