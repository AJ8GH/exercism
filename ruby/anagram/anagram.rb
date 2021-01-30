class Anagram
  private

  attr_reader :word, :sorted_word, :downcase_word
  def initialize(word)
    @word = word
    @downcase_word = word.downcase
    @sorted_word = word.sort
  end

  def anagram?(anagram)
    downcase_word != anagram.downcase and sorted_word == anagram.sort
  end

  public

  def match(anagrams)
    anagrams.select { |candidate_anagram| anagram?(candidate_anagram) }
  end
end

class String
  def sort
    self.downcase.chars.sort.join
  end
end
