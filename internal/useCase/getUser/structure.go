package getUser

type User struct {
    Id          int32   `json:"id"`
    FirstName   string  `json:"first_name"`
    LastName    string  `json:"last_name"`
    Username    string  `json:"username"`
    PhoneNumber string  `json:"phone_number"`
    Type        string  `json:"type"`
    Bio         string  `json:"bio"`
    PhotoList   []Photo `json:"photo_list"`
}

type Photo struct {
    Id       string `json:"id"`
    UniqueId string `json:"uniq_id"`
    Link     string `json:"link"`
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
                Remote       struct {
                    Type                 string `json:"@type"`
                    Id                   string `json:"id"`
                    UniqueId             string `json:"unique_id"`
                    IsUploadingActive    bool   `json:"is_uploading_active"`
                    IsUploadingCompleted bool   `json:"is_uploading_completed"`
                    UploadedSize         bool   `json:"uploaded_size"`
                } `json:"remote"`
                Local struct {
                    Type            string `json:"@type"`
                    CanBeDownloaded bool   `json:"can_be_downloaded"`
                    CanBeDeleted    bool   `json:"can_be_deleted"`
                    Path            string `json:"path"`
                } `json:"local"`
            } `json:"photo"`
        } `json:"sizes"`
    } `json:"photos"`
}
