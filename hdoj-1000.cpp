#include <iostream>
using namespace std;

int solve(int a, int b);

int main(int argc ,char **argv)
{
    int a, b;
    while (cin >> a >> b) {
        cout << solve(a, b) << endl;
    }
    return 0;
}

int solve(int a, int b)
{
    return a + b;
}
