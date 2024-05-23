package securefilechanger

import "errors"

type File struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"file_name" db:"name"`
	Path      string `json:"path" db:"path"`
	SizeBytes int    `json:"size_bytes" db:"size_bytes"`
	Type      string `json:"type" db:"type"`
	FolderId  int    `json:"folder_id" db:"folder_id"`
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
