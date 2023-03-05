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

var mockProject = model.Project{
	Base: model.Base{
		ID: 1,
	},
	Name:             "app 1",
	Description:      "app 1 description",
	URL:              "http://app-1.example.com",
	GithubRepository: "https://github.com/user/repo-1",
}

func TestProjectService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mocks.NewMockProjectRepositoryInterface(ctrl)
	s := NewProjectService(mr)

	newProject := model.Project{
		Name:             mockProject.Name,
		Description:      mockProject.Description,
		URL:              mockProject.URL,
		GithubRepository: mockProject.GithubRepository,
	}

	t.Run(constant.TEST_CREATED, func(t *testing.T) {
		mr.EXPECT().UFindByName(newProject.Name).Return(model.Project{}, gorm.ErrRecordNotFound)
		mr.EXPECT().Create(newProject).Return(mockProject, nil)
		project, err := s.Create(newProject)
		newProject.ID = project.ID

		require.NoError(t, err)
		require.Equal(t, newProject, project)
	})

	t.Run(constant.TEST_ALREADY_EXISTS, func(t *testing.T) {
		mr.EXPECT().UFindByName(newProject.Name).Return(newProject, nil)
		_, err := s.Create(newProject)

		require.EqualError(t, err, errors.New(constant.ALREADY_EXISTS).Error())
	})
}

func TestProjectService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mocks.NewMockProjectRepositoryInterface(ctrl)
	s := NewProjectService(mr)

	updatedProject := mockProject
	updatedProject.Name = "new app"

	t.Run(constant.TEST_UPDATED, func(t *testing.T) {
		mr.EXPECT().Find(mockProject.ID).Return(mockProject, nil)
		mr.EXPECT().UFindByName(updatedProject.Name).Return(model.Project{}, gorm.ErrRecordNotFound)
		mr.EXPECT().Update(mockProject, updatedProject).Return(updatedProject, nil)
		project, err := s.Update(mockProject.ID, updatedProject)

		require.NoError(t, err)
		require.Equal(t, updatedProject, project)
	})

	t.Run(constant.TEST_ALREADY_EXISTS, func(t *testing.T) {
		mr.EXPECT().Find(mockProject.ID).Return(mockProject, nil)
		mr.EXPECT().UFindByName(updatedProject.Name).Return(updatedProject, nil)
		_, err := s.Update(mockProject.ID, updatedProject)

		require.EqualError(t, err, fmt.Errorf(constant.RECORD_ALREADY_EXISTS, updatedProject.Name).Error())
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mr.EXPECT().Find(mockProject.ID).Return(model.Project{}, gorm.ErrRecordNotFound)
		_, err := s.Update(mockProject.ID, updatedProject)

		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}

func TestProjectService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mocks.NewMockProjectRepositoryInterface(ctrl)
	s := NewProjectService(mr)

	t.Run(constant.TEST_DELETED, func(t *testing.T) {
		mr.EXPECT().Find(mockProject.ID).Return(mockProject, nil)
		mr.EXPECT().Delete(mockProject).Return(nil)
		err := s.Delete(mockProject.ID)

		require.NoError(t, err)
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mr.EXPECT().Find(mockProject.ID).Return(model.Project{}, gorm.ErrRecordNotFound)
		err := s.Delete(mockProject.ID)

		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}
