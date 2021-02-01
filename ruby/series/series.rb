class Series
  private

  attr_reader :digits

  def initialize(numbers)
    @digits = numbers.chars
  end

  def valid?(size)
    digits.length >= size
  end

  public

  def slices(size)
    raise SliceSizeError unless valid?(size)
    digits.each_cons(size).map(&:join)
  end
end

class SliceSizeError < ArgumentError
  def initialize
    raise self, 'Slice requested must be smaller than series.'
  end
end
