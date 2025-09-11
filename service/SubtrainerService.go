package service

import (
	"fmt"
	becrypt "greenstech/helper/Becrypt"
	logger "greenstech/helper/Logger"
	timeZone "greenstech/helper/TimeZone"
	"greenstech/model"
	"greenstech/query"

	"gorm.io/gorm"
)

func NewSubtrainerRegistrationServices(db *gorm.DB, reqVal model.ReqnewSubtrainerRegistrationModel, idValue int) (bool, string) {
	log := logger.InitLogger()

	tx := db.Begin()
	if tx.Error != nil {
		log.Printf("ERROR: Failed to begin transaction: %v\n", tx.Error)
		return false, "Something went wrong, Try Again"
	}

	defer func() {
		if r := recover(); r != nil {
			log.Error("ERROR: Recovered from panic, rolling back transaction:", r)
			tx.Rollback()
		}
	}()

	//checkEmail and Mobile Number is Already Registered
	var checkEmail model.CheckemailPhoneNumberModel
	checkEmailErr := tx.Raw(query.Checkmailphonenumber, reqVal.Emailid, reqVal.Phonenumber).Scan(&checkEmail).Error
	if checkEmailErr != nil {
		log.Error(checkEmailErr.Error())
		return false, "Something went wrong, Try Again"
	}

	//Check Email Already Exits
	switch checkEmail.IsEmailExist {
	case true:
		return false, "Email Already Exits"
	case false:
		break
	}

	//Phone Number Already Exits
	switch checkEmail.IsPhoneNumberExist {
	case true:
		return false, "Phone Number Already Exits"
	case false:
		break
	}

	//Generating CustId
	var count int

	countErr := db.Raw(query.GetUserLatestCountSQL, 3).Scan(&count).Error
	if countErr != nil {
		log.Error(countErr.Error())
		return false, "Something went wrong, Try Again"
	}

	// Format the count to a 5-digit string with leading zeros
	formattedCount := fmt.Sprintf("%05d", count+100001)

	// Generate the custom user ID
	customerID := "GTST" + formattedCount

	//Performing the Student Registration
	var userId int

	//Inserting into the User Table
	userDataErr := tx.Raw(
		query.NewRegistrationSubtrainerUserSQL,
		reqVal.Fullname,
		true,
		3, //Sub Trainer
		reqVal.Dob,
		reqVal.ProfileImage,
		timeZone.GetPacificTime(),
		idValue,
		customerID,
	).Scan(&userId).Error
	if userDataErr != nil {
		log.Error(userDataErr.Error())
		return false, "Something went wrong, Try Again"
	}

	//Inserting Communication
	communicationDataErr := tx.Exec(
		query.NewRegistrationCommunicationSQL,
		userId,
		reqVal.CurrentLocation,
		reqVal.Phonenumber,
		nil,
		reqVal.Emailid,
	).Error
	if communicationDataErr != nil {
		log.Error(communicationDataErr.Error())
		return false, "Something went wrong, Try Again"
	}

	//Password Encrypt
	hashPassword, hashPassworderr := becrypt.HashPassword(reqVal.Dob)

	if hashPassworderr != nil {
		log.Printf("ERROR: Failed to hash password: %v\n", hashPassworderr)
		return false, "Something went wrong, Try Again"
	}

	//Inserting the Password
	AuthPasswordErr := tx.Exec(
		query.NewRegistrationPasswordSQL,
		userId,
		reqVal.Dob,
		hashPassword,
		true,
	).Error
	if AuthPasswordErr != nil {
		log.Error(AuthPasswordErr.Error())
		return false, "Something went wrong, Try Again"
	}

	//Inserting the Sub Trainer Domain
	SubtrainerDomainDataErr := tx.Exec(
		query.NewRegistrationSubtrainerDomainSQL,
		userId,
		reqVal.Workexprience,
		reqVal.Aadhar,
		reqVal.Resume,
		timeZone.GetPacificTime(),
		idValue,
	).Error
	if SubtrainerDomainDataErr != nil {
		log.Error(SubtrainerDomainDataErr.Error())
		return false, "Something went wrong, Try Again"
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("ERROR: Failed to commit transaction: %v\n", err)
		tx.Rollback()
		return false, "Something went wrong, Try Again"
	}

	return true, "Successfully Registered"

}
