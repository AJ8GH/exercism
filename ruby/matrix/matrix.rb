class Matrix
  attr_accessor :rows, :string, :columns
  def initialize(string)
    @string = string
    @rows = get_rows
    @columns = rows.transpose
  end

  def get_rows
    string.split("\n").map { |row| row.split.map(&:to_i) }
  end
end
