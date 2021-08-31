package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
)

func GetAllVoucher() (interface{}, error) {
	var vouchers []models.Voucher
	if e := config.DB.Find(&vouchers).Error; e != nil {
		return nil, e
	}
	return vouchers, nil
}

func GetMyVoucher(userId int) (interface{}, error) {

	var vouchers []models.My_voucher
	query := config.DB.Raw("SELECT voucher_code, description, voucher_disc, minimum_transaction FROM user_vouchers LEFT JOIN vouchers ON vouchers.voucher_id = user_vouchers.voucher_id WHERE user_id = ? AND status = 1", userId).Scan(&vouchers)
	if err := query.Error; err != nil {
		return nil, err
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return vouchers, nil
}

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
	return disc
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
