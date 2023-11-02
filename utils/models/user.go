package models

type SignupDetail struct {
	Firstname    string `json:"firstname"  validate:"required"`
	Lastname     string `json:"lastname"  validate:"required"`
	Email        string `json:"email"  validate:"required"`
	Phone        string `json:"phone"  validate:"required"`
	Password     string `json:"password"  validate:"required"`
	ReferralCode string `json:"referral_code"`
}

type SignupDetailResponse struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
type TokenUser struct {
	Users        SignupDetailResponse
	AccessToken  string
	RefreshToken string
}
type LoginDetail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type UserDetailsAtAdmin struct {
	Id          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	BlockStatus bool   `json:"block_status"`
}

type UpdatePassword struct {
	OldPassword        string `json:"old_password" binding:"required"`
	NewPassword        string `json:"new_password" binding:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}

type PaymentDetails struct {
	ID           uint   `json:"id"`
	Payment_Name string `json:"payment_name"`
}
type AddressInfoResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Street    string `json:"street"`
	City      string `json:"city"`
}
type AddressInfo struct {
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
	Street    string `json:"street"`
	City      string `json:"city"`
}
type UsersProfileDetails struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email" `
	Phone        string `json:"phone"`
	ReferralCode string `json:"referral_code" binding:"required"`
}

type CheckoutDetails struct {
	AddressInfoResponse []AddressInfoResponse
	Payment_Method      []PaymentDetails
	ReferralAmount      ReferralAmount
	Cart                []Cart
	Grand_Total         float64
	Total_Price         float64
	DiscountReason      []string
}
