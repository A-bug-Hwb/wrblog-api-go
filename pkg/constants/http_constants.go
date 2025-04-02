package constants

import "net/http"

// SUCCESS 200 成功
const SUCCESS = http.StatusOK

const NOT_FOUND = http.StatusNotFound

// BAD_REQUEST 400 参数有误
const BAD_REQUEST = http.StatusBadRequest

// UNAUTHORIZED 401 未授权
const UNAUTHORIZED = http.StatusUnauthorized

// FORBIDDEN 403 授权过期
const FORBIDDEN = http.StatusForbidden

// IP_LOCKED 423 ip锁定，禁止访问
const IP_LOCKED = http.StatusLocked

// ERROR 500 系统错误
const ERROR = http.StatusInternalServerError

// WARN 601 系统警告
const WARN = 601
