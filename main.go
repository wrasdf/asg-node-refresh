package main

import (
  "fmt"
  "flag"
  "time"
  "path/filepath"

  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/util/homedir"
  kube "github.com/wrasdf/asg-node-roller/services/kube"
  utils "github.com/wrasdf/asg-node-roller/services/utils"
  "github.com/golang/protobuf/ptypes"

  "k8s.io/klog/v2"
)


func main() {

  klog.InitFlags(nil)

  var kubeconfig string
  var ttlHours string

  flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) Absolute path to the kubeconfig file")
  flag.StringVar(&ttlHours, "ttlHours", "48", "TTL time for node")

  config, err:= kube.BuildConfig(kubeconfig)
  if err != nil {
		klog.Fatal(err)
	}

  client := kubernetes.NewForConfigOrDie(config)

  // Step1: Get node, which is ruuning time longger than 48 hours in cluster

  nodes, _ := kube.GetNodes(client)

  for _, node := range nodes.Items {
    nodeTimestamp, _ := ptypes.TimestampProto(node.CreationTimestamp.Time)
    nowTimestamp := time.Now().Unix()
    ttlHours, _ := utils.StringToInt(ttlHours)
    if ((nowTimestamp - nodeTimestamp.GetSeconds())/3600 - ttlHours) > 0 {
      fmt.Printf("Node: %s \n", node.Name)
    }
  }

  // Step2: Find ASG by node name

  // Step3: Terminate node via ASG

}
