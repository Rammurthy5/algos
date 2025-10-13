is typically used to find max sum of subarray (contigous), submatrix. 

## Real world usecases
Stock Trading Profit Maximization: Used in real-time or near real-time analysis of daily stock price changes
(where the array elements are daily profits/losses) to quickly find the time window (contiguous subarray) that would have yielded
the maximum profit for a single buy-and-sell transaction.

Budget Optimization: Identifying the consecutive period (e.g., months) with the maximum net savings or highest disposable income in financial tracking applications.

Peak Detection: In telecommunication, radar systems, or general data streams (like from environmental sensors), 
it can identify the segment of a signal that shows the maximum intensity or most prominent change over a continuous duration,
which is crucial for noise reduction and signal enhancement.

Event Detection: Analyzing sensor network data (e.g., temperature, pollution) to quickly pick out periods of unusual activity, such as a continuous spike in pollution levels.

Trend and Anomaly Detection: Analyzing time-series data to quickly spot the period of maximum positive or negative trend (growth or decline) in datasets, 
which is often used in business intelligence or market research tools to highlight significant changes.

Sequence Feature Identification: Used to identify significant regions in DNA, RNA, or protein sequences 
(e.g., a high concentration of specific bases like GC-rich regions), which helps in understanding biological processes or diseases.

Score Optimization/Streak Tracking: Calculating the maximum continuous score streak or combo a player has achieved over a series of levels
or time limits for high-score tracking and analysis.

## time complexity 
O(n)

## space complexity
O(1)
