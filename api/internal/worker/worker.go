package worker

import (
	repository "api/internal/repositories"
	service "api/internal/services"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Worker interface
type WorkerInterface interface {
	UpdateEC2Statuses()
}

// Worker
type Worker struct {
	Config                  aws.Config
	ProjectStatusRepository repository.ProjectStatusRepositoryInterface
	AWSService              service.AWSServiceInterface
	ProjectStatusService    service.ProjectStatusServiceInterface
}

// New worker
func NewWorker(config aws.Config, db *gorm.DB) WorkerInterface {
	projectStatusRepository := repository.NewProjectStatusRepository(db)
	projectStatusService := service.NewProjectStatusService(projectStatusRepository)
	awsService := service.NewAWSService()
	return &Worker{
		Config:                  config,
		ProjectStatusRepository: projectStatusRepository,
		AWSService:              awsService,
		ProjectStatusService:    projectStatusService,
	}
}

// Updates the ec2 statuses of the projects
func (w Worker) UpdateEC2Statuses() {
	client := ec2.NewFromConfig(w.Config)
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
	instances, err := w.AWSService.GetEC2Status(client, context.TODO(), params)
	if err != nil {
		log.Warn(err)
	}

	fmt.Println(instances)
}
