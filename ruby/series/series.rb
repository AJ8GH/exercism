class Series
  private

  attr_reader :number_string, :number_string_length
  def initialize(number_string)
    @number_string = number_string
    @number_string_length = number_string.length
  end

  def argument_length_valid(sub_string_length)
    raise ArgumentError if sub_string_length > number_string_length
  end

  public

  def slices(sub_string_length)
    argument_length_valid(sub_string_length)
    difference = number_string_length - sub_string_length
    sub_strings = []
    (difference + 1).times do |start_of_sub_string|
      sub_strings << number_string[start_of_sub_string, sub_string_length]
    end
    sub_strings
  end
end
