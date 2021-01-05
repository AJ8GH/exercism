class Darts
  def initialize (x, y)
    @x = x
    @y = y
  end

  def score
  r2 = @x ** 2 + @y ** 2
  r2 > 100 ? 0 :
  r2 > 25 ? 1 :
  r2 > 1 ? 5 : 10
  end
end
