package utils

type PositionUtils struct {
	Name      string `json:"name"`
	Salary    int    `json:"salary"`
	ID        int    `json:"id"`
	CreatedAt string `json:"create_at"`
	UpdatedAt string `json:"updated_at"`
}

type CompanyUtils struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Balance   int    `json:"balance"`
	ID        int    `json:"id"`
	CreatedAt string `json:"create_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserUtils struct {
	ID         int            `json:"id"`
	CreatedAt  string         `json:"create_at"`
	UpdatedAt  string         `json:"updated_at"`
	SecretID   string         `json:"secret_id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	Address    string         `json:"address"`
	PositionID int            `json:"position_id"`
	Position   *PositionUtils `json:"position"`
}
