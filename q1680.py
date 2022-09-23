class Solution:
    def concatenatedBinary(self, n: int) -> int:
        binaries = []
        for i in range(n):
            binaries.append(format(i+1, "b"))
        conc = "".join(binaries)
        dec = int(conc, 2)
        print(f"Called with {n}, returned {dec}")
        return dec % ((10 ** 9) + 7)

if __name__ == "__main__":
    s = Solution()
    assert 1 == s.concatenatedBinary(1)
    assert 27 == s.concatenatedBinary(3)
    assert 505379714 == s.concatenatedBinary(12)
