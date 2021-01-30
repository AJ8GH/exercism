class Series
  attr_accessor :string

  def initialize(string)
    @string = string
  end

  def slices(number)
    slices = []
    (string.length - number + 1).times do |index|
      slices << string[index, number]
    end
    slices
  end
end
