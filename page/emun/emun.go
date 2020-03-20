package emun

//Status 启用/禁用状态
var Status = map[int]string{
	0: "禁用", //禁用
	1: "正常", //正常
}

//GetStatus 获取状态信息
func GetStatus(code int) string {
	msg, ok := Status[code]
	if ok {
		return msg
	}

	return Status[0]
}

//UserType 用户类型
var UserType = map[int]string{
	1: "管理员",  //管理员
	2: "测试用户", //测试用户
}

//GetUserType 获取状态信息
func GetUserType(code int) string {
	msg, ok := UserType[code]
	if ok {
		return msg
	}

	return UserType[0]
}

//Deleted 逻辑删除状态
var Deleted = map[int]string{
	0: "正常",  //正常
	1: "已删除", //已删除
}

//GetDeleted 获取删除状态
func GetDeleted(code int) string {
	msg, ok := Deleted[code]
	if ok {
		return msg
	}

	return Deleted[1]
}
