package service

import (
	model "worker/internal/models"

	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// AWS service interface
type AWSServiceInterface interface {
	GetEC2Status(client *ec2.Client, context context.Context, params *ec2.DescribeInstancesInput) ([]model.EC2, error)
}

// AWS service
type AWSService struct{}

// New AWS service
func NewAWSService() AWSServiceInterface {
	return &AWSService{}
}

// Gets the ec2 instances' status
func (a AWSService) GetEC2Status(client *ec2.Client, context context.Context, params *ec2.DescribeInstancesInput) ([]model.EC2, error) {
	res, err := client.DescribeInstances(context, params)
	if err != nil {
		return []model.EC2{}, err
	}

	var instances []model.EC2
	for _, reservation := range res.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, model.EC2{
				InstanceID: *instance.InstanceId,
				State:      instance.State.Name,
				Tags:       collectEC2Tags(instance.Tags),
			})
		}
	}
	if len(instances) == 0 {
		instances = []model.EC2{}
	}

	return instances, nil
}

func collectEC2Tags(tags []types.Tag) map[string]string {
	ec2Tags := make(map[string]string)
	for _, tag := range tags {
		ec2Tags[strings.ToLower(*tag.Key)] = *tag.Value
	}
	return ec2Tags
}
