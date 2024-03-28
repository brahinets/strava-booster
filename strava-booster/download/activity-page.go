package download

type ActivitiesPage struct {
	Activities []ActivityEntity `json:"models"`
	Page       int              `json:"page"`
	PerPage    int              `json:"perPage"`
	Total      int              `json:"total"`
}
