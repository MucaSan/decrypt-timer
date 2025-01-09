package usecases

import (
	"context"
	model "decrypter-timer/domain/models"
)

type DecryptRequest struct {
	Message   string
	Algorithm string
	SecretKey string
}

func (uc *UseCasesContainer) Decrypt(ctx context.Context, in *DecryptRequest) (*model.EncryptData, error) {
	return nil, nil
}
