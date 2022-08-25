package model

type Post struct {
	Id      string `json:"id"`
	User    string `json:"user"`
	Message string `json:"message"`
	Url     string `json:"url"`
	Type    string `json:"type"`
}

type User struct {
    Username string `json:"username"` // `` 避免双引号意义重复
    Password string `json:"password"`
    Age      int64  `json:"age"`
    Gender   string `json:"gender"`
}

