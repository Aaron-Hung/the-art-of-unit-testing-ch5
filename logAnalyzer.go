package ch5

type LogAnalyzer struct {
	webService IWebService
	emailService IEmailService
}

func (la *LogAnalyzer) Analyze(filename string){
	if(len(filename) < 8){
		err := la.webService.LogError("File name too short: " + filename)
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
