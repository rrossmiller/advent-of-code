#include "util.h"

bool contains(const std::vector<int> &v, int i) {
    for (int d : v) {
        if (d == i) {
            return true;
        }
    }
    return false;
}
std::string repeat(char s, int n) {
    std::string rtn = "";
    for (int i = 0; i < n; i++) {
        rtn.push_back(s);
    }
    return rtn;
}

std::vector<std::string> getTestData() {  // test data
    auto data = std::istringstream(
        "...#......\n"
        ".......#..\n"
        "#.........\n"
        "..........\n"
        "......#...\n"
        ".#........\n"
        ".........#\n"
        "..........\n"
        ".......#..\n"
        "#...#.....");
    std::string line;
    std::vector<std::string> lines;
    while (std::getline(data, line)) {
        lines.push_back(line);
    }

    return lines;
}
std::vector<std::string> getFileData() {
    std::ifstream dataFile("../data/11.txt");
    std::string line;
    std::vector<std::string> lines;
    while (std::getline(dataFile, line)) {
        lines.push_back(line);
    }
    dataFile.close();


    return lines;
}
