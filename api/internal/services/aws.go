package service

import (
	model "api/internal/models"
	"context"

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

	return instances, nil
}

func collectEC2Tags(tags []types.Tag) []model.EC2Tag {
	var ec2Tags []model.EC2Tag
	for _, tag := range tags {
		ec2Tags = append(ec2Tags, model.EC2Tag{
			Name:  *tag.Key,
			Value: *tag.Value,
		})
	}
	return ec2Tags
}
