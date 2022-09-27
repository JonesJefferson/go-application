package userschema

import (
	"github.com/tutorial/api/proto/_generated/proto/user"
)

type User struct {
	Name         string `bson:"name,omitempty"`
	Age          int32  `bson:"age,omitempty"`
	Gender       string `bson:"gender,omitempty"`
	MobileNumber int32  `bson:"mobile_number,omitempty"`
	EmailId      string `bson:"email_id,omitempty"`
}

// ConvertToSchema converts proto struct to User.
func (w *User) ConvertToSchema(user *user.User) {
	w.Name = user.GetName()
	w.Age = user.GetAge()
	w.Gender = user.GetGender()
	w.MobileNumber = user.GetMobileNumber()
	w.EmailId = user.GetEmailId()
}

// ConvertToProto converts User struct into proto
func (u *User) ConvertToProto() *user.User {
	return &user.User{
		Name:         u.Name,
		Age:          u.Age,
		Gender:       u.Gender,
		MobileNumber: u.MobileNumber,
		EmailId:      u.EmailId,
	}
}
