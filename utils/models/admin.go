package models

type AdminDetail struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type DashboardUser struct {
	TotalUsers  int
	BlockedUser int
}

type DashBoardProduct struct {
	TotalProducts     int
	OutOfStockProduct int
}
type CompleteAdminDashboard struct {
	DashboardUser    DashboardUser
	DashBoardProduct DashBoardProduct
}
