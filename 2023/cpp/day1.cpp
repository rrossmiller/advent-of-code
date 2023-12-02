#include <cctype>
#include <cstdio>
#include <fstream>
#include <string>
#include <vector>

using std::ifstream;
using std::printf;
using std::string;

void pt1(ifstream &dataFile) {
    string line;
    std::vector<int> lineNums;
    int sum = 0;
    // int l = 0;
    // for each line
    while (std::getline(dataFile, line)) {
        std::vector<int> nums;
        // printf("%d: %s\n", l, line.c_str());

        for (int i = 0; i < line.length(); i++) {
            if (std::isdigit(line[i])) {
                nums.push_back(line[i] - '0');
            }
        }

        // l++;
        sum += nums[0] * 10 + nums[nums.size() - 1];
    }
    printf("pt1: %d\n", sum);
}

int main() {
    // open the file
    ifstream dataFile("../data/1.txt"); // dataFile.open("../data/1.txt");
    pt1(dataFile);
    dataFile.close();
}
