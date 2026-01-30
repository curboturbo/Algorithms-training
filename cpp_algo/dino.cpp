#include <iostream>
#include <vector>
#include <algorithm>
#include <memory>

struct Island {
    int id;
    long long treasure;
    std::vector<Island*> neighbors;
    Island(int id, long long treasure) : id(id), treasure(treasure) {}
};

long long dfs(Island* node, int visitedMask, std::vector<std::vector<long long>>& memo) {
    int nodeIdx = node->id - 1;
    if (memo[visitedMask][nodeIdx] != -1) {
        return memo[visitedMask][nodeIdx];
    }
    visitedMask |= (1 << nodeIdx);
    long long maxSum = node->treasure;
    for (Island* neighbor : node->neighbors) {
        int neighborIdx = neighbor->id - 1;
        if (visitedMask & (1 << neighborIdx)){continue;}
        long long candidate = node->treasure + dfs(neighbor, visitedMask, memo);
        if (candidate >= maxSum) {maxSum=candidate;}
    }
    memo[visitedMask][nodeIdx] = maxSum;
    return maxSum;
}


int main() {
    int n, m;
    std::cin >> n >> m;
    std::vector<std::unique_ptr<Island>> islands(n);
    for (int i = 0; i < n; i++) {
        long long treasure;
        std::cin >> treasure;
        islands[i] = std::make_unique<Island>(i + 1, treasure);
    }
    for (int i = 0; i < m; i++) {
        int a, b;
        std::cin >> a >> b;
        a--; b--;
        islands[a]->neighbors.push_back(islands[b].get());
        islands[b]->neighbors.push_back(islands[a].get());
    }
    int fullMask = 1 << n;
    std::vector<std::vector<long long>> memo(fullMask,std::vector<long long>(n, -1));
    long long answer = dfs(islands[0].get(), 0, memo);
    std::cout << answer << std::endl;
    return 0;
}