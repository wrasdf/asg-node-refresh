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

type KubeClient struct {
	client    kubernetes.Interface
	namespace string
}

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

func NewKubeClient(kubeconfig string) (*KubeClient, error) {

  config, err:= buildConfig(kubeconfig)
  if err != nil {
    return nil, err
  }

  return &KubeClient{
    client: kubernetes.NewForConfigOrDie(config),
    namespace: "default",
  }, nil
}

func (c *KubeClient) SetNamespace(ns string) {
	if len(ns) > 0 {
		c.namespace = ns
	}
}

func (c *KubeClient) GetNodes() (*corev1.NodeList, error) {
  nodes, err := c.client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
  if err != nil {
    return nil, err
  }
  return nodes, nil
}

func (c *KubeClient) GetDeployments(selector string) (*appsv1.DeploymentList, error) {
  deploys, err := c.client.AppsV1().Deployments(c.namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: selector})
  if err != nil {
    return nil, err
  }
  return deploys, nil
}

func (c *KubeClient) GetPods(selector string) (*corev1.PodList, error) {
  pods, err := c.client.CoreV1().Pods(c.namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: selector})
  if err != nil {
    return nil, err
  }
  return pods, nil
}


// kube.IsLongerThanTTL(node, "48")
func IsLongerThanTTL(node corev1.Node, ttlHours string) bool {
  utc, _ := time.LoadLocation("UTC")
  nodeTimestamp, _ := ptypes.TimestampProto(node.CreationTimestamp.Time.In(utc))
  nowTimestamp := time.Now().In(utc).Unix()
  ttl, _ := utils.StringToInt64(ttlHours)
  if ((nowTimestamp - nodeTimestamp.GetSeconds())/3600 - ttl) >= 0 {
    return true
  }
  return false
}
