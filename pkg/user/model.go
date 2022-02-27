package user

type Model struct {
	ID      string `gorm:"primarykey"`
	Name    string
	Surname string
}
