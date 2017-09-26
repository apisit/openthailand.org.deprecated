package constant

import "time"

const (
	AUTHENTICATED_USER_HEADER       = "OpenThailand-Authenticated-User"
	AUTHENTICATED_USER_TOKEN_HEADER = "OpenThailand-Authenticated-User-Access-Token"
)
const (
	COOKIE_AGE_MINUTE     = time.Duration(24 * 60)
	WEB_TOKEN_COOKIE_NAME = "cookie_token"
)

const (
	HTTP_HEADER_KEY_REST_API               = "X-OpenThailand-Rest-API-Key"
	HTTP_HEADER_KEY_APPLICATION_ID         = "X-OpenThailand-Application-Id"
	HTTP_HEADER_AUTHORIZATION_TOKEN_PREFIX = "OpenThailand"
)

const (
	CACHE_KEY_PROVINCES = "PROVINCES"
)
