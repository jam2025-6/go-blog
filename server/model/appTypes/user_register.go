package appTypes

import "encoding/json"

// Register 用户注册来源
type Register int

const (
	Email Register = iota // 邮箱验证码注册
	QQ                    // QQ 登录注册
)

// MarshalJSON 实现了 json.Marshaler 接口
func (r Register) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Register) UmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	*r = ToRegister(str)
	return nil
}

func (r Register) String() string {
	var str string
	switch r {
	case Email:
		str = "邮箱"
	case QQ:
		str = "QQ"
	default:
		str = "未知"
	}
	return str
}

func ToRegister(strings string) Register {
	switch strings {
	case "邮箱":
		return Email
	case "QQ":
		return QQ
	case "未知":
		return -1
	default:
		return -1
	}
}
