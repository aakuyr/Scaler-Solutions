class Solution:
    # @param A : list of integers
    # @param B : list of integers
    # @return an integer
    def moveToLastOfList(self, A):
        first = A[0]
        A = A[1:]
        A.append(first)
        return A
    
    def solve(self, A, B):
        count = 0
        while A:
            if A[0] != B[0]:
                while A[0] != B[0]:
                    A = self.moveToLastOfList(A)
                    count += 1
            count += 1
            A = A[1:]
            B = B[1:]
        return count
