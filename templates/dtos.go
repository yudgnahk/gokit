package templates

// BaseDtos ...
const BaseDtos = `package dtos

// Meta is common meta.
type Meta struct {
	Code    int    BACK_STICKjson:"code"BACK_STICK
	Message string BACK_STICKjson:"message"BACK_STICK
}
`

// HealthDtos ...
const HealthDtos = `package dtos

// HealthResponse represent body response of health check API.
type HealthResponse struct {
	Meta Meta BACK_STICKjson:"meta"BACK_STICK
}
`
