class Phrase
  def initialize(phrase)
    @phrase = phrase
    @words = get_words
  end

  def word_count
    words.inject(Hash.new(0)) do |hash, word|
      hash[format_word(word)] += 1
      hash
    end
  end

  private

  attr_reader :phrase, :words

  def get_words
    phrase.split.map { |word| word.split(',') }.flatten
  end

  def format_word(word)
    word.gsub(/[^a-zA-Z0-9']/, '').downcase.gsub(/\B'\b|\b'\B/, '')
  end
end
