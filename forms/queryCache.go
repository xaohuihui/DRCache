package forms

// author: xaohuihui
// datetime: 2022/3/7 15:10:35
// software: GoLand

type QueryCacheForm struct {
	Key string `form:"key" json:"key" binding:"required"`
}

type SetCacheParse struct {
	Key   string `form:"key" json:"key" binding:"required"`
	Value string `form:"value" json:"value" binding:"required"`
	Timeout int `form:"timeout" json:"timeout" binding:"required,default=1000"`
}
