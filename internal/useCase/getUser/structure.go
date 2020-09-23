package getUser

type User struct {
	Id          int32    `json:"id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Username    string   `json:"username"`
	PhoneNumber string   `json:"phone_number"`
	Type        string   `json:"type"`
	PhotoList   []string `json:"photo_list"`
}

type ResponseUser struct {
	Id          int32  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Type        struct {
		Type string `json:"@type"`
	} `json:"type"`
}

type ResponsePhotos struct {
	TotalCount int32 `json:"total_count"`
	PhotoList  []struct {
		Type      string `json:"@type"`
		AddedDate int32  `json:"added_date"`
		Id        int64  `json:"id"`
		Sizes     []struct {
			Height int32 `json:"height"`
			Width  int32 `json:"width"`
			Photo  struct {
				Id           int32 `json:"id"`
				ExpectedSize int32 `json:"expected_size"`
				Size         int32 `json:"size"`
				Remote       map[string]struct {
					Id                   string `json:"id"`
					UniqueId             string `json:"unique_id"`
					IsUploadingActive    bool   `json:"is_uploading_active"`
					IsUploadingCompleted bool   `json:"is_uploading_completed"`
					UploadedSize         bool   `json:"uploaded_size"`
				}
			}
		}
	} `json:"photos"`
}
