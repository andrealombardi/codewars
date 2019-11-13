package countip;

import java.util.Arrays;
import java.util.function.Function;

public class CountIPAddresses {

    public static long ipsBetween(String start, String end) {
        long a = toIntArray.andThen(toBaseTen).apply(start);
        long z = toIntArray.andThen(toBaseTen).apply(end);
        return z-a;
    }

    private static Function<int[], Long> toBaseTen = ip ->
            (long) (ip[0] * Math.pow(256, 3) + ip[1] * Math.pow(256, 2) + ip[2] * 256 + ip[3]);

    private static Function<String, int[]> toIntArray = ip ->
        Arrays.stream(ip.split("\\.")).mapToInt(Integer::parseInt).toArray();

}
