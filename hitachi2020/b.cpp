#include <iostream>
#include <vector>

using namespace std;

int main(void){
    int an, bn, m;

    cin >> an >> bn >> m;
    int a_min, b_min, c_min;
    a_min = b_min = c_min = 1000000;

    vector<int> a(an), b(an);
    for (int i = 0; i < an; i++) {
        cin >> a[i];
        if (a_min > a[i]) {
            a_min = a[i];
        }
    }
    for (int i = 0; i < bn; i++) {
        cin >> b[i];
        if (b_min > b[i]) {
            b_min = b[i];
        }
    }
    for (int i = 0; i < m; i++) {
        int x, y, c;
        cin >> x >> y >> c;
        int p = a[x-1] + b[y-1] - c;
        if (c_min > p) {
            c_min = p;
        }
    }

    if (c_min < a_min + b_min) {
        cout << c_min << endl;
    } else {
        cout << a_min + b_min << endl;
    }
    
    return 0;
}