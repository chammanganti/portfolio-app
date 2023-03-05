package model

// Project model
type Project struct {
	Base
	Name             string `gorm:"unique;not null" json:"name"`
	Description      string `json:"description"`
	URL              string `json:"url"`
	GithubRepository string `json:"github_repository"`
}

// Project request model
type ProjectRequest struct {
	Name             string `validate:"required,min=4" json:"name"`
	Description      string `validate:"required,min=8" json:"description"`
	URL              string `validate:"required,url" json:"url"`
	GithubRepository string `validate:"required,url" json:"github_repository"`
}

// Project status model
type ProjectStatus struct {
	Base
	Name      string `json:"name"`
	IsHealthy bool   `gorm:"type:boolean;default:false" json:"is_healthy"`
	ProjectID uint   `json:"project_id"`
}

// Project status request model
type ProjectStatusRequest struct {
	Name      string `validate:"required,min=3" json:"name"`
	IsHealthy bool   `validate:"boolean" json:"is_healthy"`
	ProjectID uint   `validate:"required,numeric" json:"project_id"`
}
