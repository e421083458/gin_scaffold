package dao

import (
	"github.com/e421083458/gin_scaffold/dto"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"time"
)

type ListPageOutput struct {
	List  []User `form:"list" json:"list" comment:"用户列表" validate:""`
	Total int64  `form:"page" json:"page" comment:"用户总数" validate:"required"`
}

type User struct {
	Id        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Addr      string    `json:"addr" gorm:"column:addr"`
	Age       int       `json:"age" gorm:"column:age"`
	Birth     string    `json:"birth" gorm:"column:birth"`
	Sex       int       `json:"sex" gorm:"column:sex"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at"" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
}

func (f *User) TableName() string {
	return "user"
}

func (f *User) Del(c *gin.Context, tx *gorm.DB, idSlice []string) error {
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where("id in (?)", idSlice).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *User) Find(c *gin.Context, tx *gorm.DB, id int64) (*User, error) {
	var user *User
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (f *User) PageList(c *gin.Context, tx *gorm.DB, params *dto.ListPageInput) ([]User, int64, error) {
	var list []User
	var count int64
	offset := (params.Page - 1) * params.PageSize
	query := tx.SetCtx(public.GetGinTraceContext(c))
	if params.Name != "" {
		query = query.Where("name = ?", params.Name)
	}
	err := query.Limit(params.PageSize).Offset(offset).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Table("user").Count(&count).Error
	if errCount != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (f *User) Save(c *gin.Context, tx *gorm.DB) error {
	if err := tx.SetCtx(public.GetGinTraceContext(c)).Save(f).Error; err != nil {
		return err
	}
	return nil
}
