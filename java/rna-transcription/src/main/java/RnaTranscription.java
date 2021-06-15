import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.stream.Collectors;

class RnaTranscription {
    String transcribe(String dnaStrand) {
        HashMap<String, String> rna = new HashMap<>();
        rna.put("G", "C");
        rna.put("C", "G");
        rna.put("T", "A");
        rna.put("A", "U");
        rna.put("", "");

        List<String> strands = Arrays.asList(dnaStrand.split(""));
        List<String> transcribed = strands.stream()
                .map(rna::get)
                .collect(Collectors.toList());
        System.out.println(transcribed);

        return String.join("", transcribed);
    }
}