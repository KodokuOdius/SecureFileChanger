package securefilechanger

import "errors"

type File struct {
	Id        int    `json:"file_id" db:"id"`
	Name      string `json:"file_name" db:"name"`
	Path      string `json:"-" db:"path"`
	SizeBytes int    `json:"size_bytes" db:"size_bytes"`
	Type      string `json:"file_type" db:"type"`
	FolderId  int    `json:"-" db:"folder_id"`
	OwnerId   int    `json:"-" db:"user_id"`
}

type UpdateFile struct {
	Name *string `json:"file_name"`
}

type Folder struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"folder_name" db:"name" binding:"required"`
	Is_root bool   `json:"is_root" db:"is_root"`
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
