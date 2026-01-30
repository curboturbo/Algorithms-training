#include <iostream>
#include <vector>
#include <algorithm>


int main(){
    int n;
    std::cin >> n;
    
    int min_vas = 10*2000;
    int min_mash = 10*2000;
    int max_vas = -1;
    int max_mash = -1;
    int total = 0;
    int x;

    for (int i = 0; i < n; ++i) {
        std::cin >> x;
        if (i % 2 == 0){
            total += x;
            max_vas = std::max(max_vas, x);
            min_vas = std::min(min_vas, x);
        } else {
            total -= x;
            max_mash = std::max(max_mash, x);
            min_mash = std::min(min_mash, x);
        }
    }

    int a = 2 * (max_mash - min_vas);
    int b = 2 * (min_mash - max_vas);
    int ans = std::max(a,b);
    ans = std::max(ans, 0);

    std::cout << total + ans << std::endl;
    return 0;
}
