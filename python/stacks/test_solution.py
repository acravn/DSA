import unittest
from solution import Solution


class TestSolution(unittest.TestCase):

    def setUp(self):
        self.s = Solution()

    def test_simplifyPath(self):
        cases = [
            {
                "input": "/home/",
                "output": "/home"
            },
            {
                "input": "/../",
                "output": "/"
            },
            {
                "input": "/home//foo/",
                "output": "/home/foo"
            },
            {
                "input": "/a/./b/../../c/",
                "output": "/c"
            }
        ]

        for case in cases:
            res = self.s.simplifyPath(case["input"])
            self.assertEqual(res, case['output'])

    def test_makeGood(self):
        cases = [
            {
                "input": "leEeetcode",
                "output": "leetcode"
            },
            {
                "input": "abBAcC",
                "output": ""
            },
            {
                "input": "s",
                "output": "s"
            }
        ]

        for case in cases:
            res = self.s.makeGood(case["input"])
            self.assertEqual(res, case['output'])
