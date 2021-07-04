import java.util.Locale;

public class Bob {
    private final String DEFAULT_RESPONSE = "Whatever.";
    private final String EMPTY_RESPONSE = "Fine. Be that way!";
    private final String QUESTION_RESPONSE = "Sure.";
    private final String YELLING_RESPONSE = "Whoa, chill out!";
    private final String YELLING_QUESTION = "Calm down, I know what I'm doing!";

    public String hey(String speech) {
        if (isEmpty(speech)) return EMPTY_RESPONSE;
        if (isQuestion(speech) && isYelling(speech)) return YELLING_QUESTION;
        if (isQuestion(speech)) return QUESTION_RESPONSE;
        if (isYelling(speech)) return YELLING_RESPONSE;
        return DEFAULT_RESPONSE;
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
