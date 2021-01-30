class Triangle
  attr_reader :sides, :valid_triangle, :all_sides_positive, :triangle_inequality

  def initialize(sides)
    @sides = sides
    @all_sides_positive = sides.all? { |side| side > 0 }
    @triangle_inequality = sides.all? { |side| sides.sum - side >= side }
    @valid_triangle = all_sides_positive && triangle_inequality
  end

  def equilateral?
    sides.all?(sides.first) and valid_triangle
  end

  def isosceles?
    valid_triangle and sides.any? { |side| sides.count(side) >= 2 }
  end

  def scalene?
    valid_triangle and not equilateral? and not isosceles?
  end
end
