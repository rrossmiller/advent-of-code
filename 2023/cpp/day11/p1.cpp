#include <cstdio>
#include <cstdlib>
#include <string>
#include <unordered_set>
#include <vector>

#include "util.h"
using std::string;
using std::vector;

void part1(vector<string> &data, vector<coordinate> &galaxies) {
    int sum = 0;
    std::unordered_set<string> done;
    for (int i = 0; i < galaxies.size(); i++) {
        auto gi = galaxies[i];
        // printf("%d...\n", gi.id);
        for (int j = 0; j < galaxies.size(); j++) {
            if (i == j) {
                continue;
            }
            auto gj = galaxies[j];
            string gix = std::to_string(gi.id);
            string gjx = std::to_string(gj.id);
            if (!(done.contains(gix + "|" + gjx) ||
                  done.contains(gjx + "|" + gix))) {
                done.insert(gix + "|" + gjx);
                int dist = abs(gi.x - gj.x) + abs(gi.y - gj.y);
                sum += dist;
            }
        }
    }

    printf("part1: %d\n", sum);
}
