[Test]
public void EventFiringManual()
{
    bool loadFired = false;
    SomeView view = new SomeView();
    view.Load += delegate { loadedFired = true };
    view.DoSomethingThatEventuallyFiresThisEvent();

    Assert.IsTrue(loadFired);
}