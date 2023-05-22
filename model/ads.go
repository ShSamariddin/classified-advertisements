package model

type AdType int
type AdStatus int
type SaleType int
type Currency int

const (
	house AdType = iota
)

const (
	active AdStatus = iota
	inactive
)

const (
	longTerm SaleType = iota
	daily
)

const (
	USD Currency = iota
	EUR
	TJS
	RUB
	UZS
)
