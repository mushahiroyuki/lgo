public class Client {
    private final Logic logic;
    // この型はインタフェースで、実装ではない

    public Client(Logic logic) {
        this.logic = logic;
    }

    public void program() {
      // どこからかデータを取得
      String data = "";
      this.logic.process(data);
    }
}
