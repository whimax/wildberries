package types

// MonthDetailSale Структура данных отчёта валберис "Ежемесячный отчет о продажах по реализации"
type MonthDetailSale struct {
	RealizationreportId      uint64          `json:"realizationreport_id"`
	DateFrom                 WildberriesTime `json:"date_from"`
	DateTo                   WildberriesTime `json:"date_to"`
	CreateDt                 WildberriesTime `json:"create_dt"`
	SuppliercontractCode     interface{}     `json:"suppliercontract_code"`
	RrdId                    uint64          `json:"rrd_id"`
	GiId                     uint64          `json:"gi_id"`
	SubjectName              string          `json:"subject_name"`
	NmId                     uint64          `json:"nm_id"`
	BrandName                string          `json:"brand_name"`
	SaName                   string          `json:"sa_name"`
	TsName                   string          `json:"ts_name"`
	Barcode                  string          `json:"barcode"`
	DocTypeName              string          `json:"doc_type_name"`
	Quantity                 int             `json:"quantity"`
	RetailPrice              float64         `json:"retail_price"`
	RetailAmount             float64         `json:"retail_amount"`
	SalePercent              float64         `json:"sale_percent"`
	CommissionPercent        float64         `json:"commission_percent"`
	OfficeName               string          `json:"office_name"`
	SupplierOperName         string          `json:"supplier_oper_name"`
	OrderDt                  WildberriesTime `json:"order_dt"`
	SaleDt                   WildberriesTime `json:"sale_dt"`
	RrDt                     WildberriesTime `json:"rr_dt"`
	ShkId                    int64           `json:"shk_id"`
	RetailPriceWithdiscRub   float64         `json:"retail_price_withdisc_rub"`
	DeliveryAmount           uint64          `json:"delivery_amount"`
	ReturnAmount             uint64          `json:"return_amount"`
	DeliveryRub              float64         `json:"delivery_rub"`
	GiBoxTypeName            string          `json:"gi_box_type_name"`
	ProductDiscountForReport float64         `json:"product_discount_for_report"`
	SupplierPromo            float64         `json:"supplier_promo"`
	Rid                      int64           `json:"rid"`
	PpvzSppPrc               float64         `json:"ppvz_spp_prc"`
	PpvzKvwPrcBase           float64         `json:"ppvz_kvw_prc_base"`
	PpvzKvwPrc               float64         `json:"ppvz_kvw_prc"`
	PpvzSalesCommission      float64         `json:"ppvz_sales_commission"`
	PpvzForPay               float64         `json:"ppvz_for_pay"`
	PpvzReward               float64         `json:"ppvz_reward"`
	AcquiringFee             float64         `json:"acquiring_fee"`
	AcquiringBank            string          `json:"acquiring_bank"`
	PpvzVw                   float64         `json:"ppvz_vw"`
	PpvzVwNds                float64         `json:"ppvz_vw_nds"`
	PpvzOfficeId             uint64          `json:"ppvz_office_id"`
	PpvzOfficeName           string          `json:"ppvz_office_name"`
	PpvzSupplierId           uint64          `json:"ppvz_supplier_id"`
	PpvzSupplierName         string          `json:"ppvz_supplier_name"`
	PpvzInn                  string          `json:"ppvz_inn"`
	DeclarationNumber        string          `json:"declaration_number"`
	BonusTypeName            string          `json:"bonus_type_name"`
	StickerId                string          `json:"sticker_id"`
	SiteCountry              string          `json:"site_country"`
	Penalty                  float64         `json:"penalty"`
	AdditionalPayment        int             `json:"additional_payment"`
	Kiz                      string          `json:"kiz"`
	Srid                     string          `json:"srid"`
}
