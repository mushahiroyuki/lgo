class Foo:
    def __init__(self, x):
        self.x = x

def outer():
    f = Foo(10)
    inner1(f)
    print(f.x)  ## 20
    inner2(f)
    print(f.x)  ## これも20
    g = None
    inner2(g)
    print(g is None)  ## True

def inner1(f):
    f.x = 20

def inner2(f):
    f = Foo(30)

outer()
