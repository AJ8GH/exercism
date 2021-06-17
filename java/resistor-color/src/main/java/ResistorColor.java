import java.util.Arrays;

class ResistorColor {
    private final String[] COLOR_CODES = new String[] {
            "black",
            "brown",
            "red",
            "orange",
            "yellow",
            "green",
            "blue",
            "violet",
            "grey",
            "white"
    };

    int colorCode(String color) {
         return Arrays.asList(COLOR_CODES).indexOf(color);
    }

    String[] colors() {
        return COLOR_CODES;
    }
}
