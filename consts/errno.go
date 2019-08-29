package consts

const (
	Success = 0

	// 参数不合法
	ParamsDecodeError = 10001 + iota

	NetWorkError
)

var ErrMsg = map[int]string{
	Success:           "success",
	ParamsDecodeError: "参数不合法",
	NetWorkError:      "网络错误",
}
