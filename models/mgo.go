package models

import (
	"StudentManager/database"
	"gopkg.in/mgo.v2/bson"
)

// Insert 添加一条数据
func (s *Student) Insert() error {
	a, c := database.Connect("student")
	defer a.Close()
	s.Uid = bson.NewObjectId().Hex()
	return c.Insert(&s)
}

// All 获取全部数据
func (s *Student) All() ([]Student, error) {
	a, c := database.Connect("student")
	defer a.Close()
	var group []Student
	err := c.Find(nil).All(&group)
	return group, err
}

// Get 通过ID获取数据
func (s *Student) Get(id string) error {
	a, c := database.Connect("student")
	defer a.Close()
	return c.Find(bson.M{"id": id}).One(&s)
}

// Delete 删除一天记录
func (s *Student) Delete() error {
	a, c := database.Connect("student")
	defer a.Close()
	return c.Remove(bson.M{"id": s.Uid})
}

// Update 更新一条数据
func (s *Student) Update() error {
	a, c := database.Connect("goods")
	defer a.Close()
	c.Update(bson.M{"id": s.Uid}, s)
	return s.Get(s.Uid)
}
