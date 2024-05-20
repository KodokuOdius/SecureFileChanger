package securefilechanger

import "errors"

type File struct {
	Id        int    `json:"-"`
	Name      string `json:"file_name"`
	Path      string `json:"path"`
	SizeBytes string `json:"size_bytes"`
	Type      string `json:"type"`
	FolderId  string `json:"folder_id"`
}

type Folder struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"folder_name" db:"name" binding:"required"`
	Is_root bool   `json:"is_root" db:"is_root"`
	Is_bin  bool   `json:"is_bin" db:"is_bin"`
}

type UpdateFolder struct {
	Name *string `json:"folder_name"`
}

func (u UpdateFolder) Validate() error {
	if u.Name == nil {
		return errors.New("update has no values")
	}

	return nil
}
