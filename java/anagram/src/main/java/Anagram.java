import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Locale;

class Anagram {
    private final String word;
    private final String normalizedWord;

    Anagram(String word) {
        this.word = word;
        this.normalizedWord = normalize(word);
    }

    public List<String> match(List<String> args) {
        List<String> matches = new ArrayList<>();

        for (String wordToCheck: args) {
            String normalizedWordToCheck = normalize(wordToCheck);
            if (!isSameWord(wordToCheck) && normalizedWordToCheck.equals(normalizedWord)) {
                matches.add(wordToCheck);
            }
        }
        return matches;
    }

    private String normalize(String word) {
        char[] lowerCaseChars = word.toLowerCase(Locale.ROOT).toCharArray();
        Arrays.sort(lowerCaseChars);
        return new String(lowerCaseChars);
    }

    private boolean isSameWord(String wordToCheck) {
        return wordToCheck.toLowerCase(Locale.ROOT).equals(word.toLowerCase(Locale.ROOT));
    }
}