package tosmallest;

import java.util.Comparator;
import java.util.stream.IntStream;

public class ToSmallest {

    public static long[] smallest(long n) {

        var s = String.valueOf(n);
        var comparator = Comparator.<Candidate>comparingLong(c -> c.n).thenComparingLong(c -> c.i).thenComparingLong(c -> c.j);

        return IntStream.range(0, s.length()).boxed()
                .flatMap(i -> (IntStream.range(0, s.length()).boxed()
                        .map(j -> new Candidate(s, i, j))))
                .min(comparator).map(Candidate::toArray).orElseThrow();
    }

    static class Candidate {
        long n, i, j;

        Candidate(String s, int i, int j) {
            this.n = build(s, i, j); this.i = i; this.j = j;
        }

        long[] toArray() {
            return new long[]{n, i, j};
        }

        private static long build(String s, int i, int j) {
            StringBuilder sb = new StringBuilder(s);
            char c = sb.charAt(i);
            sb.deleteCharAt(i);
            sb.insert(j, c);
            return Long.parseLong(sb.toString());
        }
    }

}