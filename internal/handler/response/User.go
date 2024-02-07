package response

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar,omitempty"`
	Token      string `json:"token"`
}
