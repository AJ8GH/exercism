import java.util.*;

class Scrabble {
    private final String word;
    private final HashMap<Set<Character>, Integer> SCORES = createScoresMap();

    Scrabble(String word) {
        this.word = word;
    }

    int getScore() {
        char[] chars = word.toUpperCase(Locale.ROOT).toCharArray();
        int score = 0;
        for (char c : chars) {
            for (Set<Character> letters : SCORES.keySet()) {
                if(letters.contains(c)) score += SCORES.get(letters);
            }
        }
        return score;
    }

    private HashMap<Set<Character>, Integer> createScoresMap() {
        HashMap<Set<Character>, Integer> scoresMap = new HashMap<>();
        scoresMap.put(new HashSet<>(Arrays.asList('A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T')), 1);
        scoresMap.put(new HashSet<>(Arrays.asList('D', 'G')), 2);
        scoresMap.put(new HashSet<>(Arrays.asList('B', 'C', 'M', 'P')), 3);
        scoresMap.put(new HashSet<>(Arrays.asList('F', 'H', 'V', 'W', 'Y')), 4);
        scoresMap.put(new HashSet<>(Arrays.asList('K')), 5);
        scoresMap.put(new HashSet<>(Arrays.asList('J', 'X')), 8);
        scoresMap.put(new HashSet<>(Arrays.asList('Q', 'Z')), 10);
        return scoresMap;
    }
}