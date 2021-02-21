package types

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	IsMale   bool   `json:"isMale"`
	Age      int    `json:"age"`
}

type NewRegister struct {
	Message string `json:"message"`
	ID      int64  `json:"id"`
}
