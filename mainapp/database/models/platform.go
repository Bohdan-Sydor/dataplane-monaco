package models

import "time"

func (Platform) IsEntity() {}

func (Platform) TableName() string {
	return "platform"
}

type Platform struct {
	ID           string     `gorm:"PRIMARY_KEY;type:varchar(48);" json:"id"`
	BusinessName string     `json:"business_name"`
	Timezone     string     `json:"timezone"`
	Complete     bool       `json:"complete"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

func (Environment) IsEntity() {}

func (Environment) TableName() string {
	return "environment"
}

type Environment struct {
	ID         string     `gorm:"PRIMARY_KEY;type:varchar(48);" json:"id"`
	Name       string     `gorm:"index:idx_env_name,unique;" json:"name"`
	PlatformID string     `json:"platform_id"`
	Active     bool       `json:"active"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

func (EnvironmentUser) IsEntity() {}

func (EnvironmentUser) TableName() string {
	return "environment_user"
}

type EnvironmentUser struct {
	EnvironmentID string     `gorm:"PRIMARY_KEY;type:varchar(48);" json:"environment_id"`
	UserID        string     `gorm:"PRIMARY_KEY;type:varchar(48);" json:"user_id"`
	Active        bool       `json:"active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}
