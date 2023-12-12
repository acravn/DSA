import unittest
from main import Solution

class Test_Solution(unittest.TestCase):

    def setUp(self):
        self.solution = Solution()

    def test_findMaxLength(self):
        cases = [
            {
                "input": [0,1],
                "output": 2
            },
            {
                "input": [0,1,0],
                "output": 2
            }
        ]

        for case in cases:
            res = self.solution.findMaxLength(case['input'])
            self.assertEqual(res, case['output'])