package model

type ReqnewSubtrainerRegistrationModel struct {
	Fullname        string `json:"fullname" binding:"required" mapstructure:"fullname"`
	Phonenumber     string `json:"phonenumber" binding:"required" mapstructure:"phonenumber"`
	Emailid         string `json:"emailid" binding:"required" mapstructure:"emailid"`
	Dob             string `json:"dob" binding:"required" mapstructure:"dob"`
	CurrentLocation string `json:"currentLocation" mapstructure:"currentLocation"`
	Workexprience   string `json:"workexprience" mapstructure:"workexprience"`
	Aadhar          string `json:"aadhar" mapstructure:"aadhar"`
	ProfileImage    string `json:"profileImage" mapstructure:"profileImage"`
	Resume          string `json:"resume" mapstructure:"resume"`
}

type ReqGetSubtrainerRegistrationModel struct {
	Id int `json:"id" binding:"required" mapstructure:"id"`
}

type SubtrainerListModel struct {
	UserId              int    `json:"refUserId" gorm:"column:refUserId"`
	RefUserName         string `json:"refUserName" gorm:"column:refUserName"`
	RefUserStatus       string `json:"refUserStatus" gorm:"column:refUserStatus"`
	RefUserDOB          string `json:"refUserDOB" gorm:"column:refUserDOB"`
	RefUserProfile      string `json:"refUserProfile" gorm:"column:refUserProfile"`
	RefUserCustId       string `json:"refUserCustId" gorm:"column:refUserCustId"`
	RefUCAddress        string `json:"refUCAddress" gorm:"column:refUCAddress"`
	RefUCMobileno       string `json:"refUCMobileno" gorm:"column:refUCMobileno"`
	RefUCMail           string `json:"refUCMail" gorm:"column:refUCMail"`
	RefSTDWorkExprience string `json:"refSTDWorkExprience" gorm:"column:refSTDWorkExprience"`
	RefSDTAadhar        string `json:"refSDTAadhar" gorm:"column:refSDTAadhar"`
	RefSDTResume        string `json:"refSDTResume" gorm:"column:refSDTResume"`
}
