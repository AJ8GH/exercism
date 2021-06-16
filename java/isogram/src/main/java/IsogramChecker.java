import java.util.*;

class IsogramChecker {
    boolean isIsogram(String phrase) {
        char[] characters = normalize(phrase).toCharArray();
        Set<Character> set = new HashSet<>();
        for (char character : characters) {
            if (!set.add(character)) { return false; }
        }
        return true;
    }

    private String normalize(String phrase) {
        return phrase.replaceAll("[^a-zA-Z]", "").toLowerCase(Locale.ROOT);
    }
}
