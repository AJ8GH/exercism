import java.util.Locale;

public class Bob {
    public String hey(String speech) {
        if (isEmpty(speech)) return "Fine. Be that way!";
        if (isQuestion(speech) && isYelling(speech)) {
            return "Calm down, I know what I'm doing!";
        }
        if (isQuestion(speech)) return "Sure.";
        if (isYelling(speech)) return "Whoa, chill out!";
        return "Whatever.";
    }

    private boolean isYelling(String speech) {
        String yelledSpeech = speech.toUpperCase(Locale.ROOT);
        String spokenSpeech = speech.toLowerCase(Locale.ROOT);
        return speech.equals(yelledSpeech) && !speech.equals(spokenSpeech);
    }

    private boolean isQuestion(String speech) {
        return speech.replaceAll("\\s+","").endsWith("?");
    }

    private boolean isEmpty(String speech) {
        return speech.replaceAll("\\s+","").length() == 0;
    }
}
