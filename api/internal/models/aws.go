package model

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// EC2 model
type EC2 struct {
	InstanceID string                  `json:"name"`
	State      types.InstanceStateName `json:"state"`
	Tags       []EC2Tag                `json:"tags"`
}

// EC2 tags model
type EC2Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
