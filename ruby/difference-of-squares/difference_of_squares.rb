class Squares
  attr_reader :number
  def initialize(number)
    @number = number
  end

  def square_of_sum
    (0..number).step(1).sum ** 2
  end

  def sum_of_squares
  end
end
