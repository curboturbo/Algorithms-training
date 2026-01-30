#include <iostream>
#include <stack>
#include <vector>



//Input: nums = [1,2,3,4,3]
//Output: [2,3,4,-1,4]


std::vector<int> nextGreaterElements(std::vector<int>& nums) {
    int n = nums.size();
    std::vector<int> answer(n, -1);
    std::stack<int> st;
    for (int i =n-1; i>=0;i--) {
        while (!st.empty() && st.top() <= nums[i]) {
            st.pop();
        }
        if (!st.empty()) {
            answer[i] = st.top();
        }
        st.push(nums[i]);
    }
    for (int i =n-1; i>=0;i--) {
        while (!st.empty() && st.top() <= nums[i]) {
            st.pop();
        }
        if (!st.empty()) {
            answer[i] = st.top();
        }
        st.push(nums[i]);
    }

    return answer;

}


int main() {
    std::vector<int> nums = {5,4,3,2,1};
    auto res = nextGreaterElements(nums);

    for (int x : res) {
        std::cout << x << " ";
    }
}
