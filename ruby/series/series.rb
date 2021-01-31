class Series
  private

  attr_reader :numbers, :numbers_length, :digits

  def initialize(numbers)
    @numbers = numbers
    @numbers_length = numbers.length
    @digits = numbers.chars
  end

  def validate(sub_string)
    raise ArgumentError if sub_string > numbers_length
  end

  public

  def slices(sub_string)
    validate(sub_string)
    digits.map.with_index do |digit, index|
      numbers.slice(index, sub_string) unless index > numbers_length - sub_string
    end.compact
  end
end
