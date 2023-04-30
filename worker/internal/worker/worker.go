package worker

import (
	"worker/internal/libs/constant"
	model "worker/internal/models"
	"worker/internal/proto/project"
	"worker/internal/proto/project_status"
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
	"google.golang.org/grpc"
)

// Worker interface
type WorkerInterface interface {
	CacheProjects() string
	CacheProjectStatuses() string
	UpdateEC2Statuses()
}

// Worker
type Worker struct {
	Config               aws.Config
	GRPC                 *grpc.ClientConn
	Redis                store.RedisInterface
	ProjectService       project.ProjectClient
	ProjectStatusService project_status.ProjectStatusClient
	AWSService           service.AWSServiceInterface
}

// New worker
func NewWorker(config aws.Config, grpc *grpc.ClientConn, rdb store.RedisInterface) WorkerInterface {
	projectService := project.NewProjectClient(grpc)
	projectStatusService := project_status.NewProjectStatusClient(grpc)
	awsService := service.NewAWSService()
	return &Worker{
		Config:               config,
		GRPC:                 grpc,
		Redis:                rdb,
		ProjectService:       projectService,
		ProjectStatusService: projectStatusService,
		AWSService:           awsService,
	}
}

// Caches all projects
func (w Worker) CacheProjects() string {
	if _, err := w.Redis.Get("projects"); err == nil {
		return ""
	}

	projects, err := w.ProjectService.Find(context.Background(), &project.FindRequest{})
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	p, err := json.Marshal(projects)
	fmt.Println(string(p))
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

	projectStatuses, err := w.ProjectStatusService.Find(context.Background(), &project_status.FindRequest{})
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

	var projects model.Projects
	if err := json.Unmarshal([]byte(projectsStr), &projects); err != nil {
		log.Warn(err)
		return
	}

	var projectStatuses model.ProjectStatuses
	if err := json.Unmarshal([]byte(projectStatusesStr), &projectStatuses); err != nil {
		log.Warn(err)
		return
	}

	for _, instance := range instances {
		project := getProject(instance.Tags["name"], projects.Projects)
		projectStatus := getProjectStatus(instance.Tags["sid"], project.ProjectId, projectStatuses.ProjectStatuses)
		isHealthy := types.InstanceStateNameRunning == instance.State

		cachedHealthKey := fmt.Sprintf("%s:%s:%s", constant.PROJECT_STATUS_KEY_PREFIX, projectStatus.ProjectStatusId, projectStatus.Name)

		cachedHealth, _ := w.Redis.Get(cachedHealthKey)
		if cachedHealth == strconv.FormatBool(isHealthy) {
			continue
		}
		updateRequest := &project_status.UpdateRequest{
			ProjectStatusId: projectStatus.ProjectStatusId,
			Name:            projectStatus.Name,
			IsHealthy:       isHealthy,
			ProjectId:       projectStatus.ProjectId,
		}
		updatedProjectStatus, err := w.ProjectStatusService.Update(context.Background(), updateRequest)
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
