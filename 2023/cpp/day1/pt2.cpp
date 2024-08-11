#include <cstdio>
#include <fstream>
#include <iostream>
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
    string word = "";

    while (std::getline(dataFile, line)) {
        std::cout << line << std::endl;
    }
    return;
    // for each line
    while (std::getline(dataFile, line)) {
        printf("line: %s\n", line.c_str());
        std::vector<int> nums;
        for (int i = 0; i < line.length(); i++) {
            word += line[i];
            if (std::isdigit(line[i])) {
                // std::cout << "c: " << line[i] << std::endl;
                nums.push_back(line[i] - '0');
                word.clear();
                // word = "";
            } else {
                for (const std::pair<const std::string, int>& n : hashmap) {
                    if (word.contains(n.first)) {
                        // printf("map %s, %s\n", word.c_str(),
                        // n.first.c_str());
                        nums.push_back(n.second);
                        // word = "";
                        word.clear();
                        break;
                    }
                }
            }
        }

        l++;
        for (int i = 0; i < nums.size(); i++) {
            printf("%d ", nums[i]);
        }
        sum += nums[0] * 10 + nums[nums.size() - 1];
        printf("%d\n", nums[0] * 10 + nums[nums.size() - 1]);
        std::cout << std::endl;
    }
    printf("pt2: %d - %s\n", sum, sum == 55093 ? "ok" : "not right");
}
