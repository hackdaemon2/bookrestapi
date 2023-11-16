package models

type (
	BookRequestDTO struct {
		Title  string `json:"title" binding:"required" validate:"min=1,max=255"`
		Author string `json:"author" binding:"required" validate:"min=1,max=255"`
	}

	BookPaginationResponseDTO struct {
		BaseErrorDTO
		Data     any    `json:"data"`
		Page     *int   `json:"page,omitempty"`
		Size     *int   `json:"size,omitempty"`
		RowCount *int64 `json:"row_count,omitempty"`
	}

	BookResponseDTO struct {
		BaseErrorDTO
		Data any `json:"data,omitempty"`
	}

	ResourcePaginationConfigDTO struct {
		PageText    string
		SizeText    string
		Page        string
		PageSize    string
		MinPageSize int
		MaxPageSize int
	}
)
