package errorss

import "errors"

var ErrOfferAlreadyexist = errors.New("the offer already exists")
var ErrCouponAlreadyexist = errors.New("the offer already exists")
var ErrFailedTovalidateOtp = errors.New("failed to validate otp")
var ErrEmailAlreadyExist = errors.New("user with this email is already exists")
var ErrPhoneAlreadyExist = errors.New("user with this phonenumber is already exists")
