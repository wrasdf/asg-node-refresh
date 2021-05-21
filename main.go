package main

import (
	"flag"
	"os"
	"fmt"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"k8s.io/client-go/util/homedir"

	"k8s.io/klog/v2"

	awsLib "github.com/wrasdf/asg-node-roller/services/aws"
	kube "github.com/wrasdf/asg-node-roller/services/kube"
)

func main() {

	klog.InitFlags(nil)

	var kubeconfig string
	var ttlHours string
	var region string

	flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) Absolute path to the kubeconfig file")
	flag.StringVar(&ttlHours, "ttlHours", os.Getenv("ttlHours"), "TTL time for node")
	flag.StringVar(&region, "region", os.Getenv("region"), "AWS Region")

	flag.Parse()

	k8sClient, _ := kube.NewKubeClient(kubeconfig)
	ec2Client, _ := awsLib.NewEC2Client(region)
	kubeNodes, _ := k8sClient.GetNodes()
	ttlNodes := []string{}

	for _, node := range kubeNodes.Items {
		if kube.IsLongerThanTTL(node, ttlHours) {
			ttlNodes = append(ttlNodes, node.Name)
		}
	}

	if len(ttlNodes) > 0 {
		Node := ttlNodes[0]
		instance, _ := ec2Client.DescribeInstances(&ec2.DescribeInstancesInput{
			Filters: []types.Filter{
				{
					Name:   aws.String("network-interface.private-dns-name"),
					Values: []string{Node},
				},
			}})

		NodeInstanceId := *instance.Reservations[0].Instances[0].InstanceId
		falseBool := false

		asgClient, _ := awsLib.NewASGClient(region)
		_, err := asgClient.TerminateInstanceInAutoScalingGroup(&autoscaling.TerminateInstanceInAutoScalingGroupInput{
			InstanceId:                     &NodeInstanceId,
			ShouldDecrementDesiredCapacity: &falseBool,
		})
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			return
		}
		fmt.Printf("The Node: %s will be drained soon.", Node)
	}
}
