package common

// author: songyanhui
// datetime: 2022/2/17 14:32:42
// software: GoLand

type BusError struct {}

func (b *BusError) ThrowError(err error)  {
	if err != nil {
		panic(err.Error())
	}
}
