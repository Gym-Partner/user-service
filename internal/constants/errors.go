package constants

const (
	ServiceErrRun     = "[RUN] failed to run use cases | [ORIGINAL-ERROR] %s"
	ServiceErrMigrate = "[DB][Postgres][MIGRATE] failed to run migration | [ORIGINAL-ERROR] %s"

	ServiceErrDBUserNotFound = "[USER][SERVICE][REPOSITORY] User not found or not exist in database"
	ServiceErrDBCreateUser   = "[USER][SERVICE][REPOSITORY] Failed to create user | [ORIGINAL-ERROR] %s"
)

const (
	ServiceErrAppDBCreateUser = "Failed to create user {%s} in database."

	ServiceErrAppINTUserAlreadyExist = "User {%s} already exists."
)
