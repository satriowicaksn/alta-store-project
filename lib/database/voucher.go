package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
)

func ClaimUserVoucher(userId, voucherId int) (bool, error) {
	userVoucher := models.User_voucher{
		User_id:    userId,
		Voucher_id: voucherId,
		Status:     1,
	}
	if err := config.DB.Create(&userVoucher).Error; err != nil {
		return false, err
	}
	return true, nil
}

func ValidateUserVoucher(userId int, voucherCode string) (bool, string) {
	vouchers := models.Voucher{}
	getVoucherDetail := config.DB.Where("voucher_code = ?", voucherCode).Find(&vouchers)
	if getVoucherDetail.RowsAffected == 0 {
		return false, "Invalid voucher code"
	}
	voucherId := vouchers.Voucher_id
	user_voucher := models.User_voucher{}
	validateVoucher := config.DB.Where("user_id = ? and voucher_id = ?", userId, voucherId).Find(&user_voucher)
	if user_voucher.Status != 1 {
		return false, "You have used this voucher code"
	}
	if validateVoucher.RowsAffected == 0 {
		return false, "You don't have this voucher"
	}
	return true, ""
}

func GetVoucherDiscount(amount int, voucherCode string) int {
	vouchers := models.Voucher{}
	config.DB.Where("voucher_code = ?", voucherCode).Find(&vouchers)
	voucherType := vouchers.Voucher_type
	if voucherType == 1 {
		return vouchers.Voucher_disc
	}
	disc := amount * vouchers.Voucher_disc / 100
	totalDisc := amount - disc
	return totalDisc
}

func UseVoucher(user_voucher_id int) bool {
	userVoucher := models.User_voucher{
		Status: 0,
	}
	query := config.DB.Where("user_voucher_id = ?", user_voucher_id).Updates(&userVoucher)
	if query.RowsAffected == 0 {
		return false
	}
	return true
}
