package tosmallest;

import org.junit.Test;

import java.util.Arrays;

import static org.junit.Assert.*;

public class ToSmallestTest {
    private static void testing(long n, String res) {
        assertEquals(res,
                Arrays.toString(ToSmallest.smallest(n)));
    }

    @Test
    public void test() {
        System.out.println("Basic Tests smallest");
        testing(261235, "[126235, 2, 0]");

    }
    @Test
    public void test1() {
        System.out.println("Basic Tests smallest");
        testing(209917, "[29917, 0, 1]");
    }
    @Test
    public void test2() {
        System.out.println("Basic Tests smallest");
        testing(285365, "[238565, 3, 1]");
    }
    @Test
    public void test3() {
        System.out.println("Basic Tests smallest");
        testing(269045, "[26945, 3, 0]");
    }
    @Test
    public void test4() {
        System.out.println("Basic Tests smallest");
        testing(296837, "[239687, 4, 1]");
    }
    @Test
    public void test5() {
        testing(11131, "[11113, 4, 0]");
    }

    @Test
    public void test6() {
        testing(964821571181538688L, "[196482157118538688, 11, 0]");
    }
    @Test
    public void test7() {
        testing(903638528994158336L, "[36385289941583369, 0, 17]");
    }
    @Test
    public void test8() {
        testing(51291664257631672L, "[12591664257631672, 0, 2]");
    }
    @Test
    public void test9() {
        testing(588895571116954368L, "[158889557116954368, 8, 0]");
    }
    @Test
    public void test10() {
        testing(104996969700471488L, "[10499696970471488, 10, 0]");
    }
    @Test
    public void test11() {
        testing(803765573648623488L, "[37655736486234888, 0, 15]");
    }


}