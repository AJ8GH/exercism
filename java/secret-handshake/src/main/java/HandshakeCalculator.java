import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

class HandshakeCalculator {

    List<Signal> calculateHandshake(int number) {
        List<Signal> signals = Arrays.asList(Signal.values());
        List<Signal> output = new ArrayList<>();

        String binary = Integer.toBinaryString(number);

        int count = 0;
        for (int i = binary.length() -1; i >= 0; i--) {
            if (binary.charAt(i) == '1') {
                if (count == 4) {
                    Collections.reverse(output);
                } else {
                    output.add(signals.get(count));
                }
            }
            count++;
        }
        return output;
    }
}
