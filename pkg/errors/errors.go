package artifactsErrors

type ErrorResponseSchema struct {
	Error ErrorSchema `json:"error"`
}

type ErrorSchema struct {
	Code    int    `json:"code"`
	Message string `json:"string"`
}

type ArtifactsMmoError struct {
	ErrorSchema ErrorSchema
}

func (error *ArtifactsMmoError) Error() string {
	return error.ErrorSchema.Message
}

const (
	// General
	InvalidPayload  int = 422
	TooManyRequests int = 429
	NotFound        int = 404
	FatalError      int = 500

	// intAccount Error Codes
	TokenInvalid           int = 452
	TokenExpired           int = 453
	TokenMissing           int = 454
	TokenGenerationFail    int = 455
	UsernameAlreadyUsed    int = 456
	EmailAlreadyUsed       int = 457
	SamePassword           int = 458
	CurrentPasswordInvalid int = 459

	// intCharacter Error Codes
	CharacterNotEnoughHp            int = 483
	CharacterMaximumUtilitesEquiped int = 484
	CharacterItemAlreadyEquiped     int = 485
	CharacterLocked                 int = 486
	CharacterNotThisTask            int = 474
	CharacterTooManyItemsTask       int = 475
	CharacterNoTask                 int = 487
	CharacterTaskNotCompleted       int = 488
	CharacterAlreadyTask            int = 489
	CharacterAlreadyMap             int = 490
	CharacterSlotEquipmentError     int = 491
	CharacterGoldInsufficient       int = 492
	CharacterNotSkillLevelRequired  int = 493
	CharacterNameAlreadyUsed        int = 494
	MaxCharactersReached            int = 495
	CharacterNotLevelRequired       int = 496
	CharacterInventoryFull          int = 497
	CharacterNotFound               int = 498
	CharacterInCooldown             int = 499

	// intItem Error Codes
	ItemInsufficientQuantity int = 471
	ItemInvalidEquipment     int = 472
	ItemRecyclingInvalidItem int = 473
	ItemInvalidConsumable    int = 476
	MissingItem              int = 478

	// intGrand Exchange Error Codes
	GeMaxQuantity           int = 479
	GeNotInStock            int = 480
	GeNotThePrice           int = 482
	GeTransactionInProgress int = 436
	GeNoOrders              int = 431
	GeMaxOrders             int = 433
	GeTooManyItems          int = 434
	GeSameAccount           int = 435
	GeInvalidItem           int = 437
	GeNotYourOrder          int = 438

	// intBank Error Codes
	BankInsufficientGold      int = 460
	BankTransactionInProgress int = 461
	BankFull                  int = 462

	// intMaps Error Codes
	MapNotFound        int = 597
	MapContentNotFound int = 598
)
