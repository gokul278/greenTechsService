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