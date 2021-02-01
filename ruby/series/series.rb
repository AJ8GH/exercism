class Series
  private

  attr_reader :numbers

  def initialize(numbers)
    @numbers = numbers
  end

  def valid?(size)
    numbers.length >= size
  end

  public

  def slices(size)
    raise SliceSizeError unless valid?(size)
    numbers.chars.each_cons(size).map(&:join)
  end
end

class SliceSizeError < ArgumentError
end
