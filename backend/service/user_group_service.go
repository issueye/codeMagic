package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/issueye/code_magic/backend/common/model"
	"github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/config"
	"github.com/issueye/code_magic/backend/global"
	"github.com/issueye/code_magic/backend/pkg/utils"
	"github.com/issueye/code_magic/backend/repository"
	"gorm.io/gorm"
)

type UserGroup struct {
	service.BaseService
}

func (owner *UserGroup) Self() *UserGroup {
	return owner
}

func (owner *UserGroup) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewUserGroup(args ...service.ServiceContext) *UserGroup {
	return service.NewService(&UserGroup{}, args...)
}

// FindUserGroup
// 查找用户
func (UserGroup *UserGroup) FindUserGroup(info *repository.Login) (*model.UserGroupInfo, error) {
	query := UserGroup.GetDB().Model(&model.UserGroupInfo{}).Order("id")
	query = query.Where("account = ?", info.Account)

	// 判断是否需要验证密码
	r := config.GetParam("SERVER-MODE", "release")
	fmt.Printf("当前运行模式：%s", r.String())
	if strings.EqualFold(r.String(), "release") {
		query = query.Where("password = ?", info.Password)
	}

	data := new(model.UserGroupInfo)
	err := query.Find(data).Error
	return data, err
}

// CreateAdminNonExistent
// 创建管理员用户，如果不存在
func (UserGroup *UserGroup) CreateAdminNonExistent() error {
	info := new(model.UserGroupInfo)
	err := UserGroup.GetDB().Model(&model.UserGroupInfo{}).Where("name = ?", global.AdminGroupName).Where("id = ?", global.AdminGroupId).Find(info).Error
	if err != nil {
		return err
	}

	if info.ID != "" {
		return nil
	}

	info.ID = global.AdminGroupId
	info.Name = global.AdminGroupName
	info.Mark = "系统自动生成的管理员数据"
	info.State = 1
	info.Sys = 1
	return UserGroup.GetDB().Create(info).Error
}

// Create
// 创建用户信息
func (UserGroup *UserGroup) Create(data *repository.CreateUserGroup) error {
	tx := UserGroup.GetDB().Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()
	}()

	info := new(model.UserGroupInfo)
	info.ID = strconv.FormatInt(utils.GenID(), 10)
	info.Name = data.Name
	info.Mark = data.Mark
	info.State = 1
	info.Sys = 0
	err = UserGroup.GetDB().Create(info).Error
	if err != nil {
		return err
	}

	return nil
}

// GetById
// 根据用户ID查找用户信息
func (UserGroup *UserGroup) GetById(id string) (*model.UserGroupInfo, error) {
	info := new(model.UserGroupInfo)
	err := UserGroup.GetDB().Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// Modify
// 修改用户信息
func (UserGroup *UserGroup) Modify(info *repository.ModifyUserGroup) error {
	m := make(map[string]any)
	m["name"] = info.Name
	m["mark"] = info.Mark

	return UserGroup.GetDB().Model(&model.UserGroupInfo{}).Where("id = ?", info.ID).Updates(m).Error
}

// Status
// 修改用户信息
func (UserGroup *UserGroup) Status(info *repository.StatusUserGroup) error {
	return UserGroup.GetDB().
		Model(&model.UserGroupInfo{}).
		Where("id = ?", info.ID).
		Update("state", info.State).
		Error
}

// Delete
// 删除用户信息
func (UserGroup *UserGroup) Delete(id string) error {
	return UserGroup.GetDB().Where("id = ?", id).Delete(&model.UserGroupInfo{}).Error
}

// List
// 获取用户列表
func (UserGroup *UserGroup) List(info *model.Page[*repository.QueryUserGroup]) ([]*model.UserGroupInfo, error) {
	UserGroupInfo := new(model.UserGroupInfo)
	list := make([]*model.UserGroupInfo, 0)

	err := service.Filter(UserGroup, UserGroupInfo.TableName(), info, &list, func(db *gorm.DB) (*gorm.DB, error) {
		query := db.Order("id")

		// 组名称
		if info.Condition.Name != "" {
			query = query.Where("name like ?", fmt.Sprintf("%%%s%%", info.Condition.Name))
		}

		// 备注
		if info.Condition.Mark != "" {
			query = query.Where("mark like ?", fmt.Sprintf("%%%s%%", info.Condition.Name))
		}

		return query, nil
	})

	return list, err
}
