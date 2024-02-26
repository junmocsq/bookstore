package tools

var m = map[string]map[int8]string{
	"gender": {
		0: "未知",
		1: "男",
		2: "女",
	},
	"recharge_channel": {
		1: "系统充值",
		2: "微信",
		3: "支付宝",
		4: "苹果",
	},
	"recharge_status": {
		0: "未充值",
		1: "已充值",
	},
	"transaction_category": {
		1: "充值",
		2: "订阅",
		3: "登录",
		4: "签到",
	},
	"transaction_status": {
		1: "收入",
		2: "支出",
	},
}

func Mapping(key string, id int8) string {
	r, ok := m[key]
	if !ok {
		return "unknown key"
	}
	res, ok := r[id]
	if !ok {
		return "unknown"
	}
	return res
}
