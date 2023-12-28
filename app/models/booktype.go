package models

type (
	BookRequestType struct {
		Title  string `json:"title" binding:"required" validate:"min=1,max=255"`
		Author string `json:"author" binding:"required" validate:"min=1,max=255"`
	}

	BookPaginationResponseType struct {
		BaseErrorType
		Data     any    `json:"data"`
		Page     *int   `json:"page,omitempty"`
		Size     *int   `json:"size,omitempty"`
		RowCount *int64 `json:"row_count,omitempty"`
	}

	BookResponseType struct {
		BaseErrorType
		Data any `json:"data,omitempty"`
	}

	ResourcePaginationConfigType struct {
		PageText    string
		SizeText    string
		Page        string
		PageSize    string
		MinPageSize int
		MaxPageSize int
	}
)
