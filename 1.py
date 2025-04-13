import math
s = input()
n = len(s)
dp = []
pos = []
for i in range(n):
    dp.append([0] * n)
    pos.append([0] * n)
for i in range(n):
    for j in range(n):
        if i == j:
            dp[i][j] = 1
for right in range(n):
    for left in range(right, -1, -1):
        if left == right:
            dp[left][right] = 1
        else:
            min = math.inf
            mink = -1
            b1 = s[left] == '(' and s[right] == ')'
            b2 = s[left] == '[' and s[right] == ']'
            b3 = s[left] == '{' and s[right] == '}'
            if b1 or b2 or b3:
                min = dp[left + 1][right - 1]
            for k in range(left, right):
                if min > dp[left][k] + dp[k + 1][right]:
                    min = dp[left][k] + dp[k + 1][right]
                    mink = k
            dp[left][right] = min
            pos[left][right] = mink
def rec(l, r):
    temp = r - l + 1
    if dp[l][r] == temp:
        return
    if dp[l][r] == 0:
        print(s[l:r + 1], end="")
        return
    if pos[l][r] == -1:
        print(s[l], end="")
        rec(l + 1, r - 1)
        print(s[r], end="")
        return
    rec(l, pos[l][r])
    rec(pos[l][r] + 1, r)
rec(0, n - 1)