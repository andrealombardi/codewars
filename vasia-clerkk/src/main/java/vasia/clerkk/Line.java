package vasia.clerkk;

public class Line {
    public static String Tickets(int[] peopleInLine) {

        var twentyFive = 0; var fifty = 0;

        for (int i : peopleInLine) {
            switch (i) {
                case 25: twentyFive++; break;
                case 50: fifty++; twentyFive--; break;
                case 100: if (fifty > 0) { fifty--; twentyFive--; } else twentyFive -= 3; break;
            }
            if (twentyFive < 0) return "NO";
        }
        return "YES";
    }
}