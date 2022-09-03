public interface IWebService
{
    void LogError(string message);
}

public class FakeWebService: IWebService
{
    public string LastError;
    pubic void LogError(string message)
    {
        LastError = message;
    }
}

    [Test]
public void Analyze_TooShortFileName_CallsWebService()
{
    FakeWebService mockService = new FakeWebService();

    LogAnalyzer log = new LogAnalyzer(mockService);
    string tooShortFileName = "abc.exe";
    log.Analyze(tooShortFileName);

    StringAssert.Contains("FileName too short: abc.exe", mockService.LastError); // 針對模擬物件進行驗證
}

public class LogAnalyzer
{
    private IWebService service;
    public LogAnalyzer(IWebService service)
    {
        this.service = service;
    }
    public void Analyze(string fileName)
    {
        if (fileName.Length < 8)
        {
            service.LogError("FileName too short: " + fileName); // 在產品程式中紀錄檔名過短的錯誤
        }
    }
}