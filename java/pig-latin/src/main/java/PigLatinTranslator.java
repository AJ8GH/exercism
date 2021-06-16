import java.util.*;

class PigLatinTranslator {
    private final char[] VOWELS = {'a', 'e', 'i', 'o', 'u', 'y'};

    public String translate(String phrase) {
        String[] words = phrase.split(" ");

        List<String> translation = new ArrayList<>();

        for (String word: words) {
            int firstVowel = findFirstVowel(word);

            if (firstVowel == 0 || word.startsWith("xr")) {
                translation.add(word + "ay");
            } else if (word.startsWith("qu")) {
                translation.add(word.substring(2) + "quay");
            } else if (word.startsWith("squ")) {
                translation.add(word.substring(3) + "squay");
            } else {
                translation.add(
                        word.substring(firstVowel) +
                        word.substring(0, firstVowel) +
                        "ay"
                );
            }
        }
        return String.join(" ", translation);
    }

    private int findFirstVowel(String word) {
        List<Integer> vowelIndices = new ArrayList<>();
        for (char vowel : VOWELS) {
            if (word.indexOf(vowel) != -1) {
                vowelIndices.add(word.indexOf(vowel));
            }
        }
        Collections.sort(vowelIndices);
        if (word.startsWith("y") && vowelIndices.get(1) == 1) {
            return vowelIndices.get(1);
        }
        return vowelIndices.get(0);
    }
}