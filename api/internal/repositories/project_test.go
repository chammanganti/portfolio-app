package repository

import (
	"api/internal/libs/constant"
	model "api/internal/models"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var mockProjects = []model.Project{
	{
		Base: model.Base{
			ID: 1,
		},
		Name:             "app 1",
		Description:      "app 1 description",
		URL:              "http://app-1.example.com",
		GithubRepository: "https://github.com/user/repo-1",
	},
	{
		Base: model.Base{
			ID: 2,
		},
		Name:             "app 2",
		Description:      "app 2 description",
		URL:              "http://app-2.example.com",
		GithubRepository: "https://github.com/user/repo-2",
	},
	{
		Base: model.Base{
			ID: 3,
		},
		Name:             "app 3",
		Description:      "app 3 description",
		URL:              "http://app-3.example.com",
		GithubRepository: "https://github.com/user/repo-3",
	},
}

var projectRows = []string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "url", "github_repository"}

func TestProjectRepository_FindAll(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectRepository(db)

	rows := sqlmock.NewRows(projectRows)
	for _, p := range mockProjects {
		rows.AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.Description, p.URL, p.GithubRepository)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "projects" WHERE "projects"."deleted_at" IS NULL`)).
		WillReturnRows(rows)

	res, err := r.FindAll()

	require.NoError(t, err)
	require.Equal(t, mockProjects, res)
}

func TestProjectRepository_Find(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectRepository(db)

	p := mockProjects[0]

	t.Run(constant.TEST_FOUND, func(t *testing.T) {
		rows := sqlmock.NewRows(projectRows).AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.Description, p.URL, p.GithubRepository)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "projects" WHERE "projects"."id" = $1 AND "projects"."deleted_at" IS NULL ORDER BY "projects"."id" LIMIT 1`)).
			WithArgs(p.ID).
			WillReturnRows(rows)

		res, err := r.Find(p.ID)

		require.NoError(t, err)
		require.Equal(t, p, res)
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
		_, err := r.Find(1)
		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}

func TestProjectRepository_FindByName(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectRepository(db)

	p := mockProjects[0]

	t.Run(constant.TEST_FOUND, func(t *testing.T) {
		rows := sqlmock.NewRows(projectRows).AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.Description, p.URL, p.GithubRepository)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "projects" WHERE name = $1 AND "projects"."deleted_at" IS NULL ORDER BY "projects"."id" LIMIT 1`)).
			WithArgs(p.Name).
			WillReturnRows(rows)

		res, err := r.FindByName(p.Name)

		require.NoError(t, err)
		require.Equal(t, p, res)
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
		_, err := r.Find(1)
		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}

func TestProjectRepository_UFindByName(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectRepository(db)

	p := mockProjects[0]

	t.Run(constant.TEST_FOUND, func(t *testing.T) {
		rows := sqlmock.NewRows(projectRows).AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.Description, p.URL, p.GithubRepository)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "projects" WHERE name = $1 ORDER BY "projects"."id" LIMIT 1`)).
			WithArgs(p.Name).
			WillReturnRows(rows)

		res, err := r.UFindByName(p.Name)

		require.NoError(t, err)
		require.Equal(t, p, res)
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
		_, err := r.Find(1)
		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}

func TestProjectRepository_Create(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectRepository(db)

	p := mockProjects[0]

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "projects" ("created_at","updated_at","deleted_at","name","description","url","github_repository") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), p.Name, p.Description, p.URL, p.GithubRepository).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(p.ID))
	mock.ExpectCommit()

	res, err := r.Create(model.Project{
		Name:             p.Name,
		Description:      p.Description,
		URL:              p.URL,
		GithubRepository: p.GithubRepository,
	})

	p.CreatedAt = res.CreatedAt
	p.UpdatedAt = res.UpdatedAt

	require.NoError(t, err)
	require.Equal(t, p, res)
}

func TestProjectRepository_Update(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectRepository(db)

	p := mockProjects[0]

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "projects" SET "updated_at"=$1,"name"=$2,"description"=$3,"url"=$4,"github_repository"=$5 WHERE "projects"."deleted_at" IS NULL AND "id" = $6`)).
		WithArgs(sqlmock.AnyArg(), p.Name, p.Description, p.URL, p.GithubRepository, p.ID).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	res, err := r.Update(p, p)

	p.UpdatedAt = res.UpdatedAt

	require.NoError(t, err)
	require.Equal(t, p, res)
}

func TestProjectRepository_Delete(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectRepository(db)

	p := mockProjects[0]

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "projects" SET "deleted_at"=$1 WHERE "projects"."id" = $2 AND "projects"."deleted_at" IS NULL`)).
		WithArgs(sqlmock.AnyArg(), p.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := r.Delete(p)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}
