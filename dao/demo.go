package dao

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"time"
)

type Area struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	AreaName  string    `json:"area_name" gorm:"column:area_name" description:"区域名称"`
	CityId    int       `json:"city_id" gorm:"column:city_id" description:"城市id"`
	UserId    int64     `json:"user_id" gorm:"column:user_id" description:"操作人"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

func (t *Area) TableName() string {
	return "area"
}

func (t *Area) Find(c *gin.Context, tx *gorm.DB, id string) (*Area, error) {
	area:=&Area{}
	err := tx.WithContext(c).Where("id = ?", id).Find(area).Error
	if err != nil {
		return nil, err
	}
	return area, nil
}
