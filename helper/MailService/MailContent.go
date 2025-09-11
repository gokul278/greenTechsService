package mailservice

import (
	timeZone "greenstech/helper/TimeZone"
	"html"
	"strconv"
	"time"
)

func LoginOTPContent(otp int) string {
	return `
	<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Login Passcode</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f6f8fa;
      margin: 0;
      padding: 0;
    }
    .container {
      max-width: 600px;
      background-color: #EDD1CE;
      margin: 40px auto;
      padding: 30px;
      border-radius: 8px;
      box-shadow: 0 0 10px rgba(0,0,0,0.05);
    }
    .header {
      text-align: center;
      padding-bottom: 20px;
    }
    .header h1 {
      margin: 0;
      color: #525252;
    }
    .otp {
      font-size: 32px;
      font-weight: 700;
      color: #333333;
      text-align: center;
      margin: 20px 0;
      letter-spacing: 2px;
    }
    .content {
      font-size: 16px;
      color: #525252;
      text-align: center;
      margin-bottom: 30px;
    }
    .footer {
      font-size: 12px;
      text-align: center;
      color: #525252;
      border-top: 1px solid #dddddd;
      padding-top: 15px;
    }
  </style>
</head>
<body>
  <div class="container">
    <div class="header">
      <h1>Login Passcode Verification</h1>
    </div>
    <div class="content">
      <p>Use the following Passcode to log in. This Passcode is valid for 10 minutes:</p>
      <div class="otp">` + html.EscapeString(strconv.Itoa(otp)) + `</div>
      <p>If you did not request this Passcode, please ignore this email.</p>
    </div>
    <div class="footer">
      &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + ` Wellthgreen Theranostics. All rights reserved.
    </div>
  </div>
</body>
</html>
	`
}

func ForgetPasswordOTPContent(otp int) string {
	return `
			<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Login Passcode</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f6f8fa;
      margin: 0;
      padding: 0;
    }
    .container {
      max-width: 600px;
      background-color: #EDD1CE;
      margin: 40px auto;
      padding: 30px;
      border-radius: 8px;
      box-shadow: 0 0 10px rgba(0,0,0,0.05);
    }
    .header {
      text-align: center;
      padding-bottom: 20px;
    }
    .header h1 {
      margin: 0;
      color: #525252;
    }
    .otp {
      font-size: 32px;
      font-weight: bold;
      color: #333333;
      text-align: center;
      margin: 20px 0;
      letter-spacing: 2px;
    }
    .content {
      font-size: 16px;
      color: #525252;
      text-align: center;
      margin-bottom: 30px;
    }
    .footer {
      font-size: 12px;
      text-align: center;
      color: #525252;
      border-top: 1px solid #dddddd;
      padding-top: 15px;
    }
  </style>
</head>
<body>
  <div class="container">
    <div class="header">
      <h1>Forget Password Passcode Verification</h1>
    </div>
    <div class="content">
      <p>Use the following Passcode to Forget Password. This Passcode is valid for 10 minutes:</p>
      <div class="otp">` + html.EscapeString(strconv.Itoa(otp)) + `</div>
      <p>If you did not request this Passcode, please ignore this email.</p>
    </div>
    <div class="footer">
      &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + ` Wellthgreen Theranostics. All rights reserved.
    </div>
  </div>
</body>
</html>
	`
}

func RegisterationMailContent(userName string) string {
	return `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Welcome to Wellthgreen HealthCare Pvt Ltd</title>
    <style>
      body {
        font-family: Arial, sans-serif;
        background-color: #f6f8fa;
        margin: 0;
        padding: 0;
      }
      .container {
        max-width: 600px;
        background-color: #EDD1CE;
        margin: 40px auto;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
      }
      .header {
        text-align: center;
        padding-bottom: 20px;
      }
      .header h1 {
        margin: 0;
        color: #525252;
      }
      .content {
        font-size: 16px;
        color: #525252;
        text-align: center;
        margin-bottom: 30px;
      }
      .button {
        display: inline-block;
        padding: 12px 25px;
        background-color: #c6d4c0;
        color: #f7f7f7;
        border-radius: 5px;
        text-decoration: none;
        font-weight: bold;
        margin-top: 20px;
      }
      .footer {
        font-size: 12px;
        text-align: center;
        color: #525252;
        border-top: 1px solid #dddddd;
        padding-top: 15px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>Welcome, ` + html.EscapeString(userName) + `!</h1>
      </div>
      <div class="content">
        <p>
          Thank you for registering with <strong>Wellthgreen Theranostics</strong>.
        </p>
        <p>
          We’re excited to have you on board. You can now log in using your
          registered email and the password you created.
        </p>
        <a href="https://easeqt.brightoncloudtech.com/" class="button">Login Now</a>
        <p>If you didn’t register with us, please ignore this email.</p>
      </div>
      <div class="footer">
        &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + `
        Wellthgreen Theranostics. All rights reserved.
      </div>
    </div>
  </body>
</html>
  `
}

func GetOTPMailContent(userName string, otp int) string {
	return `
 <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Welcome to Wellthgreen HealthCare Pvt Ltd</title>
    <style>
      body {
        font-family: Arial, sans-serif;
        background-color: #f6f8fa;
        margin: 0;
        padding: 0;
      }
      .container {
        max-width: 600px;
        background-color: #edd1ce;
        margin: 40px auto;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
      }
      .header {
        text-align: center;
        padding-bottom: 20px;
      }
      .header h1 {
        margin: 0;
        color: #525252;
      }
      .otp-box {
        font-size: 24px;
        font-weight: bold;
        background-color: #f0f0f0;
        padding: 15px;
        border-radius: 8px;
        margin: 20px auto;
        text-align: center;
        color: #333333;
        letter-spacing: 4px;
        max-width: 200px;
      }
      .content {
        font-size: 16px;
        color: #525252;
        text-align: center;
        margin-bottom: 30px;
      }
      .button {
        display: inline-block;
        padding: 12px 25px;
        background-color: #c6d4c0;
        color: #ffffff;
        border-radius: 5px;
        text-decoration: none;
        font-weight: bold;
        margin-top: 20px;
      }
      .footer {
        font-size: 12px;
        text-align: center;
        color: #525252;
        border-top: 1px solid #dddddd;
        padding-top: 15px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>Welcome, ` + html.EscapeString(userName) + `!</h1>
      </div>
      <div class="content">
        <p>
          Thank you for registering with
          <strong>Wellthgreen Theranostics</strong>.
        </p>
        <p>
          Please use the Passcode below to verify your email address and complete
          your registration:
        </p>
        <div class="otp-box">` + html.EscapeString(strconv.Itoa(otp)) + `</div>
        <p>This Passcode is valid for the next 10 minutes.</p>
        <p>If you didn’t register with us, please ignore this email.</p>
      </div>
      <div class="footer">
        &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + `
        Wellthgreen Theranostics. All rights reserved.
      </div>
    </div>
  </body>
</html>
  `
}

func RegistrationMailContent(userName, patientID, gmail, password string, role string) string {
	return `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Welcome to Wellthgreen HealthCare Pvt Ltd</title>
    <style>
      body {
        font-family: Arial, sans-serif;
        background-color: #f6f8fa;
        margin: 0;
        padding: 0;
      }
      .container {
        max-width: 600px;
        background-color: #edd1ce;
        margin: 40px auto;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
      }
      .header {
        text-align: center;
        padding-bottom: 20px;
      }
      .header h1 {
        margin: 0;
        color: #525252;
      }
      .content {
        font-size: 16px;
        color: #525252;
        text-align: center;
        margin-bottom: 30px;
      }
      .credentials {
        background-color: #fff;
        padding: 15px;
        border-radius: 5px;
        margin: 20px auto;
        width: fit-content;
        text-align: left;
        font-family: monospace;
        border: 1px solid #ccc;
      }
      .button {
        display: inline-block;
        padding: 12px 25px;
        background-color: #c6d4c0;
        color: #ffffff;
        border-radius: 5px;
        text-decoration: none;
        font-weight: bold;
        margin-top: 20px;
      }
      .footer {
        font-size: 12px;
        text-align: center;
        color: #525252;
        border-top: 1px solid #dddddd;
        padding-top: 15px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>Welcome, ` + html.EscapeString(userName) + `!</h1>
      </div>
      <div class="content">
        <p>
          You have successfully been onboarded as a ` + html.EscapeString(role) + ` at <strong>Wellthgreen Theranostics</strong>.
        </p>
        <p>Your login credentials are as follows:</p>
        <div class="credentials">
          <p>
            <strong>` + html.EscapeString(role) + ` ID:</strong> ` +
		html.EscapeString(patientID) + `
          </p>
          <p>
            <strong>Email (Gmail):</strong> ` + html.EscapeString(gmail) + `
          </p>
          <p><strong>Password:</strong> ` + html.EscapeString(password) + `</p>
        </div>
        <a href="https://easeqt.brightoncloudtech.com/" class="button"
          >Login to Dashboard</a
        >
        <p style="margin-top: 20px">
          If you did not request this registration, please ignore this email.
        </p>
      </div>
      <div class="footer">
        &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + `
        Wellthgreen Theranostics. All rights reserved.
      </div>
    </div>
  </body>
</html>

`
}

func PatientReportSignOff(userName string, patientID string, AppintmentDate string, scancenterCode string) string {
	return `
  <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Your Report is Ready - Wellthgreen HealthCare Pvt Ltd</title>
    <style>
      body {
        font-family: Arial, sans-serif;
        background-color: #f6f8fa;
        margin: 0;
        padding: 0;
      }
      .container {
        max-width: 600px;
        background-color: #edd1ce;
        margin: 40px auto;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
      }
      .header {
        text-align: center;
        padding-bottom: 20px;
      }
      .header h1 {
        margin: 0;
        color: #525252;
      }
      .content {
        font-size: 16px;
        color: #525252;
        text-align: center;
        margin-bottom: 30px;
      }
      .report-info {
        background-color: #fff;
        padding: 15px;
        border-radius: 5px;
        margin: 20px auto;
        width: fit-content;
        text-align: left;
        font-family: monospace;
        border: 1px solid #ccc;
      }
      .button {
        display: inline-block;
        padding: 12px 25px;
        background-color: #c6d4c0;
        color: #ffffff;
        border-radius: 5px;
        text-decoration: none;
        font-weight: bold;
        margin-top: 20px;
      }
      .highlight {
        background-color: #e8f5e8;
        padding: 15px;
        border-radius: 5px;
        border-left: 4px solid #28a745;
        margin: 20px 0;
        font-weight: bold;
        color: #155724;
      }
      .footer {
        font-size: 12px;
        text-align: center;
        color: #525252;
        border-top: 1px solid #dddddd;
        padding-top: 15px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>Report Ready - ` + html.EscapeString(userName) + `!</h1>
      </div>
      <div class="content">
        <div class="highlight">
          📋 Your scan report has been completed!
        </div>
        <p>
          Dear ` + html.EscapeString(userName) + `,
        </p>
        <p>
          Your report has been processed and is now available for download.
        </p>
        
        <div class="report-info">
          <p><strong>Patient ID:</strong> ` + html.EscapeString(patientID) + `</p>
          <p><strong>Appointment Date:</strong> ` + html.EscapeString(AppintmentDate) + `</p>
          <p><strong>Scan Center Code:</strong> ` + html.EscapeString(scancenterCode) + `</p>
        </div>

        <p>To access your report, please log in with your credentials:</p>
        
        <a href="https://easeqt.brightoncloudtech.com/" class="button">
          Login to View Report
        </a>
        
        <p style="margin-top: 15px; font-size: 14px;">
          If you have any questions about your report or need assistance accessing it, 
          please contact your scan center.
        </p>
      </div>
      <div class="footer">
        &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + `
        Wellthgreen Theranostics. All rights reserved.<br>
      </div>
    </div>
  </body>
</html>
  `
}

func ManagerReportSignOff(patientName string, patientID string, appointmentDate string, scanCenterCode string) string {
	return `
  <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Report Ready for Admin Review - Wellthgreen HealthCare Pvt Ltd</title>
    <style>
      body {
        font-family: Arial, sans-serif;
        background-color: #f6f8fa;
        margin: 0;
        padding: 0;
      }
      .container {
        max-width: 600px;
        background-color: #edd1ce;
        margin: 40px auto;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
      }
      .header {
        text-align: center;
        padding-bottom: 20px;
      }
      .header h1 {
        margin: 0;
        color: #525252;
      }
      .content {
        font-size: 16px;
        color: #525252;
        margin-bottom: 30px;
      }
      .report-info {
        background-color: #fff;
        padding: 20px;
        border-radius: 5px;
        margin: 20px 0;
        border: 1px solid #ccc;
        font-family: monospace;
        text-align: left;
      }
      .button {
        display: inline-block;
        padding: 12px 25px;
        background-color: #c6d4c0;
        color: #ffffff;
        border-radius: 5px;
        text-decoration: none;
        font-weight: bold;
        margin: 10px 5px;
      }
      .highlight {
        background-color: #e8f5e8;
        padding: 15px;
        border-radius: 5px;
        border-left: 4px solid #28a745;
        margin: 20px 0;
        font-weight: bold;
        color: #155724;
      }
      .action-required {
        background-color: #f8d7da;
        padding: 15px;
        border-radius: 5px;
        border-left: 4px solid #dc3545;
        margin: 20px 0;
        font-weight: bold;
        color: #721c24;
      }
      .footer {
        font-size: 12px;
        text-align: center;
        color: #525252;
        border-top: 1px solid #dddddd;
        padding-top: 15px;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>📋 Admin Notification - Report Ready</h1>
      </div>
      <div class="content">
        <div class="highlight">
          ✅ Medical report has been processed    </div>
        
        <p>Dear Scan Center Administrator,</p>
        
        <p>A medical report has been completed.</p>
        
        <div class="report-info">
          <h3 style="margin-top: 0; color: #525252;">Report Details</h3>
          <p><strong>Patient Name:</strong> ` + html.EscapeString(patientName) + `</p>
          <p><strong>Patient ID:</strong> ` + html.EscapeString(patientID) + `</p>
          <p><strong>Appointment Date:</strong> ` + html.EscapeString(appointmentDate) + `</p>
          <p><strong>Scan Center Code:</strong> ` + html.EscapeString(scanCenterCode) + `</p>
          <p><strong>Processing Status:</strong> <span style="color: #28a745;">Completed</span></p>
        </div>

        <div style="text-align: center; margin: 25px 0;">
          <a href="https://easeqt.brightoncloudtech.com/" class="button">
            Login
          </a>
        </div>
        
        <p style="font-size: 14px;">
          If you need assistance or have questions about this report, please contact 
          wellthgreen theranostics.
        </p>
      </div>
      <div class="footer">
        &copy; ` + html.EscapeString(strconv.Itoa(time.Now().In(timeZone.MustGetPacificLocation()).Year())) + `
        Wellthgreen Theranostics. All rights reserved.<br>
        <small>This is an automated notification. Please do not reply to this email.</small>
      </div>
    </div>
  </body>
</html>
 `
}
