import java.util.*;

public class PangramChecker {
    private final String ALPHABET = "abcdefghijklmnopqrstuvwxyz";

    public boolean isPangram(String input) {
        String normalizedInput = normalize(input);
        String pangramCheck = createPangramCheck(normalizedInput);
        return pangramCheck.equals(ALPHABET);
    }

    private String normalize(String input) {
        String lowerCaseInput = input.toLowerCase(Locale.ROOT);
        return lowerCaseInput.replaceAll("[^a-zA-Z]", "");
    }

    private String createPangramCheck(String normalizedInput) {
        List<String> letters = Arrays.asList(normalizedInput.split(""));
        List<String> uniqueLetters = new ArrayList<>(new HashSet<>(letters));
        Collections.sort(uniqueLetters);
        return String.join("", uniqueLetters);
    }
}