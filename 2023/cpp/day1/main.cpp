#include <cstdio>
#include <fstream>

using std::ifstream;

void pt1(ifstream& d);
void pt2(ifstream& d);
int main() {
    printf("Day 1\n");
    // open the file
    ifstream dataFile("../data/1.txt");
    pt1(dataFile);
    dataFile.close();

    ifstream f("../data/1.txt");
    // ifstream f("test.txt");
    pt2(f);
    f.close();
}
