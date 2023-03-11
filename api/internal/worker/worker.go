package worker

import (
	repository "api/internal/repositories"
	service "api/internal/services"
	store "api/internal/store/redis"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Worker interface
type WorkerInterface interface {
	CacheProjects() string
	CacheProjectStatuses() string
	UpdateEC2Statuses()
}

// Worker
type Worker struct {
	Config                  aws.Config
	Redis                   store.RedisInterface
	ProjectRepository       repository.ProjectRepositoryInterface
	ProjectStatusRepository repository.ProjectStatusRepositoryInterface
	AWSService              service.AWSServiceInterface
	ProjectStatusService    service.ProjectStatusServiceInterface
}

// New worker
func NewWorker(config aws.Config, db *gorm.DB, rdb store.RedisInterface) WorkerInterface {
	projectRepository := repository.NewProjectRepository(db)
	projectStatusRepository := repository.NewProjectStatusRepository(db)
	projectStatusService := service.NewProjectStatusService(projectStatusRepository)
	awsService := service.NewAWSService()
	return &Worker{
		Config:                  config,
		Redis:                   rdb,
		ProjectRepository:       projectRepository,
		ProjectStatusRepository: projectStatusRepository,
		AWSService:              awsService,
		ProjectStatusService:    projectStatusService,
	}
}

// Caches all projects
func (w Worker) CacheProjects() string {
	if _, err := w.Redis.Get("projects"); err == nil {
		return ""
	}

	projects, err := w.ProjectRepository.FindAll()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	p, err := json.Marshal(projects)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	w.Redis.Set("projects", string(p), time.Second*28)

	return "projects have been cached"
}

// Caches all project statuses
func (w Worker) CacheProjectStatuses() string {
	if _, err := w.Redis.Get("project_statuses"); err == nil {
		return ""
	}

	projectStatuses, err := w.ProjectStatusRepository.FindAll()
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	p, err := json.Marshal(projectStatuses)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	w.Redis.Set("project_statuses", string(p), time.Second*28)

	return "project statuses have been cached"
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
