class Triangle
  attr_reader :sides,
              :all_sides_positive,
              :triangle_inequality,
              :valid_triangle

  alias :valid_triangle? :valid_triangle
  undef :valid_triangle

  def initialize(sides)
    @sides = sides
    @all_sides_positive = sides.all? { |side| side > 0 }
    @triangle_inequality = sides.all? { |side| sides.sum - side >= side }
    @valid_triangle = all_sides_positive && triangle_inequality
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
