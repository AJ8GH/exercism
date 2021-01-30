class Series
  attr_accessor :string

  def initialize(string)
    @string = string
  end

  def slices(number)
    @string.split('')
  end
end
