class Solution:
    def simplifyPath(self, path: str) -> str:
        stack = []
        paths = path.split("/")

        for p in paths:
            if p == "..":
                if stack:
                    stack.pop()
            elif (p != "/") and (p != "") and (p != "."):
                stack.append(p)

        return "/" + '/'.join(stack)

    def makeGood(self, s: str) -> str:
        stack = []
        for curr_char in s:
            if stack and abs(ord(curr_char) - ord(stack[-1])) == 32:
                stack.pop()
            else:
                stack.append(curr_char)


        return "".join(stack)
