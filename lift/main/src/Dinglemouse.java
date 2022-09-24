import java.util.*;
import java.util.function.Predicate;

import static java.util.stream.Collectors.toList;

public class Dinglemouse {

    public static int[] theLift(final int[][] queues, final int capacity) {
        Building b = new Building(queues, capacity);
        return b.fulfilRequest();
    }

    static class Building {
        private final int maxCapacity;
        private List<Integer> visited = new ArrayList<>();
        private List<Integer> lift = new ArrayList<>();
        private Map<Integer, List<Integer>> floormap = new TreeMap<>();

        public Building(int[][] queues, int capacity) {
            for (int i = 0; i < queues.length; i++) {
                floormap.put(i, Arrays.stream(queues[i]).boxed().collect(toList()));
            }
            maxCapacity = capacity;
        }

        public int[] fulfilRequest(){
            visited.add(0);
            while (!finished()) {
                goUp();
                goDown();
            }
            if (getLastVisited() != 0) visited.add(0);
            return visited.stream().mapToInt(Integer::intValue).toArray();
        }

        private void goUp() {
            for (Map.Entry<Integer, List<Integer>> floor : floormap.entrySet()) {
                evaluateFloor(floor, destination -> destination > floor.getKey());
            }
        }
        private void goDown() {
            var down = new ArrayList<>(floormap.entrySet());
            Collections.reverse(down);
            for (Map.Entry<Integer, List<Integer>> floor : down) {
                evaluateFloor(floor, destination -> destination < floor.getKey());
            }
        }

        private void evaluateFloor(Map.Entry<Integer, List<Integer>> floor, Predicate<Integer> isGoingInTheSameDirection) {
            boolean stopped = false;
            var out = lift.stream().filter(destination -> destination == floor.getKey()).collect(toList());
            if (!out.isEmpty()) {
                stopped = true;
                lift.removeAll(out);
            }

            //Necessary because the lift should stop even if there is no capacity
            var callers= floor.getValue().stream().filter(isGoingInTheSameDirection).findAny();
            if(!callers.isEmpty()) {
                stopped = true;
                var in = floor.getValue().stream()
                        .filter(isGoingInTheSameDirection)
                        .limit(getCurrentCapacity())
                        .collect(toList());
                lift.addAll(in);
                in.stream().forEach(floor.getValue()::remove);
            }

            if (stopped && floor.getKey() != getLastVisited()) {
                visited.add(floor.getKey());
            }
        }

        private int getLastVisited() {
            return visited.get(visited.size() -1);
        }

        private boolean finished() {
            return floormap.entrySet().stream().filter(e -> ! e.getValue().isEmpty()).findAny().isEmpty();
        }

        private int getCurrentCapacity(){
            return maxCapacity - lift.size();
        }
    }
}