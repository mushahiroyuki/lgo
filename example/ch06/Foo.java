//
// 1. javac Foo.java
// 2. java Foo

public class Foo {
    private int x;

    public Foo(int x) {
        this.x = x;
    }

    public static void main(String[] args) {
        outer();
    }

    private static void outer() {
        Foo f = new Foo(10);
        inner1(f);
        System.out.println(f.x);
        inner2(f);
        System.out.println(f.x);
        Foo g = null;
        inner2(g);
        System.out.println(g == null);
    }

    private static void inner1(Foo f) {
        f.x = 20;
    }

    private static void inner2(Foo f) {
        f = new Foo(30);
    }
}
