package worker

import (
	"worker/internal/libs/constant"
	model "worker/internal/models"
	repository "worker/internal/repositories"
	service "worker/internal/services"
	store "worker/internal/store/redis"

	"context"
	"encoding/json"
	"fmt"
	"strconv"
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
	w.Redis.Set(constant.PROJECTS_KEY, string(p), time.Second*28)

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
	w.Redis.Set(constant.PROJECT_STATUSES_KEY, string(p), time.Second*28)

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
		return
	}

	projectsStr, _ := w.Redis.Get(constant.PROJECTS_KEY)
	projectStatusesStr, _ := w.Redis.Get(constant.PROJECT_STATUSES_KEY)

	var projects []model.Project
	if err := json.Unmarshal([]byte(projectsStr), &projects); err != nil {
		log.Warn(err)
		return
	}

	var projectStatuses []model.ProjectStatus
	if err := json.Unmarshal([]byte(projectStatusesStr), &projectStatuses); err != nil {
		log.Warn(err)
		return
	}

	for _, instance := range instances {
		project := getProject(instance.Tags["name"], projects)
		projectStatus := getProjectStatus(instance.Tags["sid"], project.ProjectId, projectStatuses)
		isHealthy := types.InstanceStateNameRunning == instance.State
		projectStatus.IsHealthy = isHealthy

		cachedHealthKey := fmt.Sprintf("%s:%s:%s", constant.PROJECT_STATUS_KEY_PREFIX, projectStatus.ProjectStatusId, projectStatus.Name)

		cachedHealth, _ := w.Redis.Get(cachedHealthKey)
		if cachedHealth == strconv.FormatBool(isHealthy) {
			continue
		}
		updatedProjectStatus, err := w.ProjectStatusService.Update(projectStatus.ProjectStatusId, projectStatus)
		if err != nil {
			log.Warn(err)
		}
		w.Redis.Set(cachedHealthKey, strconv.FormatBool(isHealthy), time.Hour)

		log.Info(updatedProjectStatus)
	}

	fmt.Println(instances)
}

func getProject(name string, projects []model.Project) model.Project {
	for _, project := range projects {
		if name == project.Name {
			return project
		}
	}
	return model.Project{}
}

func getProjectStatus(name string, projectId string, projectStatuses []model.ProjectStatus) model.ProjectStatus {
	for _, projectStatus := range projectStatuses {
		if name == projectStatus.Name && projectId == projectStatus.ProjectId {
			return projectStatus
		}
	}
	return model.ProjectStatus{}
}
