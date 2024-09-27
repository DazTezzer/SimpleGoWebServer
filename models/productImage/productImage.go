package productImage

type ProductImage struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image []byte `json:"image"` // или string для хранения URL
}
