package model

import (
	"github.com/jinzhu/gorm"
	//"instance/controllers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Good struct {
	Behaviour   string
	Persinality string
}

type Connect interface {
	Create(value interface{})
	Save(value interface{})
	Find(dest interface{}, conds interface{})
}
type Hello struct {
	db *gorm.DB
}

func Getinstance() *Hello {
	d, err := gorm.Open("mysql", "root:Bhavana@1998@tcp(127.0.0.1:3306)/code?charset=utf8&parseTime=True")
	if err != nil {

	}
	return &Hello{db: d}
}
func (p *Hello) Create(value interface{}) {
	p.db.Create(value)
}
func (p *Hello) Save(value interface{}) {

	p.db.Save(value)
}
func (p *Hello) Find(dest interface{}, conds interface{}) {
	p.db.Find(dest, conds)
}
