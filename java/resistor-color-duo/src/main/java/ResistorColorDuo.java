import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

class ResistorColorDuo {
    private final String[] COLORS = new String[] {
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

    int value(String[] colors) {
        List<String> intList = Arrays.stream(colors)
                .map(color -> String.valueOf(Arrays.asList(COLORS)
                        .indexOf(color))).collect(Collectors.toList());

        return Integer.parseInt(String.join("", intList)
                .substring(0, 2));
    }
}
