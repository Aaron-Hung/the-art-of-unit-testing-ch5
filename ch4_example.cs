public interface IWebService
{
    void LogError(string message, int a, int b, int c, int d);
}

public class FakeWebService: IWebService
{
    public string LastError;
    public int Log_a, Log_b, Log_c, Log_d;
    pubic void LogError(string message, int a, int b, int c, int d)
    {
        LastError = message;
        Log_a = a;
        Log_b = b;
        Log_c = c;
        Log_d = d;
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
    StringAssert.Contains(1, mockService.a); // 針對模擬物件進行驗證
    StringAssert.Contains(2, mockService.b); // 針對模擬物件進行驗證
    StringAssert.Contains(3, mockService.c); // 針對模擬物件進行驗證
    StringAssert.Contains(4, mockService.d); // 針對模擬物件進行驗證

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
            int a = 1;
            int b = 2;
            int c = 3;
            int d = 4;
            service.LogError("FileName too short: " + fileName, a, b, c, d); // 在產品程式中紀錄檔名過短的錯誤
        }
    }
}