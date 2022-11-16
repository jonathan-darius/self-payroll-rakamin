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
