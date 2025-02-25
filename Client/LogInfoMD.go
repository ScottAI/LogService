package Client

import (
	"encoding/json"
	"time"
)

//日志本身

type LogInfoMD struct {
	Topic1 string    `json:"Topic1"` //主题1
	Topic2 string    `json:"Topic2"` //主题2
	Topic3 string    `json:"Topic3"` //主题3
	MID    int       `json:"MID"`    //用户ID
	SID    int       `json:"SID"`    //服务器ID
	UpTime time.Time //写入时间
	UpNum  int64     //更新数量
	Total  int64     //更新后总数量
	Datas  string    //其他数据
}

var (
	sid int
)

func NewLogInfo(topid1, topid2, topid3 string, mid, sid int, dt time.Time, upnum, total int64, datas string) *LogInfoMD {
	result := new(LogInfoMD)
	result.Topic1 = topid1
	result.Topic2 = topid2
	result.Topic3 = topid3
	result.MID = mid
	result.SID = sid
	result.UpTime = dt
	result.UpNum = upnum
	result.Total = total
	result.Datas = datas
	return result
}

func NewLogInfoByNum(topid1, topid2, topid3 string, mid, sid int, dt time.Time, upnum, total int64) *LogInfoMD {
	return NewLogInfo(topid1, topid2, topid3, mid, sid, dt, upnum, total, "")
}

func NewLogInfoByIface(topid1, topid2, topid3 string, mid, sid int, dt time.Time, upnum, total int64, datas interface{}) *LogInfoMD {
	if datas == nil {
		return NewLogInfo(topid1, topid2, topid3, mid, sid, dt, upnum, total, "")
	}
	d, _ := json.Marshal(datas)
	return NewLogInfo(topid1, topid2, topid3, mid, sid, dt, upnum, total, string(d))
}

func NewLoginfoByTime(topid1, topid2, topid3 string, mid, sid int, dt time.Time) *LogInfoMD {
	return NewLogInfo(topid1, topid2, topid3, mid, sid, dt, 0, 0, "")
}
