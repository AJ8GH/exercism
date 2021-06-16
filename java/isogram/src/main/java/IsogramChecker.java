import java.util.*;

class IsogramChecker {

    boolean isIsogram(String phrase) {
        String letters = phrase.replaceAll("[^a-zA-Z]", "");
        char[] characters = letters.toLowerCase(Locale.ROOT).toCharArray();
        Set<Character> set = new HashSet<>();

        for (char character : characters) {
            if (!set.add(character)) { return false; }
        }
        return true;
    }
}
