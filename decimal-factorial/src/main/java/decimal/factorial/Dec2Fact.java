package decimal.factorial;

import java.util.Arrays;
import java.util.List;

public class Dec2Fact {

    static final List<Character> NUMS = List.of(
            '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B',
            'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
            'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z');

    public static String dec2FactString(long nb) {
        var quotient = nb;
        var remainders = new StringBuilder();

        for (int i = 1; i < NUMS.size() && quotient > 0; i++) {
            remainders.append(NUMS.get((int) (quotient%i)));
            quotient = quotient/i;
        }

        return remainders.reverse().toString();
    }

    public static long factString2Dec(String str) {
        String reversed = new StringBuilder(str).reverse().toString();

        long acc = 0;
        for (int i = 0; i < reversed.length(); i++) {
            Character c = reversed.charAt(i);
            acc += NUMS.indexOf(c) * factorial(i);
        }
        return acc;
    }

    private static long factorial(long i) {
        if (i < 1) return 1;
        return i * factorial(i - 1);
    }

}