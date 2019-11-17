package persistence.bugger;

class Persist {
    public static int persistence(long n) {
        return doPersistence(0, n);
    }

    private static int doPersistence(int count, long bugger) {
        String str = String.valueOf(bugger);
        if (str.length() == 1) {
            return count;
        }

        return doPersistence(++count,  str.codePoints()
                .map(Character::getNumericValue)
                .reduce(1, (t, s) -> t * s));
    }
}