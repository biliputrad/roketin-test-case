package challenge_two_dto

type CreateMovie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
}
