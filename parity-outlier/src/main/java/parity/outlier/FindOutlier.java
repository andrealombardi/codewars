package parity.outlier;

import java.util.function.IntPredicate;
import java.util.stream.Collectors;

import static java.util.Arrays.stream;
import static java.util.stream.Collectors.groupingBy;

public class FindOutlier {
    static int find(int[] integers) {
        var firstThree = stream(integers).boxed().limit(3).collect(groupingBy(i -> i % 2, Collectors.counting()));
        IntPredicate oddOrEven = i -> firstThree.getOrDefault(0, 0L) > 1 ? Math.abs(i) % 2 == 1 : i % 2 == 0;
        return stream(integers).parallel().filter(oddOrEven).findAny().orElseThrow();
    }
}