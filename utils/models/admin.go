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
type DashboardOrder struct {
	CompletedOrder int
	PendingOrder   int
	CancelledOrder int
	TotalOrder     int
	TotalOrderItem int
}
type DashboardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}
type DashboardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}
type CompleteAdminDashboard struct {
	DashboardUser    DashboardUser
	DashBoardProduct DashBoardProduct
	DashboardOrder   DashboardOrder
	DashboardRevenue DashboardRevenue
	DashboardAmount  DashboardAmount
}
type SalesReport struct {
	TotalSales      float64
	TotalOrders     int
	CompletedOrders int
	PendingOrders   int
	TrendingProduct string
}
