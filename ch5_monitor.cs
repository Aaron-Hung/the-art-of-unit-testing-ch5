class Presenter
{
    public readonly IView _view;
    public Presenter(IView view)
    {
        _view = view;
        this._view.Loaded += OnLoaded;
    }
    private void OnLoaded()
    {
        _view.Render("Hello World");
    }
}

public interface IView
{
    event Action Loaded;
    void Render(string text);
}

// ----- TESTS
[TestFixture]
public class EventRelatedTests
{
    [Test]
    public void ctor_WhenViewIsLoaded_CallsViewRender()
    {
        var mockView = Substitute.For<IView>();
        Presenter p = new Presenter(mockView);
        mockView.Loaded += Raise.Event<Action>(); // 使用 NSubstitute 觸發事件
        mockView.Received().Render(Arg.Is<string>(s => s.Contains("Hello World"))); // 驗證在測試中 view 物件的 Render 方法是否被呼叫
    }
}