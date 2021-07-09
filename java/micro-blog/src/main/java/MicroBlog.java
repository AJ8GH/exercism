class MicroBlog {

    public String truncate(String input) {
        if (Character.codePointCount(input, 0, input.length()) <= 5) {
            return input;
        }

        StringBuilder result = new StringBuilder(5);
        input.codePoints().limit(5).forEach(result::appendCodePoint);
        return result.toString();
    }
}