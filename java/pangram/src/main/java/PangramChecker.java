import java.util.*;

public class PangramChecker {
    private String alphabet = "abcdefghijklmnopqrstuvwxyz";

    public boolean isPangram(String input) {
        String filteredInput = input.toLowerCase(Locale.ROOT).replaceAll("[^a-zA-Z]", "");
        List<String> letters = Arrays.asList(filteredInput.split(""));
        List<String> uniqueLetters = new ArrayList<>(new HashSet<>(letters));
        Collections.sort(uniqueLetters);
        String pangramCheck = String.join("", uniqueLetters);
        return pangramCheck.equals(alphabet);
    }
}
