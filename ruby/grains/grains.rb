class Grains
  @@board = 1..64

  def self.square(number)
    validate(number)
    2 ** (number - 1)
  end

  def self.total
    @@board.map { |number| square(number) }.sum
  end

  def self.validate(number)
    raise ArgumentError unless @@board.include?(number)
  end
end
