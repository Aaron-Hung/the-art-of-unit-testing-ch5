package ch5

import "ch5/er"

type LogAnalyzer struct {
	webService IWebService
	emailService IEmailService
}

func (la *LogAnalyzer) Analyze(filename string){
	if(len(filename) < 8){
		errorInfo := er.NewErrorInfo("File name too short", filename)
		err := la.webService.LogError(errorInfo)
		if err != nil {
			la.emailService.SendEmail("support@going.cloud","Fake exception!", "Can't Log")
		}
	}
}

func NewLogAnalyzer(fs IWebService, fe IEmailService) *LogAnalyzer {
    return &LogAnalyzer{
        webService: fs,
		emailService: fe,
    }
}
