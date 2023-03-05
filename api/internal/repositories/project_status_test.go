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

var mockProjectStatuses = []model.ProjectStatus{
	{
		Base: model.Base{
			ID: 1,
		},
		Name:      "ec2",
		IsHealthy: true,
		ProjectID: 1,
	},
	{
		Base: model.Base{
			ID: 2,
		},
		Name:      "app",
		IsHealthy: true,
		ProjectID: 1,
	},
	{
		Base: model.Base{
			ID: 3,
		},
		Name:      "eks",
		IsHealthy: true,
		ProjectID: 2,
	},
}

var projectStatusRows = []string{"id", "created_at", "updated_at", "deleted_at", "name", "is_healthy", "project_id"}

func TestProjectStatusRepository_FindAll(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectStatusRepository(db)

	rows := sqlmock.NewRows(projectStatusRows)
	for _, p := range mockProjectStatuses {
		rows.AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.IsHealthy, p.ProjectID)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project_statuses" WHERE "project_statuses"."deleted_at" IS NULL`)).
		WillReturnRows(rows)

	res, err := r.FindAll()

	require.NoError(t, err)
	require.Equal(t, mockProjectStatuses, res)
}

func TestProjectStatusRepository_Find(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectStatusRepository(db)

	p := mockProjectStatuses[0]

	t.Run(constant.TEST_FOUND, func(t *testing.T) {
		rows := sqlmock.NewRows(projectStatusRows).AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.IsHealthy, p.ProjectID)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project_statuses" WHERE "project_statuses"."id" = $1 AND "project_statuses"."deleted_at" IS NULL ORDER BY "project_statuses"."id" LIMIT 1`)).
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

func TestProjectStatusRepository_FindByName(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectStatusRepository(db)

	p := mockProjectStatuses[0]

	t.Run(constant.TEST_FOUND, func(t *testing.T) {
		rows := sqlmock.NewRows(projectStatusRows).AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.IsHealthy, p.ProjectID)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project_statuses" WHERE name = $1 AND "project_statuses"."deleted_at" IS NULL ORDER BY "project_statuses"."id" LIMIT 1`)).
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

func TestProjectStatusRepository_FindByNameProject(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectStatusRepository(db)

	p := mockProjectStatuses[0]

	t.Run(constant.TEST_FOUND, func(t *testing.T) {
		rows := sqlmock.NewRows(projectStatusRows).AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.IsHealthy, p.ProjectID)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project_statuses" WHERE name = $1 AND project_id = $2 AND "project_statuses"."deleted_at" IS NULL ORDER BY "project_statuses"."id" LIMIT 1`)).
			WithArgs(p.Name, p.ProjectID).
			WillReturnRows(rows)

		res, err := r.FindByNameProject(p.Name, p.ProjectID)

		require.NoError(t, err)
		require.Equal(t, p, res)
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
		_, err := r.Find(1)
		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}

func TestProjectStatusRepository_UFindByNameProject(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectStatusRepository(db)

	p := mockProjectStatuses[0]

	t.Run(constant.TEST_FOUND, func(t *testing.T) {
		rows := sqlmock.NewRows(projectStatusRows).AddRow(p.ID, p.CreatedAt, p.UpdatedAt, p.DeletedAt, p.Name, p.IsHealthy, p.ProjectID)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "project_statuses" WHERE name = $1 AND project_id = $2 ORDER BY "project_statuses"."id" LIMIT 1`)).
			WithArgs(p.Name, p.ProjectID).
			WillReturnRows(rows)

		res, err := r.UFindByNameProject(p.Name, p.ProjectID)

		require.NoError(t, err)
		require.Equal(t, p, res)
	})

	t.Run(constant.TEST_NOT_FOUND, func(t *testing.T) {
		mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
		_, err := r.Find(1)
		require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	})
}

func TestProjectStatusRepository_Create(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectStatusRepository(db)

	p := mockProjectStatuses[0]

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "project_statuses" ("created_at","updated_at","deleted_at","name","is_healthy","project_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), p.Name, p.IsHealthy, p.ProjectID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(p.ID))
	mock.ExpectCommit()

	res, err := r.Create(model.ProjectStatus{
		Name:      p.Name,
		IsHealthy: p.IsHealthy,
		ProjectID: p.ProjectID,
	})

	p.CreatedAt = res.CreatedAt
	p.UpdatedAt = res.UpdatedAt

	require.NoError(t, err)
	require.Equal(t, p, res)
}

func TestProjectStatusRepository_Update(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectStatusRepository(db)

	p := mockProjectStatuses[0]

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "project_statuses" SET "updated_at"=$1,"name"=$2,"is_healthy"=$3 WHERE "project_statuses"."deleted_at" IS NULL AND "id" = $4`)).
		WithArgs(sqlmock.AnyArg(), p.Name, p.IsHealthy, p.ID).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	res, err := r.Update(p, p)

	p.UpdatedAt = res.UpdatedAt

	require.NoError(t, err)
	require.Equal(t, p, res)
}

func TestProjectStatusRepository_Delete(t *testing.T) {
	db, mockDb, mock := SetUp(t)
	defer mockDb.Close()
	r := NewProjectStatusRepository(db)

	p := mockProjectStatuses[0]

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "project_statuses" SET "deleted_at"=$1 WHERE "project_statuses"."id" = $2 AND "project_statuses"."deleted_at" IS NULL`)).
		WithArgs(sqlmock.AnyArg(), p.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := r.Delete(p)

	require.NoError(t, err)
	require.Equal(t, nil, err)
}
