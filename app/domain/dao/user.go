package dao

type User struct {
	ID       		int    `gorm:"column:id; primary_key; not null;autoIncrement" json:"id"`
	Name     		string `gorm:"column:name" json:"name"`
	Designation    	string `gorm:"column:designation" json:"designation"`
	BaseModel
}