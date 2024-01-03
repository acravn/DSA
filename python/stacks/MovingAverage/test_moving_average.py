import unittest
from moving_average import MovingAverage

class TestMovingAverage(unittest.TestCase):

    def test_next(self):
        testcases = [
            {
                "size": 3,
                "vals": [1,10,3,5],
                "expected": [1.0, 5.5, 4.666666666666667, 6.0]
            },
            {
                "size": 1,
                "vals": [4, 0],
                "expected": [4, 0]
            }
        ]

        for case in testcases:
            obj = MovingAverage(case['size'])
            for i in range(len(case['vals'])):
                avg = obj.next(case['vals'][i])
                self.assertEqual(avg,case['expected'][i])

