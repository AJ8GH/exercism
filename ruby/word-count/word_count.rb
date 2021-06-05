class Phrase
  def initialize(phrase)
    @phrase = phrase
    @words = normalize
  end

  def word_count
    words.tally
  end

  private

  attr_reader :phrase, :words

  def normalize
    phrase.downcase.scan(/\b[\w']+\b/)
  end
end
