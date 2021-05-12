package aws

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2Client struct {
  client  ec2.Client
  region  string
}

func NewEC2Client(region string) (*EC2Client, error) {

  cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
  if err != nil {
    return nil, err
  }

  return &EC2Client{
    client: *ec2.NewFromConfig(cfg),
    region: region,
  }, nil

}

func (c *EC2Client) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
  results, err := c.client.DescribeInstances(context.TODO(), input)
  if err != nil {
    return nil, err
  }
  return results, nil
}
