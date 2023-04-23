package base

import "strings"

const (
	ZeusMediaName = "zeusmedia"

	ZeusMediaVersion = "v0.1.0"
)

var (
	ZeusMediaVersionDot = strings.TrimPrefix(ZeusMediaVersion, "v")
	ZeusPullSessionUa   = ZeusMediaName + "/" + ZeusMediaVersionDot
)
