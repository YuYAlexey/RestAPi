package connect

type Todoinfo struct {
	Id    int    `json:"id" gorm:"primary_key"`
	State string `json:"state"`
	Date  string `json:"date"`
	Name  string `json:"Name"`
}
