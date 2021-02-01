class Triangle
  attr_reader :sides,
              :no_zero_sides,
              :triangle_inequality,
              :valid_triangle

  alias :valid_triangle? :valid_triangle
  undef :valid_triangle

  def initialize(sides)
    @sides = sides
    @no_zero_sides = sides.none? { |side| side.zero? }
    @triangle_inequality = sides.all? { |side| sides.sum - side >= side }
    @valid_triangle = no_zero_sides && triangle_inequality
  end

  def equilateral?
    valid_triangle? and sides.all?(sides.first)
  end

  def isosceles?
    valid_triangle? and sides.any? { |side| sides.count(side) >= 2 }
  end

  def scalene?
    valid_triangle? and not equilateral? and not isosceles?
  end
end

# if defined? Minitest
#   describe ''
# end
