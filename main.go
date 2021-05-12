package main

import (
  "flag"
  "path/filepath"

  "k8s.io/client-go/util/homedir"
  "github.com/aws/aws-sdk-go-v2/service/autoscaling"
  "k8s.io/klog/v2"

  kube "github.com/wrasdf/asg-node-roller/services/kube"
  aws "github.com/wrasdf/asg-node-roller/services/aws"

)


func main() {

  klog.InitFlags(nil)

  var kubeconfig string
  var ttlHours string
  var region string

  flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) Absolute path to the kubeconfig file")
  flag.StringVar(&ttlHours, "ttlHours", "48", "TTL time for node")
  flag.StringVar(&region, "region", "us-east-1", "AWS Region")

  k8sClient, _ := kube.NewKubeClient(kubeconfig)
  kubeNodes, _ := k8sClient.GetNodes()
  ttlNodes := []string{}

  for _, node := range kubeNodes.Items {
    if kube.IsLongerThanTTL(node, ttlHours) {
      ttlNodes = append(ttlNodes, node.Name)
    }
  }

  if len(ttlNodes) > 0 {
    asgClient, _ := aws.NewASGClient(region)
    // TODO fix this
    asgClient.TerminateInstanceInAutoScalingGroup(&autoscaling.TerminateInstanceInAutoScalingGroupInput{
      InstanceId: (string)ttlNodes[0],
      ShouldDecrementDesiredCapacity: (bool)false,
    })
  }

}
