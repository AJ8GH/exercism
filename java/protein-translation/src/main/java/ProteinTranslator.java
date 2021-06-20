import java.util.*;
import java.util.stream.Collectors;

class ProteinTranslator {
    private final HashMap<String, Set<String>> CODON_MAP = createCodonsMap();

    List<String> translate(String rnaSequence) {
        List<String> codonList = new ArrayList<>();
        for (int i = 0; i < rnaSequence.length(); i += 3) {
            codonList.add(rnaSequence.substring(i, i + 3));
        }

        List<String> translation = codonList.stream().map(codon -> {
            StringBuilder correctAminoAcid = new StringBuilder();
            for (String aminoAcid : CODON_MAP.keySet()) {
                if (CODON_MAP.get(aminoAcid).contains(codon)) {
                    correctAminoAcid.append(aminoAcid);
                    break;
                }
            }
            return correctAminoAcid.toString();
        }).collect(Collectors.toList());

        if (translation.contains("STOP")) {
            return translation.subList(0, translation.indexOf("STOP"));
        }
        return translation;
    }

    private HashMap<String, Set<String>> createCodonsMap() {
        HashMap<String, Set<String>> codonMap = new HashMap<>();
        codonMap.put("Methionine", new HashSet<>(Arrays.asList("AUG")));
        codonMap.put("Phenylalanine", new HashSet<>(Arrays.asList("UUU", "UUC")));
        codonMap.put("Leucine", new HashSet<>(Arrays.asList("UUA", "UUG")));
        codonMap.put("Serine", new HashSet<>(Arrays.asList("UCU", "UCC", "UCA", "UCG")));
        codonMap.put("Tyrosine", new HashSet<>(Arrays.asList("UAU", "UAC")));
        codonMap.put("Cysteine", new HashSet<>(Arrays.asList("UGU", "UGC")));
        codonMap.put("Tryptophan", new HashSet<>(Arrays.asList("UGG")));
        codonMap.put("STOP", new HashSet<>(Arrays.asList("UAA", "UAG", "UGA")));
        return codonMap;
    }
}