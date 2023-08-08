package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRes struct {
	JWT string `json:"jwt"`
}

type Login struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}
