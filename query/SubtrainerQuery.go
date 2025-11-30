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
  u."refUserRTId" = 3
  AND (
    $1 = 0
    OR u."refUserId" = $1
  )
ORDER BY
  u."refUserId" DESC
`

var UpdateSubtrainerUserSQL = `
UPDATE
  public.users
SET
  "refUserName" = $1,
  "refUserStatus" = $2,
  "refUserDOB" = $3,
  "refUserProfile" = $4,
  "refUserUpdatedAt" = $5,
  "refUserUpdatedBy" = $6
WHERE
  "refUserId" = $7;
`

var UpdateSubtrianerCommunicationSQL = `
UPDATE
  userdomain."userCommunication"
SET
  "refUCAddress" = $1,
  "refUCMobileno" = $2
WHERE
  "refUserId" = $3;
`

var UpdateSubtrainerDomainSQL = `
UPDATE
  userdomain."userSubtrainerDomain"
SET
  "refSTDWorkExprience" = $1,
  "refSDTAadhar" = $2,
  "refSDTResume" = $3
WHERE
  "refUserId" = $4;
`
