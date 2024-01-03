# Queue implementation
# class MovingAverage:
#
#     def __init__(self, size: int):
#         self.size = size
#         self.queue = []
#
#     def next(self, val: int) -> float:
#         self.queue.append(val)
#         window_sum = sum(self.queue[-self.size:])
#
#         return window_sum / min(len(self.queue), self.size)


# Deque implementation
from collections import deque
class MovingAverage:

    def __init__(self, size: int):
        self.size = size
        self.queue = deque()
        self.window_sum = 0

    def next(self, val: int) -> float:
        self.queue.append(val)
        left = self.queue.popleft() if len(self.queue) > self.size else 0
        self.window_sum = self.window_sum + val - left

        return self.window_sum / min(len(self.queue), self.size)


