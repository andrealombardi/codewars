import java.awt.*;
import java.util.*;
import java.util.List;
import java.util.stream.Collectors;

public class BattleShips {

    public static Map<String, Double> damagedOrSunk(final int[][] board, final int[][] attacks) {

        final String SUNK = "sunk";
        final String NOT_TOUCHED = "notTouched";
        final String DAMAGED = "damaged";
        final String POINTS = "points";

        List<Set<Point>> shipPositions = Arrays.asList(new HashSet<>(), new HashSet<>(), new HashSet<>());

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (board[i][j] > 0) {
                    shipPositions.get(board[i][j] - 1).add(new Point(j + 1, board.length - i));
                }
            }
        }

        Set<Point> attacksPositions = Arrays.stream(attacks).map(attack -> new Point(attack[0], attack[1])).collect(Collectors.toSet());

        Map<String, Double> results = new HashMap<>();
        results.put(SUNK, .0);
        results.put(NOT_TOUCHED, .0);
        results.put(DAMAGED, .0);
        results.put(POINTS, .0);

        shipPositions.stream()
                .filter(s -> !s.isEmpty())
                .forEach(s -> {
                    if (attacksPositions.containsAll(s)) {
                        results.merge(SUNK, 1.0, Double::sum);
                    } else if (Collections.disjoint(s, attacksPositions)) {
                        results.merge(NOT_TOUCHED, 1.0, Double::sum);
                    } else {
                        results.merge(DAMAGED, 1.0, Double::sum);
                    }
                });

        Double damageScore = results.get(DAMAGED) > 0 ? results.get(DAMAGED) / 2 : results.get(DAMAGED);
        results.put(POINTS, results.get(SUNK) - results.get(NOT_TOUCHED) + damageScore);

        return results;
    }

}