package recreation.one;

import java.util.Collections;
import java.util.List;
import java.util.function.Predicate;
import java.util.stream.Collectors;
import java.util.stream.LongStream;

public class SumSquaredDivisors {

    public static String listSquared(long m, long n) {

        List<List<Long>> results = LongStream.rangeClosed(m, n)
                .boxed()
                .map(SumSquaredDivisors::sumOfSquares)
                .filter(Predicate.not(List::isEmpty))
                .collect(Collectors.toList());

        return results.toString();
    }

    private static List<Long> sumOfSquares(Long l) {

        long candidate = LongStream.rangeClosed(1, l)
                .filter(s -> l % s == 0)
                .map(t -> t * t)
                .distinct()
                .reduce(0L, Long::sum);

        return Math.sqrt(candidate) % 1 == 0 ? List.of(l, candidate) : Collections.emptyList();
    }
}
