import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Locale;

class Scrabble {
    private final String word;
    private final HashMap<List<Character>, Integer> SCORES = createScoresMap();

    Scrabble(String word) {
        this.word = word;
    }

    int getScore() {
        char[] chars = word.toUpperCase(Locale.ROOT).toCharArray();
        int score = 0;
        for (char c : chars) {
            for (List<Character> letters : SCORES.keySet()) {
                if(letters.contains(c)) score += SCORES.get(letters);
            }
        }
        return score;
    }

    private HashMap<List<Character>, Integer> createScoresMap() {
        HashMap<List<Character>, Integer> scoresMap = new HashMap<>();
        scoresMap.put(Arrays.asList('A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'), 1);
        scoresMap.put(Arrays.asList('D', 'G'), 2);
        scoresMap.put(Arrays.asList('B', 'C', 'M', 'P'), 3);
        scoresMap.put(Arrays.asList('F', 'H', 'V', 'W', 'Y'), 4);
        scoresMap.put(Arrays.asList('K'), 5);
        scoresMap.put(Arrays.asList('J', 'X'), 8);
        scoresMap.put(Arrays.asList('Q', 'Z'), 10);
        return scoresMap;
    }
}