package domain

type User struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" bson:"firstname" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" bson:"lastname" gorm:"type:varchar(32)"`
	Age       uint8  `json:"age" bson:"age"`
	Email     string `json:"email" bson:"email" gorm:"type:varchar(256);UNIQUE"`
}

type UserRepository interface {
	Save(user User)
	Update(user User)
	Delete(user User)
	Find(id uint64) User
	FindAll() []User
}
