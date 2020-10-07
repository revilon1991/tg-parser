package getStorageMemberList

type ResponseMember struct {
	Id          int      `json:"id"`
	UserId      int      `json:"user_id"`
	Username    string   `json:"username"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	PhoneNumber string   `json:"phone_number"`
	Type        string   `json:"type"`
	Bio         string   `json:"bio"`
	UpdatedAt   string   `json:"updated_at"`
	JoinDate    string   `json:"join_date"`
	PhotoList   []string `json:"photo_list"`
}

type MemberPhoto struct {
	MemberId int
	Link     string
}
