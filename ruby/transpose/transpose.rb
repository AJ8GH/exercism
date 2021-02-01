require 'pry'

class Transpose
  def self.transpose(matrix_string)
    output = "TT\nhh\nee\n  \nff\noi\nuf\nrt\nth\nh \n l\nli\nin\nne\ne.\n."

    lines = matrix_string.split("\n")
    chars = lines.map(&:chars)

    slices = chars.map { |array| array.each_slice(lines.length).to_a }
    slices


    binding.pry

    max_slice = max_slice(slices)
    padded = pad_slices(slices)

    columns = padded.map(&:transpose)
    binding.pry

    join_cols = columns.map { |ary_of_cols| ary_of_cols.map(&:join) }

    result = join_cols.map { |col| col.join("\n") }.join("\n")
  end
end

def max_slice(slices)
  slices_1, slices_2 = slices
  [slices_1.map(&:size).max, slices_2.map(&:size).max].max
end

def pad_slices(slices)
  max_slice = max_slice(slices)
  slices_1, slices_2 = slices
  slices.map do |ary_of_slices|
    ary_of_slices.map do |slice|
      if slice.length == max_slice
        slice
      else
        slice.unshift(' ' * (max_slice - slice.length))
      end
    end
  end
end

class String
  def number_of_lines
    self.lines.count
  end
end

["The first line. ", "The second line."]
