package filters

// Filters фильтр поиска
type Filters struct {
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	Nation     string
	Limit      int
	Offset     int
}
