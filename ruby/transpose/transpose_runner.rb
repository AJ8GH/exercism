require_relative 'transpose'

input = "The fourth line.\nThe fifth line."
output = "TT\nhh\nee\n  \nff\noi\nuf\nrt\nth\nh \n l\nli\nin\nne\ne.\n."

puts input
puts
# puts output
puts

p Transpose.transpose(input)
p output

puts

input2 = "The longest line.\nA long line.\nA longer line.\nA line."
output2 = "TAAA\nh   \nelll\n ooi\nlnnn\nogge\nn e.\nglr\nei \nsnl\ntei\n .n\nl e\ni .\nn\ne\n."

p input2
puts

puts Transpose.transpose(input2)
puts output2
