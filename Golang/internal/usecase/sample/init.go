package sample

func New(sampleRepo sampleResource) *sampleUsecase {
	return &sampleUsecase{
		sampleRepo: sampleRepo,
	}
}
