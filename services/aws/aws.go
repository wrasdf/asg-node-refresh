package aws

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

func EC2Client(region string) (*ec2.Client, error) {
  cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
  if err != nil {
    return nil, err
  }
  client := ec2.NewFromConfig(cfg)
  return client, nil
}

func DescribeInstances(client *ec2.Client, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
  results, err := client.DescribeInstances(context.TODO(), input)
  if err != nil {
    return nil, err
  }
  return results, nil
}
