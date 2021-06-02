class Phrase
  def initialize(phrase)
    @phrase = phrase
    @words = normalize_phrase
  end

  def word_count
    words.tally
  end

  private

  attr_reader :phrase, :words

  def normalize_phrase
    phrase.split(/\s|,/).join(' ').split.map do |word|
      word.scan(/\b[\w']+\b/).join.downcase
    end
  end
end
