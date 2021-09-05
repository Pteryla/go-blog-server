package errcode

var (
	Success = NewError(0,"请求成功")
	ServerError = NewError(20001,"服务器错误")
	InvalidParams = NewError(20002,"参数错误")
	NotFound = NewError(20003,"找不到")
	TooManyRequests = NewError(20004,"请求过于频繁")
	UnauthorizedAuthNotExist = NewError(20005,"鉴权失败，找不到对应权限")
	UnauthorizedTokenError = NewError(20006 ,"鉴权失败，Token错误")
	UnauthorizedTokenTimeout = NewError(20007,"鉴权失败，Token超时")
	UnauthorizedTokenGeneration = NewError(20008 ,"鉴权失败，Token生成失败")
)


