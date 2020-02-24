#include <iostream>

using namespace std;

int main(void) {
    int a, b, c;
    cin >> a >> b >> c;
    if ((a == b && b == c) || (a != b && b != c && c != a)) {
      cout << "No" << endl;
    } else {
      cout << "Yes" << endl;
    }

}
