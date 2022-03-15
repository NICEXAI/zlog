package zlog

import "log"

func Println(v ...interface{}) {
	log.Println(v)
}
