package queryParameter

import (
	"forgoproject/pkg/type/pagination"
	"forgoproject/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	/*Тут можно добавить фильтр*/
}
