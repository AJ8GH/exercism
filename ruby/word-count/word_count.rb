class Phrase
  def word_count
    words.tally
  end

  private

  attr_reader :phrase, :words

  def initialize(phrase)
    @phrase = phrase
    @words = normalize_phrase
  end

  def normalize_phrase
    create_word_array.map { |word| normalize_word(word) }
  end

  def create_word_array
    phrase.split.map { |word| word.split(',') }.flatten
  end

  def normalize_word(word)
    word.gsub(/^'|'$|[^\w']/, '').downcase
  end
end
