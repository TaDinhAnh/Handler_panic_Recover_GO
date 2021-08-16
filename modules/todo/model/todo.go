package todomodel

const EntityName = "Todo"

func (Todo) TableName() string {
	return "todo"
}

type Todo struct {
	Id     int    `json:"id" gorm:"column:id"`
	Title  string `json:"title" gorm:"column:title"`
	Status int    `json:"status" gorm:"column:status"`
	Userid int    `json:"userid" gorm:"column:userid"`
}
