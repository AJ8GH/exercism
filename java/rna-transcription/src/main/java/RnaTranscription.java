import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.stream.Collectors;

class RnaTranscription {
    private final HashMap<String, String> RNA = createRnaMap();

    String transcribe(String dnaStrand) {
        List<String> strands = Arrays.asList(dnaStrand.split(""));
        List<String> transcribed = strands.stream()
                .map(RNA::get)
                .collect(Collectors.toList());
        System.out.println(transcribed);

        return String.join("", transcribed);
    }

    private HashMap<String, String> createRnaMap() {
        HashMap<String, String> rnaMap = new HashMap<>();
        rnaMap.put("G", "C");
        rnaMap.put("C", "G");
        rnaMap.put("T", "A");
        rnaMap.put("A", "U");
        rnaMap.put("", "");
        return rnaMap;
    }
}