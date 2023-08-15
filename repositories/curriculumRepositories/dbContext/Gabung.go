package dbContext

type GetGabung struct {
	Createprogram_entityParams Createprogram_entityParams
	CreatesectionsParams       CreatesectionsParams
}

type SectionGroup struct {
	CreatesectionsParams
}

type CreateGabungParams struct {
	Createprogram_entityParams Createprogram_entityParams
	CreatesectionsParams       CreatesectionsParams
	CreateCategoryParams       CreateCategoryParams
	CreateProgEntityDescParams CreateProgEntityDescParams
	Createsection_detailParams Createsection_detailParams
	// CreatesectionDetailMaterialParams CreatesectionDetailMaterialParams
}
