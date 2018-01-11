package services

import (
	"blizzard/repositories"
)

type UserServiceIf interface {
	Create(name string, password string) (int64, error)
	GetByID(id int64) (repositories.UsrTable, error)
	GetByUsernameAndPassword(username, userPassword string) (repositories.UsrTable, error)
}

type UserService struct {
	repo repositories.DbManagerIf
}

// TODO：创建数据库管理类
func NewUserService(repo repositories.DbManagerIf) UserServiceIf {
	return &UserService{repo: repo}
}

func (c *UserService) Create(name string, password string) (int64, error) {
	//fmt.Println("c:", c)
	//fmt.Println("c.repo:", c.repo)

	id, err := c.repo.CreateUsrTable(name, password)
	return id, err
}
func (c *UserService) GetByID(id int64) (repositories.UsrTable, error) {
	return c.repo.GetUsrtableByID(id)
}
func (c *UserService) GetByUsernameAndPassword(username, userPassword string) (repositories.UsrTable, error) {
	return c.repo.GetByUsernameAndPassword(username, userPassword)
}
