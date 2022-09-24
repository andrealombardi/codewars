import java.util.*;
import java.util.function.Predicate;

import static java.util.stream.Collectors.toList;

public class DinglemouseFirst {

    static List<Integer> visited;
    static List<Integer> lift;
    static Map<Integer, List<Integer>> floormap;
    static int MAX_CAPACITY = 0;
    static int lastVisited = 0;

    public static int[] theLift(final int[][] queues, final int capacity) {

        MAX_CAPACITY = capacity;
        lastVisited = 0;
        visited = new ArrayList<>();
        lift = new ArrayList<>();
        floormap = new TreeMap<>();

        for (int i = 0; i < queues.length; i++) {
            floormap.put(i, Arrays.stream(queues[i]).boxed().collect(toList()));
        }
        System.out.println(capacity);
        System.out.println(floormap);

        visited.add(0);

        while (!finished()) {
            for (Map.Entry<Integer, List<Integer>> floor : floormap.entrySet()) {
                evaluateFloor(floor, destination -> destination > floor.getKey());
            }

            var down = new ArrayList<>(floormap.entrySet());
            Collections.reverse(down);
            for (Map.Entry<Integer, List<Integer>> floor : down) {
                evaluateFloor(floor, destination -> destination < floor.getKey());
            }
        }

        if (lastVisited != 0) visited.add(0);

        System.out.println(visited);

        // Your code here
        return visited.stream().mapToInt(Integer::intValue).toArray();
    }

    private static void evaluateFloor(Map.Entry<Integer, List<Integer>> floor, Predicate<Integer> p) {
        boolean stopped = false;
        var out = lift.stream().filter(destination -> destination == floor.getKey()).collect(toList());
        if (!out.isEmpty()) {
            stopped = true;
            lift.removeAll(out);
        }

        var callers= floor.getValue().stream().filter(p).findAny();
        if(!callers.isEmpty()) {
            stopped = true;
            var in = floor.getValue().stream().filter(p).limit(getCurrentCapacity()).collect(toList());
            lift.addAll(in);
            in.stream().forEach(floor.getValue()::remove);
        }

        if (stopped && floor.getKey() != lastVisited) {
            visited.add(floor.getKey());
            lastVisited = floor.getKey();
        }
    }

    private static boolean finished() {
        return floormap.entrySet().stream().filter(e -> ! e.getValue().isEmpty()).findAny().isEmpty();
    }

    private static int getCurrentCapacity(){
        return MAX_CAPACITY - lift.size();
    }

}