#include <cstdio>
#include <string>
#include <vector>

#include "util.h"

using std::string;
using std::vector;
void part1(vector<string>& d, vector<coordinate>& g);
// void pt2(ifstream& d);

vector<string> getData(bool test, bool pt2, vector<coordinate>& galaxies) {
    vector<string> data;
    if (test) {
        data = getTestData();
    } else {
        data = getFileData();
    }
    int rows = data.size();
    int cols = data[0].length();

    // find rows and cols with only space
    vector<int> emptyRows;
    vector<int> emptyCols;
    for (int i = 0; i < rows; i++) {
        // if the whole row is '.'
        bool isEmpty = true;
        for (int j = 0; j < cols; j++) {
            char c = data[i][j];
            if (c != '.') {
                isEmpty = false;
            }
        }
        if (isEmpty) {
            emptyRows.push_back(i);
        }
    }
    for (int i = 0; i < cols; i++) {
        // if the whole col is '.'
        bool isEmpty = true;
        for (int j = 0; j < rows; j++) {
            char c = data[j][i];
            if (c != '.') {
                isEmpty = false;
            }
        }
        if (isEmpty) {
            emptyCols.push_back(i);
        }
    }

    vector<string> expandedSpace;
    int nCols = cols + emptyCols.size();
    for (int i = 0; i < rows; i++) {
        string row = "";
        for (int j = 0; j < cols; j++) {
            row.push_back(data[i][j]);
            // add the whole row and any extra empty column space
            if (contains(emptyCols, j)) {
                if (pt2) {
                    for (int i = 0; i < 1000000; i++) {
                        row.push_back('.');
                    }
                } else {
                    row.push_back('.');
                }
            }
        }
        // add the row
        expandedSpace.push_back(row);
        // if the whole row is empty, add another row below it
        if (contains(emptyRows, i)) {
            if (pt2) {
                for (int i = 0; i < 1000000; i++) {
                    expandedSpace.push_back(repeat('.', nCols));
                }
            } else {
                expandedSpace.push_back(repeat('.', nCols));
            }
        }
    }

    int n = 1;
    // find the galexies in expandd space
    for (int i = 0; i < expandedSpace.size(); i++) {
        for (int j = 0; j < expandedSpace[0].length(); j++) {
            if (expandedSpace[i][j] != '.') {
                expandedSpace[i][j] = std::to_string(n)[0];
                galaxies.push_back(coordinate{n, i, j});
                n++;
            }
        }
    }
    return expandedSpace;
}

int main() {
    bool test = false;
    bool pt2 = true;
    printf("*******\nDay 11\n");
    vector<coordinate> galaxies;  // don't pass mutable variables in and have
                                  // something be returned...
    auto data = getData(test, pt2, galaxies);
    printf("got data...\n");

    part1(data, galaxies);
}
