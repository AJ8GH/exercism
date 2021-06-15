import java.util.regex.Matcher;
import java.util.regex.Pattern;

class Acronym {
    private final String phrase;

    Acronym(String phrase) {
        this.phrase = normalize(phrase);
    }

    String get() {
        Pattern pattern = Pattern.compile("\\b[a-zA-Z]");
        Matcher matcher = pattern.matcher(phrase);
        StringBuilder acronym = new StringBuilder();
        while (matcher.find()) { acronym.append(matcher.group().toUpperCase()); }
        return acronym.toString();
    }

    private String normalize(String phrase) {
        return phrase.replaceAll("'", "")
                .replaceAll("[-|_]", " ");
    }
}
