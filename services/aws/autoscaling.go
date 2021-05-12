package aws

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

type AutoScalingClient struct {
  client autoscaling.Client
  region string
}

func NewASGClient(region string) (*AutoScalingClient, error) {

  cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
  if err != nil {
    return nil, err
  }

  return &AutoScalingClient{
    client: *autoscaling.NewFromConfig(cfg),
    region: region,
  }, nil

}

func (c *AutoScalingClient) TerminateInstanceInAutoScalingGroup(input *autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error) {
  results, err := c.client.TerminateInstanceInAutoScalingGroup(context.TODO(), input)
  if err != nil {
    return nil, err
  }
  return results, nil
}
