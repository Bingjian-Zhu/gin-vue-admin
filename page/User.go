package page

//User 用户登录信息结构体
type User struct {
	Roles        *[]string
	Introduction string
	Avatar       string
	Name         string
}

//Users 用户管理结构体
type Users struct {
	ID        int    `json:"id"`
	Name      string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	UserType  string `json:"user_type"`
	State     string `json:"state"`
	Deteled   string `json:"deteled"`
	CreatedOn string `json:"created_on"`
}
