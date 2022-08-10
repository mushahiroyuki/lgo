# coding: utf-8
##
## $ ruby  ex0605.rb

class Foo
  def initialize(x)
    @x = x
  end

  attr_accessor :x

end

def outer
  f = Foo.new(10)
  inner1(f)
  puts f.x
  inner2(f)
  puts f.x
  g = nil
  inner2(g)
  puts g.nil?
end

def inner1(f)
  f.x = 20
end

def inner2(f)
  f = Foo.new(30)
end

outer
