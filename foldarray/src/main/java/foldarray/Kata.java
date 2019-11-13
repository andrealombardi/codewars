import java.util.Arrays;

public class Kata {

    public static int[] splitAndAdd(int[] numbers, int n) {

        if (n == 0) return numbers;
        int[] a = Arrays.copyOfRange(numbers, 0, numbers.length / 2);
        int[] b = Arrays.copyOfRange(numbers, numbers.length / 2, numbers.length);
        int[] c = new int[Math.max(a.length, b.length)];

        for (int i = 1; i <= a.length; i++) {
            c[c.length - i] = a[a.length - i] + b[b.length - i];
        }

        if(a.length < b.length) {
            c[0] = b[0];
        }

        return splitAndAdd(c, n-1);
    }

}