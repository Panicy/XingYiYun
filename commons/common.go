package commons


//输出json
type  Jsonstruct struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
}
type Datas struct {
	Src string `json:"src"`
	Title string `json:"title"`
}
type UploadImg struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
	Data Datas `json:"data"`
}