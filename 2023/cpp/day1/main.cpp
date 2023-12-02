#include <fstream>

using std::ifstream;

void pt1(ifstream& d);
void pt2(ifstream& d);
int main() {
    // open the file
    ifstream dataFile("../data/1.txt");  // dataFile.open("../data/1.txt");
    pt1(dataFile);
    pt2(dataFile);
    dataFile.close();
}
