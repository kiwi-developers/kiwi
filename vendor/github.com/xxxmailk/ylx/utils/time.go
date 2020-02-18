package utils

import "time"

// return start timestamp in this day
// 返回本天第一个时间节点的时间戳
func ThisDay(t time.Time) (start int64) {
	year, month, day := t.Date()
	m := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	start = m.Unix()
	return
}

// return start timestamp in this month
// 返回本月第一个时间节点的时间戳
func ThisMonth(t time.Time) (start int64) {
	year, month, _ := t.Date()
	m := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start = m.Unix()
	return
}

// return start timestamp in this quarte
// 返回本季度第一个时间节点的时间戳
func ThisQuarte(t time.Time) (start int64) {
	year, month, _ := t.Date()
	month = month - (month-1)%3
	m := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start = m.Unix()
	return
}

// return start timestamp in this year
// 返回本年第一个时间节点的时间戳
func ThisYear(t time.Time) (start int64) {
	year, _, _ := t.Date()
	m := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	start = m.Unix()
	return
}
