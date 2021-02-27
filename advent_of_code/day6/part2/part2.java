package advent_of_code.day6.part2;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Arrays;
import advent_of_code.helpers.Helpers;
import java.util.Collections;

public class part2 {

    public static List<List<String>> compileGroups(String data) {
        List<List<String>> groups = new ArrayList<>();
        String[] newData = data.split("\n\n");
        for (int i = 0; i < newData.length; i++) {
            List<String> group = new ArrayList<String>();
            String[] temp = newData[i].split("\n");
            for (int j = 0; j < temp.length; j++) {
                group.add(temp[j]);
            }
            groups.add(group);
        }
        return groups;
    }

    public static HashMap<String, Integer> groupData(List<String> group) {
        HashMap<String, Integer> map = new HashMap<String, Integer>();
        map.put("count", group.size());
        for (int p = 0; p < group.size(); p++) {
            String person = group.get(p);
            String[] letters = person.split("");
            for (int c = 0; c < letters.length; c++) {
                map.put(letters[c], map.getOrDefault(letters[c], 0) + 1);
            }
        }
        return map;
    }

    public static int getCountFromGroup(HashMap<String, Integer> map) {
        List<Integer> values = new ArrayList<Integer>(map.values());
        return Collections.frequency(values, map.get("count")) - 1;
    }

    public static int main(String filepath) throws Exception {
        String file = Helpers.readFile(filepath);
        List<List<String>> groups = compileGroups(file);
        int count = 0;
        for (int group = 0; group < groups.size(); group++) {
            HashMap<String, Integer> map = groupData(groups.get(group));
            count += getCountFromGroup(map);
        }
        return count;
    }
    
}
