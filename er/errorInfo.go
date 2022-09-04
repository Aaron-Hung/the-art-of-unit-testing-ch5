package er

type ErrorInfo struct{
	Message string
	Filename string
}

func NewErrorInfo(message string, filename string) *ErrorInfo {
    return &ErrorInfo{
        Message: message,
		Filename: filename,
    }
}