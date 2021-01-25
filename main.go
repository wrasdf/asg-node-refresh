package main
import (
  "fmt"
  "flag"
  "path/filepath"

  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/util/homedir"
  utils "github.com/wrasdf/asg-node-roller/utils"

  "k8s.io/klog/v2"
)


func main() {

  klog.InitFlags(nil)

  var kubeconfig string
  var ttlTime string
  var interval string

  flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) Absolute path to the kubeconfig file")
  flag.StringVar(&ttlTime, "ttlTime", "48h", "TTL time for node")
  flag.StringVar(&interval, "interval", "30m", "Interval time for each check")

  config, err:= utils.BuildConfig(kubeconfig)
  if err != nil {
		klog.Fatal(err)
	}

  client := kubernetes.NewForConfigOrDie(config)
  nodes, _ := utils.GetNodes(client)
  deploys, _ := utils.GetDeployments(client, "kube-system")

  for _, node := range nodes.Items {
    fmt.Printf("Node: %s\n", node.Name)
  }

  for _, deploy := range deploys.Items {
    fmt.Printf("deploy: %s\n", deploy.Name)
  }

  // Step1: Get node, which is ruuning time longger than 48 hours in cluster

  // Step2: Find ASG by node name

  // Step3: Terminate node via ASG

}
