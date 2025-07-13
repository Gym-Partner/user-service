package constants

// ######################################################################################
// 											LOG
// ######################################################################################

const (
	ServiceErrRun     = "[RUN] failed to run use cases | [ORIGINAL-ERROR] %s"
	ServiceErrMigrate = "[DB][Postgres][MIGRATE] failed to run migration | [ORIGINAL-ERROR] %s"

	ServiceErrDBUserNotFound      = "[USER][SERVICE][REPOSITORY] User not found or not exist in database"
	ServiceErrDBCreateUser        = "[USER][SERVICE][REPOSITORY] Failed to create user | [ORIGINAL-ERROR] %s"
	ServiceErrDBGetAllUsers       = "[USER][SERVICE][REPOSITORY] Failed to get all users | [ORIGINAL-ERROR] %s"
	ServiceErrDBGetOneUserByID    = "[USER][SERVICE][REPOSITORY] Failed to get one user by his ID | [ORIGINAL-ERROR] %s"
	ServiceErrDBGetOneUserByEmail = "[USER][SERVICE][REPOSITORY] Failed to get one user by his email | [ORIGINAL-ERROR] %s"
)

// ######################################################################################
// 									RESPONSE REQUEST
// ######################################################################################

const (
	// REPOSITORY PART

	ServiceErrAppDBCreateUser        = "Failed to create user {%s} in database."
	ServiceErrAppDBGetAllUsers       = "Failed to get all users in the database."
	ServiceErrAppDBGetOneUserByID    = "Failed to get one user by his ID."
	ServiceErrAppDBGetOneUserByEmail = "Failed to get one user by his email."

	// SERVICE PART

	ServiceErrAppINTUserAlreadyExist = "User {%s} already exists."
)
