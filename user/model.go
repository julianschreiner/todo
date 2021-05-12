package user

import pb "todo"

type User struct {
	Id        uint
	Email     string
	Role      uint
	CreatedAt string
	Dob       string
	Active    bool
	Forename  string
	Surname   string
}

func (u *User) ToPb() *pb.Usr {
	return &pb.Usr{
		Id:       uint64(u.Id),
		Email:    u.Email,
		Forename: u.Forename,
		Surname:  u.Surname,
		Dob:      u.Dob,
	}
}
