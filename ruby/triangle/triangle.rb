class Triangle
  attr_reader :sides
  def initialize(sides)
    @sides = sides
  end

  def equilateral?
    sides.all?(sides[0]) && valid_triangle?
  end

  def isosceles?
    return true if equilateral?
    sides.map { |side| sides.count(side) }.include?(2) && valid_triangle?
  end

  def scalene?
    !equilateral? && !isosceles? && valid_triangle?
  end

  private

  def valid_triangle?
    all_sides_positive? && triangle_inequality?
  end

  def all_sides_positive?
    sides.all? { |side| side > 0 }
  end

  def triangle_inequality?
    sides.map { |side| sides.sum - side >= side }.all?(true)
  end
end
