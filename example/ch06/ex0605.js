//
// $ node ex0602.js   
// （node.jsのインストールが必要です）

class Foo {
    constructor(x) {
        this.x = x;
    }
}

function outer() {
    let f = new Foo(10);
    inner1(f);
    console.log(f.x);
    inner2(f);
    console.log(f.x);
    let g = null;
    inner2(g);
    console.log(g == null);
}

function inner1(f) {
    f.x = 20;
}

function inner2(f) {
    f = new Foo(30);
}

outer();
