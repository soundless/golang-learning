/*
 * [352] Data Stream as Disjoint Intervals
 *
 * https://leetcode.com/problems/data-stream-as-disjoint-intervals
 *
 * Hard (39.69%)
 * Total Accepted:    
 * Total Submissions: 
 * Testcase Example:  '["SummaryRanges","addNum","getIntervals","addNum","getIntervals","addNum","getIntervals","addNum","getIntervals","addNum","getIntervals"]\n[[],[1],[],[3],[],[7],[],[2],[],[6],[]]'
 *
 * Given a data stream input of non-negative integers a1, a2, ..., an, ...,
 * summarize the numbers seen so far as a list of disjoint intervals.
 * 
 * For example, suppose the integers from the data stream are 1, 3, 7, 2, 6,
 * ..., then the summary will be:
 * 
 * [1, 1]
 * [1, 1], [3, 3]
 * [1, 1], [3, 3], [7, 7]
 * [1, 3], [7, 7]
 * [1, 3], [6, 7]
 * 
 * 
 * Follow up:
 * What if there are lots of merges and the number of disjoint intervals are
 * small compared to the data stream's size?
 * 
 * 
 * Credits:Special thanks to @yunhong for adding this problem and creating most
 * of the test cases.
 */
/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type SummaryRanges struct {

}


/** Initialize your data structure here. */
func Constructor() SummaryRanges {

}


func (this *SummaryRanges) Addnum(val int)  {

}


func (this *SummaryRanges) Getintervals() []Interval {

}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
 */
