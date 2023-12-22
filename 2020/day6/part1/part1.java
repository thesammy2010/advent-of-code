package advent_of_code.day6.part1;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Arrays;
import advent_of_code.helpers.Helpers;

public class part1 {

    public static List<String> compileGroups(String data) {
        List<String> groups = new ArrayList<>();
        String[] newData = data.split("\n\n");
        for (int i = 0; i < newData.length; i++) {
            groups.add(newData[i].replaceAll("\n", ""));
        }
        return groups;
    }

    public static int main(String filepath) throws Exception {
        String file = Helpers.readFile(filepath);
        List<String> groups = compileGroups(file);
        int count = 0;
        for (int i = 0; i < groups.size(); i++) {
            HashSet<String> set = new HashSet<String>(Arrays.asList(groups.get(i).split("")));
            count += set.size();
        }
        return count;
    }
    
}
