class FoodChain {
    private String[] lyrics = {
            "I don't know why she swallowed the fly. Perhaps she'll die.",
            "It wriggled and jiggled and tickled inside her",
            "How absurd to swallow a bird!",
            "Imagine that, to swallow a cat!",
            "What a hog, to swallow a dog!",
            "I don't know how she swallowed a cow!",
            "Just opened her throat and swallowed a goat!",
            "I don't know how she swallowed a cow!",
            "She's dead, of course!",
    };

    private String start = "I know an old lady who swallowed a %s.";
    private String chain = "She swallowed the %s to catch the %s.";
    private String end = "She's dead, of course!";
    private String[] animals = { "fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse" };

    String verse(int verseNumber) {
        String output = String.format(start, animals[verseNumber - 1]);

        if (verseNumber > 1) {
            for (int i = 0; i <= verseNumber - 1; i ++) {
                output += lyrics[i];
                String animal1 = animals[i];
                String animal2 = animals[i + 1];


            }
        }
        return String.format(start, animal1) + "\n" + lyrics[verseNumber - 1];
    };

    String verses(int startVerse, int endVerse) {
        return "";
    }
}