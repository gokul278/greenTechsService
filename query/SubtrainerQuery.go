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

var ListSubtrainerSQL = `
SELECT
  *
FROM
  public."users" u
  JOIN userdomain."userCommunication" uc ON uc."refUserId" = u."refUserId"
  JOIN userdomain."userSubtrainerDomain" usd ON usd."refUserId" = u."refUserId"
WHERE
  u."refUserStatus" = true
  AND u."refUserRTId" = 3
  AND (
    $1 = 0
    OR u."refUserId" = $1
  )
ORDER BY
  u."refUserId" DESC
`
