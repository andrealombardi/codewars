package tosmallest;

import java.util.Arrays;
import java.util.Collections;
import java.util.stream.Collectors;

public class ToSmallest {

    public static long[] smallest(long n) {

        char[] chars = String.valueOf(n).toCharArray();
        char[] sorted = String.valueOf(n).toCharArray();
        Arrays.sort(sorted);

        char smallest = Character.MAX_VALUE;
        int index_of_smallest = Integer.MAX_VALUE;
        int index_of_replace = Integer.MAX_VALUE;

        boolean not_found = true;
        for (int i = 0; i < sorted.length && not_found; i++) {
            for (int j = 0; j < chars.length && not_found; j++) {
                if(chars[j] == sorted[i] && j>0) {
                    for (int k = 0; k < j && not_found; k++) {
                        if (chars[k] > chars[j]) {
                            index_of_smallest = j;
                            index_of_replace = k;
                            smallest = chars[j];
                            not_found = false;
                        }
                    }
                }
            }
        }


        if (index_of_smallest == 1) {
            while(chars[0]  > chars[index_of_smallest + 1]){
                index_of_smallest ++;
            }
            chars[index_of_smallest] = chars[0];
            chars[0] = smallest;
            return new long[] {Long.parseLong(new String(chars)), 0, index_of_smallest};
        } else {
            for (int i = index_of_smallest; i > index_of_replace; i--) {
                chars[i] = chars[i-1];
            }
            chars[index_of_replace] = smallest;
            return new long[] {Long.parseLong(new String(chars)), index_of_smallest, index_of_replace};
        }
    }
}