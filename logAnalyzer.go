package ch5

type LogAnalyzer struct {
	service IWebService
}

func (la *LogAnalyzer) Analyze(filename string){
	if(len(filename) < 8){
		la.service.LogError("File name too short: " + filename)
	}
}

func NewLogAnalyzer(service IWebService) *LogAnalyzer {
    return &LogAnalyzer{
        service: service,
    }
}


