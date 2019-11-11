import java.awt.*;
import java.util.*;
import java.util.List;
import java.util.stream.Collectors;

public class BattleShips {

    public static Map<String, Double> damagedOrSunk(final int[][] board, final int[][] attacks) {

        List<Set<Point>> shipPositions = Arrays.asList(new HashSet<>(), new HashSet<>(), new HashSet<>());

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[i].length; j++) {
                if (board[i][j] > 0) {
                    shipPositions.get(board[i][j] - 1).add(new Point(board.length - i, j + 1));
                }
            }
        }

        Set<Point> attacksPositions = Arrays.stream(attacks).map(attack -> new Point(attack[0], attack[1])).collect(Collectors.toSet());

        System.out.println(shipPositions);
        System.out.println(attacksPositions);

        Map<String, Double> results = new HashMap<>();

        shipPositions.stream()
                .filter(s -> !s.isEmpty())
                .forEach(s -> {
                    if (attacksPositions.containsAll(s)) {
                        results.merge("sunk", 1.0, Double::sum);
                    } else if (Collections.disjoint(s, attacksPositions)) {
                        results.merge("notTouched", 1.0, Double::sum);
                    } else {
                        results.merge("damaged", .5, Double::sum);
                    }
                });

        results.put("points", results.getOrDefault("sunk", 0.0)
                + results.getOrDefault("damaged", 0.0)
                - results.getOrDefault("notTouched", 0.0));

        return results;
    }

}