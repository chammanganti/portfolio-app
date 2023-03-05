package handler

import (
	service "api/internal/services"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/gofiber/fiber/v2"
)

// AWS handler interface
type AWSHandlerInterface interface {
	GetEC2Status(c *fiber.Ctx) error
}

// AWS handler
type AWSHandler struct {
	Config  aws.Config
	Service service.AWSServiceInterface
}

// New AWS handler
func NewAWSHandler(config aws.Config, service service.AWSServiceInterface) AWSHandlerInterface {
	return &AWSHandler{
		Config:  config,
		Service: service,
	}
}

// Gets the ec2 instances' status
func (a AWSHandler) GetEC2Status(c *fiber.Ctx) error {
	params := &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []string{
					string(types.InstanceStateNamePending),
					string(types.InstanceStateNameRunning),
					string(types.InstanceStateNameStopped),
					string(types.InstanceStateNameTerminated),
				},
			},
			{
				Name:   aws.String("tag:Project"),
				Values: []string{*aws.String("portfolio")},
			},
		},
	}
	client := ec2.NewFromConfig(a.Config)
	instances, err := a.Service.GetEC2Status(client, context.TODO(), params)
	if err != nil {
		return RespondWithError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(instances)
}
