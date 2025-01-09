package usecases

import (
	"context"
	model "decrypter-timer/domain/models"
)

type EncryptRequest struct {
	Message   string
	Algorithm string
	SecretKey string
}

func (uc *UseCasesContainer) Encrypt(ctx context.Context, in *EncryptRequest) (*model.EncryptData, error) {
	encryptData := &model.EncryptData{
		Algorithm: in.Algorithm,
		Message:   in.Message,
		SecretKey: in.SecretKey,
	}

}
