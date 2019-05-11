package news_info_extend

import (
	"ara-news/components/mysql"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id        int64 `json:"id,omitempty"`
	Nid       int64 `json:"nid,omitempty"`
	ViewCount int64 `json:"view_count"`
}

var model Model

func init() {
	orm.RegisterModel(new(Model))
}

func (m *Model) TableName() string {
	return "news_info_extend"
}

func InitQuerySetter(genre ...string) orm.QuerySeter {
	var alias string
	if len(genre) > 0 && genre[0] != "" {
		alias = genre[0]
	}

	return mysql.GetQuerySetter(&Model{}, alias)
}

func NewModel(nid int64) Model {
	model.Nid = nid

	return model
}

func FindByNid(nid int64, cols ...string) (Model, error) {
	qs := InitQuerySetter()
	err := qs.Filter("nid", nid).One(&model, cols...)

	return model, err
}
