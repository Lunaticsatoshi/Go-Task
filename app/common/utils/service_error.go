package utils

type ServiceError struct {
	Code                 int
	Message              string
	InternalErrorMessage string
	Payload              string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func (e *ServiceError) LogMessage() string {
	return "Error: " + e.Message + " Internal Error: " + e.InternalErrorMessage + " Payload: " + e.Payload + " Code: " + ConvertIntToString(e.Code)
}
