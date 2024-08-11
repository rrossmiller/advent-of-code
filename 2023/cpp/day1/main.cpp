#include <cstdio>
#include <fstream>
#include <vector>

using std::ifstream;

void pt1(std::vector<std::string> d);
void pt2(std::vector<std::string> d);

std::vector<std::string> readF(std::string pth) {
    std::vector<std::string> rtn;
    ifstream f(pth);

    std::string line;
    while (std::getline(f, line)) {
        rtn.push_back(line);
    }
    return rtn;
}
int main() {
    printf("Day 1\n");
    // read the file
    auto dat = readF("../data/1.txt");
    // auto dat = readF("test.txt");

    pt1(dat);
    // pt2(dat);
}
