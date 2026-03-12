package calculationService

type Calculation struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
	UserID     string `json:"user_id"`
}
