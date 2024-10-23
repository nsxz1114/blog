package ctype

type ImageType int

const (
	Local ImageType = 1 // 本地
	QiNiu ImageType = 2 // 云端
)

// func (s ImageType) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(s.String())
// }

// func (s ImageType) String() string {
// 	var str string
// 	switch s {
// 	case Local:
// 		str = "本地"
// 	case QiNiu:
// 		str = "云端"
// 	default:
// 		str = "其他"
// 	}
// 	return str
// }
