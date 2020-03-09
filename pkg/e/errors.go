package e

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err Err) Error() string {
	return fmt.Sprintf("ErrCode:%d , ErrMsg:%s", err.Code, err.Msg)
}