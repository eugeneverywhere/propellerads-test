package handler

import "strings"

type Handler interface {
	GetGroups(input []string) [][]string
}

type dataHandler struct {
	groupsMapByAccount map[string]int64
	accountsMapByGroup map[int64]map[string]bool
	currentGroup       int64
}

func New() Handler {
	return &dataHandler{
		groupsMapByAccount: nil,
		accountsMapByGroup: nil,
		currentGroup:       0,
	}
}

func (h *dataHandler) GetGroups(input []string) (res [][]string) {
	h.groupsMapByAccount = make(map[string]int64)
	h.accountsMapByGroup = make(map[int64]map[string]bool)
	h.currentGroup = 0

	for _, v := range input {
		h.GroupIDs(v)
	}
	res = make([][]string, 0)
	for _, v := range h.accountsMapByGroup {
		if len(v) > 1 {
			group := make([]string, 0)
			for acc, _ := range v {
				group = append(group, acc)
			}
			res = append(res, group)
		}
	}

	return res
}

func (h *dataHandler) GroupIDs(input string) {
	var mergeGroup int64
	var currentGroup int64
	accountsArray := strings.Split(input, ",")
	groups := make([]int64, 0)

	for _, account := range accountsArray {
		group, ok := h.groupsMapByAccount[account]
		if !ok {
			group = currentGroup
			if group == 0 {
				group = h.newGroup()
			}
			h.groupsMapByAccount[account] = group
			if h.accountsMapByGroup[group] == nil {
				h.accountsMapByGroup[group] = make(map[string]bool)
			}
			h.accountsMapByGroup[group][account] = true
			currentGroup = group
		}
		if ok && mergeGroup == 0 {
			mergeGroup = group
		}
		groups = append(groups, group)
	}

	if mergeGroup != 0 {
		h.mergeGroups(mergeGroup, groups)
	}
}

func (h *dataHandler) mergeGroups(targetGroup int64, groups []int64) {
	for _, group := range groups {
		if group == targetGroup {
			continue
		}
		for account, _ := range h.accountsMapByGroup[group] {
			h.accountsMapByGroup[targetGroup][account] = true
			h.groupsMapByAccount[account] = targetGroup
		}
		delete(h.accountsMapByGroup, group)
	}
}

func (h *dataHandler) newGroup() int64 {
	h.currentGroup++
	return h.currentGroup
}
