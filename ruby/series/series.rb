class Series

  private

  attr_reader :numbers, :numbers_length
  def initialize(numbers)
    @numbers = numbers
    @numbers_length = numbers.length
  end

  def slice_greater_than_length(slice_size)
    raise ArgumentError if slice_size > numbers_length
  end

  public

  def slices(slice_size)
    slice_greater_than_length(slice_size)
    slices = []
    (numbers_length - slice_size + 1).times do |index|
      slices << numbers[index, slice_size]
    end
    slices
  end
end
