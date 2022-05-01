class Solution:
    def divide(self, a, b):
        result = 0
        limit = 2 ** 31 - 1
        c = a if a > 0 else -1 * a
        d = b if b > 0 else -1 * b
        while c >= d:
            temp = 0
            while d << (temp + 1) <= c:
                temp += 1
            c -= d << temp
            result += 1 << temp
        if a < 0 and b < 0:
            result = result if result < limit else limit
            return result
        elif a < 0 or b < 0:
            return result * -1
        return result