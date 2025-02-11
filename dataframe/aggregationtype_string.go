// Code generated by "stringer -type=AggregationType -linecomment"; DO NOT EDIT.

package dataframe

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Aggregation_MAX-1]
	_ = x[Aggregation_MIN-2]
	_ = x[Aggregation_MEAN-3]
	_ = x[Aggregation_MEDIAN-4]
	_ = x[Aggregation_STD-5]
	_ = x[Aggregation_SUM-6]
	_ = x[Aggregation_COUNT-7]
	_ = x[Aggregation_PROD-8]
}

const _AggregationType_name = "MAXMINMEANMEDIANSTDSUMCOUNTPROD"

var _AggregationType_index = [...]uint8{0, 3, 6, 10, 16, 19, 22, 27, 31}

func (i AggregationType) String() string {
	i -= 1
	if i < 0 || i >= AggregationType(len(_AggregationType_index)-1) {
		return "AggregationType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AggregationType_name[_AggregationType_index[i]:_AggregationType_index[i+1]]
}
