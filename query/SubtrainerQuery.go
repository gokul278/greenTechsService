package query

var NewRegistrationSubtrainerUserSQL = `
INSERT INTO public.users (
  "refUserName",
  "refUserStatus",
  "refUserRTId",
  "refUserDOB",
  "refUserProfile",
  "refUserCreatedAt",
  "refUserCreatedBy",
  "refUserCustId"
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING "refUserId";
`

var NewRegistrationSubtrainerDomainSQL = `
INSERT INTO
  userdomain."userSubtrainerDomain" (
    "refUserId",
    "refSTDWorkExprience",
    "refSDTAadhar",
    "refSDTResume",
    "refSDTCreatedAt",
    "refSDTCreatedBy"
  )
VALUES
  ($1, $2, $3, $4, $5, $6)
`
