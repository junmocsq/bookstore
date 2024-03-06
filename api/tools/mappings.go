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
	"book_status": {
		1: "上架",
		2: "审查",
		3: "下架",
	},
	"book_process": {
		1: "连载",
		2: "完结",
		3: "停更",
	},
	"book_is_pay": {
		1: "付费",
		2: "免费",
	},
	"section_status": {
		1: "发布",
		2: "定时发布",
		3: "审查",
		4: "下架",
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
