package entities

type Payload struct {
	Numbers []int `json:"numbers" validate:"required"`
	Target  int   `json:"target" validate:"required"`
}
