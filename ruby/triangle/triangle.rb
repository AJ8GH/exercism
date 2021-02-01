class Triangle
  attr_reader :sides,
              :no_zero_sides,
              :triangle_inequality,
              :valid_triangle

  alias :valid_triangle? :valid_triangle
  undef :valid_triangle

  def initialize(sides)
    @sides = sides.map(&:abs)
    @no_zero_sides = sides.none? { |side| side.zero? }
    @triangle_inequality = self.sides.all? { |side| self.sides.sum - side >= side }
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

if defined? Minitest
  describe Triangle do
    describe 'it takes negative sides as arguments' do
      it 'confirms all negative equilateral' do
        triangle = Triangle.new([-2, -2, -2])
        assert triangle.equilateral?, 'Expected true'
      end

      it 'confirms partially negative equilateral' do
        triangle = Triangle.new([-2, 2, -2])
        assert triangle.equilateral?, 'Expected true'
      end

      it 'confirms isosceles with one negative side' do
        triangle = Triangle.new([5, -3, 3])
        assert triangle.isosceles?, 'Expected true'
      end

      it 'confirms scalene with 2 negative sides' do
        triangle = Triangle.new([-5, 6, -8])
        assert triangle.scalene?, 'Expected true'
      end

      it 'rejects invalid triangle with a negative side' do
        triangle = Triangle.new([-5, 0, 3])
        refute triangle.scalene?, 'Expected false'
      end

      it 'rejects inequal triangle with a negative side' do
        triangle = Triangle.new([-5, 1, 1])
        refute triangle.isosceles?, 'Expected false'
      end
    end
  end
end
