class Proverb {
    private final String[] words;
    private final String part1 = "For want of a ";
    private final String part2 = " the ";
    private final String part3 = " was lost.\n";
    private final String end = "And all for the want of a ";
    private final StringBuilder output;

    Proverb(String[] words) {
        this.words = words;
        this.output = new StringBuilder();
    }

    String recite() {
        if (words.length == 0) return output.toString();

        String endProverb = end + words[0] + ".";
        if (words.length == 1) return endProverb;

        createProverbs();
        return output.append(endProverb).toString();
    }

    private void createProverbs() {
        for (int i = 0; i < words.length - 1; i++) {
            output
                    .append(part1).append(words[i])
                    .append(part2).append(words[i + 1]).append(part3);
        }
    }
}