package serializer

// Response 基础序列化器
// omitempty如果返回的数据这个字段为空，则序列化出来的数据没有这个字段
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

type SsopaResponse struct {
	Response
	ResCode int `json:"rescode"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// user相关
	USERNOTEXISTS     = 10004 // 需要跳转到注册
	USEREXISTS        = 10002 // 需要跳转到登录
	USERCREATEERROR   = 10003
	PASSWORDERROR     = 10005
	LOGINSUCCESS      = 10001 // 需要跳转到首页
	CREATEUSERSUCCESS = 10000 // 需要跳转到登录
	SETCOOKIEFAILURE  = 10006
	LOGOUTSUCCESS     = 10007

	SUCCESS     = 20001
	PARAMSERROR = 40001
	NOTFOUND    = 40004
	SYSTEMERROR = 50001
)
