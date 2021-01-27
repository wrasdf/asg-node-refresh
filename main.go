package main

import (
  "fmt"
  "flag"
  "path/filepath"

  "k8s.io/client-go/util/homedir"
  kube "github.com/wrasdf/asg-node-roller/services/kube"
  aws "github.com/wrasdf/asg-node-roller/services/aws"

  "github.com/aws/aws-sdk-go-v2/service/ec2"

  "k8s.io/klog/v2"
)


type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64  `json:"ref"`
    private string // An unexported field is not encoded.
    Created time.Time
}


func main() {

  klog.InitFlags(nil)

  var kubeconfig string
  var ttlHours string
  var region string

  flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) Absolute path to the kubeconfig file")
  flag.StringVar(&ttlHours, "ttlHours", "48", "TTL time for node")
  flag.StringVar(&region, "region", "ap-southeast-2", "AWS Region")


  // Step1: Get node, which is ruuning time longger than 48 hours in cluster
  kubeClient, _ := kube.KubeClient(kubeconfig)
  kubeNodes, _ := kube.GetNodes(kubeClient)

  for _, node := range kubeNodes.Items {
    if kube.IsLongerThanTTL(node, ttlHours) {
      // Step2: Find ASG by node name
      fmt.Printf("Node: %s \n", node.Name)
    }
  }

  ec2Client, _ := aws.EC2Client(region)
  results, _ := aws.DescribeInstances(ec2Client, &ec2.DescribeInstancesInput{})

  fmt.Println(results)

  // Step3: Terminate node via ASG

}
