#include <iostream>
#include <string>
#include <vector>

void pt1(std::vector<std::string> dat) {
    int sum = 0;
    // for each line
    for (auto line : dat) {
        // std::cout << line << std::endl;
        // fwd pass
        int i, j;
        int a = 0;
        int b = 0;
        for (i = 0; i < line.length(); i++) {
            if (std::isdigit(line[i])) {
                auto digit = line[i];
                a = std::stoi(&digit);
                sum += 10 * a;
                break;
            }
        }

        // backwards pass
        for (j = line.length(); j >= i; j--) {
            if (std::isdigit(line[j])) {
                std::cout << line[j] << " " << b << std::endl;
                auto digit = line[j];
                b = std::stoi(&digit);
                sum += b;
                break;
            }
        }

        std::cout << line + "} " << line[i] << "-" << line[j] << std::endl;
        std::cout << 10 * a << "|" << b << ": " << 10 * a + b << std::endl;
    }
    printf("pt1: %d\n", sum);
}
