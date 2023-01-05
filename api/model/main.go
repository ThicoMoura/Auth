package model

type Model interface {
	Get(string) interface{}
}
