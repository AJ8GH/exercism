import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class TwelveDays {
    private final String verseStart = "On the %s day of Christmas my true love gave to me: ";
    private final String[] verseNumbers = { "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth" };
    private final String[] verseEnds = {
            "twelve Drummers Drumming, ",
            "eleven Pipers Piping, ",
            "ten Lords-a-Leaping, ",
            "nine Ladies Dancing, ",
            "eight Maids-a-Milking, ",
            "seven Swans-a-Swimming, ",
            "six Geese-a-Laying, ",
            "five Gold Rings, ",
            "four Calling Birds, ",
            "three French Hens, ",
            "two Turtle Doves, and ",
            "a Partridge in a Pear Tree."
    };

    String verse(int verseNumber) {
        String[] verses = Arrays.copyOfRange(
                verseEnds,
                verseEnds.length - verseNumber,
                verseEnds.length
        );

        String end = String.join("", verses) + "\n";
        return String.format(verseStart, verseNumbers[verseNumber - 1]) + end;
    }

    String verses(int startVerse, int endVerse) {
        List<String> output = new ArrayList<>();
        for (int i = startVerse; i <= endVerse; i++) {
            output.add(verse(i));
        }
        return String.join("\n", output);
    }

    String sing() {
        return verses(1, 12);
    }
}