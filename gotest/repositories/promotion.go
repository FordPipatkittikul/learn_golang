package repositories

type Promotion struct {
	ID				int
	PurchaseMin 	int
	DiscountPercent int
}

type PromotionRepository interface {
	GetPromotion() (Promotion, error)
}

