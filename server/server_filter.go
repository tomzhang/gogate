package server

import (
	"github.com/alecthomas/log4go"
)

// 注册过滤器, 追加到末尾
func (s *Server) AppendPreFilter(pre *PreFilter) {
	log4go.Info("append pre filter: %s", pre.Name)
	s.preFilters = append(s.preFilters, pre)
}

// 注册过滤器, 追加到末尾
func (s *Server) AppendPostFilter(post *PostFilter) {
	log4go.Info("append post filter: %s", post.Name)
	s.postFilters = append(s.postFilters, post)
}

func (s *Server) ExportAllPreFilters() []*PreFilter {
	result := make([]*PreFilter, len(s.preFilters))
	copy(result, s.preFilters)

	return result
}

func (s *Server) ExportAllPostFilters() []*PostFilter {
	result := make([]*PostFilter, len(s.postFilters))
	copy(result, s.postFilters)

	return result
}

// 在指定前置过滤器的后面添加
func (serv *Server) InsertPreFilterBehind(filterName string, filter *PreFilter) bool {
	log4go.Info("insert pre filter: %s", filter.Name)
	targetIdx := serv.getPreFilterIndex(filterName)
	if -1 == targetIdx {
		return false
	}

	serv.ensurePreFilterCap(1)

	// move elem
	size := len(serv.preFilters)
	ix := size - 1
	for ; ix > targetIdx; ix-- {
		serv.preFilters[ix + 1] = serv.preFilters[ix]
	}

	serv.preFilters[ix] = filter

	return true
}

// 在指定后置过滤器的后面添加
func (serv *Server) InsertPostFilterBehind(filterName string, filter *PostFilter) bool {
	log4go.Info("insert post filter: %s", filter.Name)
	targetIdx := serv.getPostFilterIndex(filterName)
	if -1 == targetIdx {
		return false
	}

	serv.ensurePostFilterCap(1)

	// move elem
	size := len(serv.postFilters)
	ix := size - 1
	for ; ix > targetIdx; ix-- {
		serv.postFilters[ix + 1] = serv.postFilters[ix]
	}

	serv.postFilters[ix] = filter

	return true
}

func (serv *Server) ensurePreFilterCap(neededSpace int) {
	currentCap := cap(serv.preFilters)
	currentLen := len(serv.preFilters)
	leftSpace := currentCap - currentLen

	if leftSpace < neededSpace {
		newCap := currentCap + (neededSpace - leftSpace) + 3

		oldFilters := serv.preFilters
		serv.preFilters = make([]*PreFilter, 0, newCap)
		copy(serv.preFilters, oldFilters)
	}
}

func (serv *Server) getPreFilterIndex(name string) int {
	if nil == serv.preFilters {
		return -1
	}

	for ix, f := range serv.preFilters {
		if f.Name == name {
			return ix
		}
	}

	return -1
}




func (serv *Server) ensurePostFilterCap(neededSpace int) {
	currentCap := cap(serv.postFilters)
	currentLen := len(serv.postFilters)
	leftSpace := currentCap - currentLen

	if leftSpace < neededSpace {
		newCap := currentCap + (neededSpace - leftSpace) + 3

		oldFilters := serv.postFilters
		serv.postFilters = make([]*PostFilter, 0, newCap)
		copy(serv.postFilters, oldFilters)
	}
}

func (serv *Server) getPostFilterIndex(name string) int {
	if nil == serv.preFilters {
		return -1
	}

	for ix, f := range serv.postFilters {
		if f.Name == name {
			return ix
		}
	}

	return -1
}

