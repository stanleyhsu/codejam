/*
 * Japan acm 2.1.4 P30
 * */

#include <iostream>

using namespace std;

const int MAX_N = 20;
int g_input_counter = 5;
int g_input_array[MAX_N] = {1, 2, 4, 7, 8};
int g_expect_sum = 13;

bool dfs(int i, int sum) {
    if (i == g_input_counter) return sum == g_expect_sum;

    if (dfs(i + 1, sum)) return true;
    if (dfs(i + 1, sum + g_input_array[i])) return true;
    return false;
}

int main(int argc, char *argv[]) {
    if (dfs(0, 0)) cout << "Yes";
    else cout << "No";
    cout << endl;
    return 0;
}

