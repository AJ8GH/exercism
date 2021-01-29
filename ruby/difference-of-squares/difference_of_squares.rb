class Squares
  attr_reader :number
  def initialize(number)
    @number = number
  end

  def square_of_sum
    natural_numbers.sum ** 2
  end

  def sum_of_squares
    natural_numbers.map { |number| number ** 2 }.sum
  end

  def natural_numbers
    (0..number).step
  end

  def difference
    square_of_sum - sum_of_squares
  end
end
