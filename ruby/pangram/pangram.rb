class Pangram
  ALPHA = ('a'..'z')

  def self.pangram?(sentence)
    letters = sentence.downcase.gsub(/[^a-zA-Z]/, '')
    letters.chars.sort.uniq.join == ALPHA.to_a.join
  end
end
