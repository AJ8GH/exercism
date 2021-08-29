public class BeerSong {
    private final String verseStart = " bottles of beer on the wall, ";
    private final String verseMiddle = " bottles of beer.\nTake one down and pass it around, ";
    private final String verseEnd = " bottles of beer on the wall.\n\n";
    private final String topUp = "Go to the store and buy some more, ";

    public String sing(int bottles, int verses) {
        return buildVerse(bottles);
    }

    private String buildVerse(int bottles) {
        String start = bottles + verseStart;
        String middle = bottles + verseMiddle;
        String end = (bottles - 1) + verseEnd;

        if (bottles == 0) {
            return "No more " + verseStart + "no more" + topUp + (bottles + 99) + verseEnd;
        } else if (bottles == 1) {
            return (start + middle)
                    .replaceAll("bottles", "bottle") +
                    end.replaceAll("0", "no more");
        } else if (bottles == 2) {
            return start + middle + end
                    .replaceAll("bottles", "bottle")
                    .replaceAll("one", "it");
        }
        return start + middle + end;
    }

    public String singSong() {
        return "";
    }
}