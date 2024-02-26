package tools

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

const (
	_               = iota
	DATE            = iota
	DATETIME        = iota
	YEARMONTH       = iota
	DATEHOUR        = iota
	DATEHOURMINUTES = iota
	TIME            = iota
	HOURMINUTES     = iota
)

func CreateRandomString(n int) string {
	s := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789_"
	rand.Intn(len(s) - 1)
	var res string = ""
	for i := 0; i < n; i++ {
		res += string(s[rand.Intn(len(s)-1)])
	}
	return res
}

func Sha256(s string) string {
	sum := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", sum)
}

func CreatePasswd(salt, passwd string) string {
	return Sha256(salt + passwd)
}

func Time2Read(t int64, category ...int) string {
	if t == 0 {
		return ""
	}
	c := DATETIME
	if len(category) >= 1 {
		c = category[0]
	}
	switch c {
	case DATE:
		return time.Unix(t, 0).Format("2006-01-02")
	case DATETIME:
		return time.Unix(t, 0).Format("2006-01-02 15:04:05")
	case YEARMONTH:
		return time.Unix(t, 0).Format("2006-01")
	case DATEHOUR:
		return time.Unix(t, 0).Format("2006-01-02 15")
	case DATEHOURMINUTES:
		return time.Unix(t, 0).Format("2006-01-02 15:04")
	case TIME:
		return time.Unix(t, 0).Format("15:04:05")
	case HOURMINUTES:
		return time.Unix(t, 0).Format("15:04")
	}
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

func PublishTime2Read(t int64) string {
	if t == 0 {
		return ""
	}
	temp := time.Now().Unix() - t
	if temp < 60 {
		return "1分钟以内"
	} else if temp < 3600 {
		return fmt.Sprintf("%d分钟前", temp/60)
	} else if temp < 86400 {
		return fmt.Sprintf("%d小时前", temp/3600)
	} else if temp < 5*86400 {
		return fmt.Sprintf("%d天前", temp/86400)
	} else {
		return time.Unix(t, 0).Format("06-01-02")
	}
}
