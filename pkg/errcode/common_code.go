package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服務內部錯誤")
	InvalidParams             = NewError(10000001, "入參錯誤")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鑒權失敗，找不到對應的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鑒權失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(10000005, "鑒權失敗，Token超時")
	UnauthorizedTokenGenerate = NewError(10000006, "鑒權失敗，Token產生失敗")
	TooManyRequests           = NewError(10000007, "請求過多")
)
