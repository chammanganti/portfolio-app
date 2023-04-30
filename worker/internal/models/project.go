package model

// Project model
type Project struct {
	ProjectId        string `gorm:"primaryKey" json:"project_id"`
	Name             string `gorm:"unique;not null" json:"name"`
	Description      string `json:"description"`
	URL              string `json:"url"`
	GithubRepository string `json:"github_repository"`
}

// Project statuses model
type Projects struct {
	Projects []Project `json:"projects"`
}

// Project status model
type ProjectStatus struct {
	ProjectStatusId string `gorm:"primaryKey" json:"project_status_id"`
	Name            string `json:"name"`
	IsHealthy       bool   `gorm:"type:boolean;default:false" json:"is_healthy"`
	ProjectId       string `json:"project_id"`
}

// Project statuses model
type ProjectStatuses struct {
	ProjectStatuses []ProjectStatus `json:"project_statuses"`
}
