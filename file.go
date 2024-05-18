package securefilechanger

type File struct {
	Id   string `json:"-"`
	Name string `json:"file_name"`
}

type Folder struct {
	Id   string `json:"-"`
	Name string `json:"file_name"`
}
