package constant

// Env
const ADDR string = "ADDR"
const ALLOW_WORKER string = "ALLOW_WORKER"
const DB_HOST string = "DB_HOST"
const DB_PORT string = "DB_PORT"
const DB_USERNAME string = "DB_USERNAME"
const DB_PASSWORD string = "DB_PASSWORD"
const DB_NAME string = "DB_NAME"
const DB_SSL_MODE string = "DB_SSL_MODE"
const DB_TIMEZONE string = "DB_TIMEZONE"
const JWT_AT_SECRET string = "JWT_AT_SECRET"
const JWT_SS string = "JWT_SS"

// JWT errors
const EXPIRED_TOKEN string = "token is expired"
const INVALID_SIGNING_METHOD string = "invalid signing method"
const INVALID_TOKEN string = "invalid token"

// GORM errors
const ALREADY_EXISTS string = "already exists"
const RECORD_ALREADY_EXISTS string = "'%s' already exists"
const RECORD_NOT_FOUND string = "record not found"

// Test
const TEST_ALREADY_EXISTS string = ALREADY_EXISTS
const TEST_CREATED string = "created"
const TEST_DELETED string = "deleted"
const TEST_FOUND string = "found"
const TEST_NOT_FOUND string = "not found"
const TEST_UPDATED string = "updated"
