package types

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	IsMale   bool   `json:"isMale"`
	Age      int    `json:"age"`
}
