class Squares
  attr_reader :number, :natural_numbers
  def initialize(number)
    @number = number
    @natural_numbers = (0..number).step
  end

  def square_of_sum
    natural_numbers.sum ** 2
  end

  def sum_of_squares
    natural_numbers.map { |number| number ** 2 }.sum
  end

  def difference
    square_of_sum - sum_of_squares
  end
end
