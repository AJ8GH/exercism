class ResistorColorDuo
  BANDS = %w[black brown red orange yellow green blue violet grey white]

  def self.value(colors)
    duo = colors.take(2)
    Integer [BANDS.index(duo.first), BANDS.index(duo.last)].join
  end
end
