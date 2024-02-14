package response

type Campaign struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserId           int    `json:"user_id"`
}
