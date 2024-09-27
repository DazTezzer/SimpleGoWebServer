package productDescription

type ProductDescription struct {
	Id             uint   `json:"id"`
	Size           string `json:"size"`
	Color          string `json:"color"`
	Material       string `json:"material"`
	Gender         string `json:"gender"`
	AdditionalInfo string `json:"additional_info"`
}
