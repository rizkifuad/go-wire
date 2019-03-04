package model

type Data struct {
	ResourceID string `gorm:"column:resourceId" json:"resource_id"`
}
