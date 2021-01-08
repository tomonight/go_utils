// Copyright 2020  enmotech.chengdu. All rights reserved.
// Use of this utils to operate time

package utils

import (
	"strconv"
	"time"
)

const (
	standardTimeStr string = "2006-01-02 15:04:05"
)

//ParseStandardDateTime  2006-01-02 15:05:06
func ParseStandardDateTime(date time.Time) string {
	return date.Format(standardTimeStr)
}

//ParseNowStandardDateTime  now --> 2006-01-02 15:05:06
func ParseNowStandardDateTime() string {
	return time.Now().Format(standardTimeStr)
}

//ParseString2Timestamp  2006-01-02 15:05:06 -> 1136185506
func ParseString2Timestamp(date string) int64 {
	stamp, _ := time.ParseInLocation(standardTimeStr, date, time.Local)
	return stamp.Unix()
}

//ParseNow2Timestamp  now -> 1136185506
func ParseNow2Timestamp() int64 {
	return time.Now().Unix()
}

//1136185506 -> 2006-01-02 15:05:06
func ParseTimestamp2String(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

//"1136185506"  -> 2006-01-02 15:05:06
func ParseTimestampStr2String(ts string) string {
	t, _ := strconv.Atoi(ts)
	if t == 0 {
		return ""
	}

	return time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
}
