#include <cstdio>
#include <fstream>
#include <string>
#include <unordered_map>
#include <vector>

using std::ifstream;
using std::printf;
using std::string;
using std::unordered_map;

const unordered_map<string, int> hashmap = {
    {"one", 1}, {"two", 2},   {"three", 3}, {"four", 4}, {"five", 5},
    {"six", 6}, {"seven", 7}, {"eight", 8}, {"nine", 9},
};

void pt2(ifstream& dataFile) {
    int sum = 0;
    int l = 0;
    string line;
    string word;

    // for each line
    while (std::getline(dataFile, line)) {
        std::vector<int> nums;
        for (int i = 0; i < line.length(); i++) {
            word += line[i];
            if (std::isdigit(line[i])) {
                nums.push_back(line[i] - '0');
                word.clear();
            } else {
                for (const std::pair<const std::string, int>& n : hashmap) {
                    if (word.contains(n.first)) {
                        nums.push_back(n.second);
                        word.clear();
                        break;
                    }
                }
            }
        }
        l++;
        sum += nums[0] * 10 + nums[nums.size() - 1];
    }
    printf("pt2: %d\n", sum);
}
