package model

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}
