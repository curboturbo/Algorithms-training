#include <iostream>
#include <vector>
#include <stack>

std::vector<int> HopeItWorks(int n, const std::vector<int>& h) {
    std::vector<int> ans(n, -1);
    std::stack<std::pair<int, int>> stack;
    for (int i = n - 1; i >= 0; --i) {
        if ((i + 1) % 2 == 1) {
            while (!stack.empty() && stack.top().first <= h[i]) {
                stack.pop();
            }
            if (!stack.empty()) {
                ans[i] = stack.top().second - i;
            }
            stack.push({h[i], i});
        }
    }
    
    std::stack<std::pair<int, int>> stackend;
    for (int i = n - 1; i >= 0; --i) {
        if ((i + 1) % 2 == 0) {
            while (!stackend.empty() && stackend.top().first <= h[i]) {
                stackend.pop();
            }
            if (!stackend.empty()) {
                ans[i] = stackend.top().second - i;
            }
            stackend.push({h[i], i});
        }
    }
    return ans;
}

int main() {
    int n;
    std::cin >> n;
    std::vector<int> h(n);
    for (int i = 0; i < n; ++i) {
        std::cin >> h[i];
    }
    std::vector<int> result = HopeItWorks(n, h);
    for (int i = 0; i < n; ++i) {
        std::cout << result[i];
        if (i < n - 1) std::cout << " ";
    }
    std::cout << std::endl;
    return 0;
}