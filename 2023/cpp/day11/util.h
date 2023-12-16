#pragma once
#include <cstdint>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

struct coordinate {
    int id;
    int x;
    int y;
};

bool contains(const std::vector<int> &v, int i);
std::string repeat(char s, int n);
std::vector<std::string> getTestData();  // test data
std::vector<std::string> getFileData();
