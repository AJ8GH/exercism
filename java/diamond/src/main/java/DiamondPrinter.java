import java.util.ArrayList;
import java.util.List;

class DiamondPrinter {
    private final char[] ALPHABET = {
            'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'
        };

    List<String> printToList(char a) {
        int index = Character.getNumericValue(a) - 10;

        List<String> output = new ArrayList<>();

        for (int i = 0; i <= index; i++) {
        }

        output.add(String.valueOf(a));

        return output;
    }
}
