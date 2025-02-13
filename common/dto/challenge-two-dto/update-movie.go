package challenge_two_dto

type UpdateMovie struct {
	Id          int64
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
}
