class SumOfMultiples
  attr_reader :factors
  def initialize(*factors)
    @factors = factors
  end

  def to(number)
    multiples = factors.map do |factor|
      (1...number).select { |divisor| divisor.multiple?(factor) }
    end
    multiples.flatten.uniq.sum
  end
end

class Integer
  def multiple?(factor)
    self % factor == 0
  end
end
