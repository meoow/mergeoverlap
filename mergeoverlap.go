package mergeoverlap

import (
	"sort"
)

type (
	point struct {
		value int
		orien bool
	}
	sortpoints []*point
)

func MergeOverlap(intervals [][2]int) [][2]int {
	axis := make([]*point, 0, len(intervals)*2)
	for _, i := range intervals {
		axis = append(axis, &point{i[0], false})
		axis = append(axis, &point{i[1], true})
	}
	sort.Sort(sortpoints(axis))
	new_axis := make([]*point, 0, len(intervals))
	lower := 0
	for _, p := range axis {
		if !p.orien {
			lower++
			if lower == 1 {
				new_axis = append(new_axis, p)
			}
		} else {
			lower--
			if lower == 0 {
				new_axis = append(new_axis, p)
			}
		}
	}
	sort.Sort(sortpoints(new_axis))
	pp := make([][2]int, 0, len(new_axis)/2)
	sp := [2]int{0, 0}
	for i, p := range new_axis {
		if i%2 == 0 {
			sp[0] = p.value
		} else {
			sp[1] = p.value
			//fmt.Println(sp)
			pp = append(pp, sp)
			sp = [2]int{0, 0}
		}
	}
	return pp
}

func (sp sortpoints) Len() int {
	return len(sp)
}

func (sp sortpoints) Swap(i, j int) {
	sp[i], sp[j] = sp[j], sp[i]
}

func (sp sortpoints) Less(i, j int) bool {
	return sp[i].value < sp[j].value
}
