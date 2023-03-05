package service

import (
	"api/internal/libs/constant"
	model "api/internal/models"
	"api/internal/repositories/mocks"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var mockProjectStatus = model.ProjectStatus{
	Base: model.Base{
		ID: 1,
	},
	Name:      "ec2",
	IsHealthy: true,
	ProjectID: 1,
}

func TestProjectStatusService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mocks.NewMockProjectStatusRepositoryInterface(ctrl)
	s := NewProjectStatusService(mr)

	newProjectStatus := model.ProjectStatus{
		Name:      mockProjectStatus.Name,
		IsHealthy: mockProjectStatus.IsHealthy,
		ProjectID: mockProjectStatus.ProjectID,
	}

	t.Run(constant.TEST_CREATED, func(t *testing.T) {
		mr.EXPECT().UFindByNameProject(newProjectStatus.Name, newProjectStatus.ProjectID).Return(model.ProjectStatus{}, gorm.ErrRecordNotFound)
		mr.EXPECT().Create(newProjectStatus).Return(mockProjectStatus, nil)
		projectStatus, err := s.Create(newProjectStatus)
		newProjectStatus.ID = projectStatus.ID

		require.NoError(t, err)
		require.Equal(t, newProjectStatus, projectStatus)
	})

	t.Run(constant.TEST_ALREADY_EXISTS, func(t *testing.T) {
		mr.EXPECT().UFindByNameProject(newProjectStatus.Name, newProjectStatus.ProjectID).Return(newProjectStatus, nil)
		_, err := s.Create(newProjectStatus)

		require.EqualError(t, err, errors.New(constant.ALREADY_EXISTS).Error())
	})
}

func TestProjectStatusService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mocks.NewMockProjectStatusRepositoryInterface(ctrl)
	s := NewProjectStatusService(mr)

	updatedProjectStatus := mockProjectStatus
	updatedProjectStatus.Name = "elb"

	t.Run(constant.TEST_UPDATED, func(t *testing.T) {
		mr.EXPECT().Find(mockProjectStatus.ID).Return(mockProjectStatus, nil)
		mr.EXPECT().UFindByNameProject(updatedProjectStatus.Name, mockProjectStatus.ProjectID).Return(model.ProjectStatus{}, gorm.ErrRecordNotFound)
		mr.EXPECT().Update(mockProjectStatus, updatedProjectStatus).Return(updatedProjectStatus, nil)
		projectStatus, err := s.Update(mockProjectStatus.ID, updatedProjectStatus)

		require.NoError(t, err)
		require.Equal(t, updatedProjectStatus, projectStatus)
	})

	t.Run(constant.TEST_ALREADY_EXISTS, func(t *testing.T) {
		mr.EXPECT().Find(mockProjectStatus.ID).Return(mockProjectStatus, nil)
		mr.EXPECT().UFindByNameProject(updatedProjectStatus.Name, mockProjectStatus.ProjectID).Return(updatedProjectStatus, nil)
		_, err := s.Update(mockProjectStatus.ID, updatedProjectStatus)

		require.EqualError(t, err, fmt.Errorf(constant.RECORD_ALREADY_EXISTS, updatedProjectStatus.Name).Error())
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mr.EXPECT().Find(mockProjectStatus.ID).Return(model.ProjectStatus{}, gorm.ErrRecordNotFound)
		_, err := s.Update(mockProjectStatus.ID, updatedProjectStatus)

		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}

func TestProjectStatusService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mocks.NewMockProjectStatusRepositoryInterface(ctrl)
	s := NewProjectStatusService(mr)

	t.Run(constant.TEST_DELETED, func(t *testing.T) {
		mr.EXPECT().Find(mockProjectStatus.ID).Return(mockProjectStatus, nil)
		mr.EXPECT().Delete(mockProjectStatus).Return(nil)
		err := s.Delete(mockProjectStatus.ID)

		require.NoError(t, err)
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mr.EXPECT().Find(mockProjectStatus.ID).Return(model.ProjectStatus{}, gorm.ErrRecordNotFound)
		err := s.Delete(mockProjectStatus.ID)

		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}
