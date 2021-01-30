class Matrix
  attr_accessor :rows, :columns
  def initialize(matrix_string)
    @rows = matrix_string.lines.map { |row| row.split.map(&:to_i) }
    @columns = rows.transpose
  end
end
