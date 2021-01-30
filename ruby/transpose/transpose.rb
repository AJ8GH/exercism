class Transpose
  def self.transpose(matrix_string)
    matrix_string.split("\n").map(&:chars).transpose
  end
end

class String
  def number_of_lines
    self.lines.count
  end
end
