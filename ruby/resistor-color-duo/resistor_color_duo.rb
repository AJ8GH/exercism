class ResistorColorDuo
  BANDS = %w[black brown red orange yellow green blue violet grey white]

  def self.value(colors)
    duo = colors.take(2)
    Integer [BANDS.index(duo.first), BANDS.index(duo.last)].join
  end
end

# Hash method using iteration
# class ResistorColorDuo
#   BANDS = {
#      black: 0,
#      brown: 1,
#      red: 2,
#      orange: 3,
#      yellow: 4,
#      green: 5,
#      blue: 6,
#      violet: 7,
#      grey: 8,
#      white: 9
#    }
#   def self.value(colors)
#   Integer colors.take(2).map(&:to_sym).map(&BANDS).join
#   end
# end
