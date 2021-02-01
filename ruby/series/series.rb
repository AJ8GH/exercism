class Series
  private

  attr_reader :numbers, :length, :digits

  def initialize(numbers)
    @numbers = numbers
    @length = numbers.length
    @digits = numbers.chars
  end

  def valid?(size)
    length >= size
  end

  public

  def slices(size)
    raise ArgumentError unless valid?(size)
    digits.map.with_index do |digit, i|
      numbers.slice(i, size) unless i > length - size
    end.compact
  end
end
