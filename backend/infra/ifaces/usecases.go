package ifaces

import (
	"context"
	model "decrypter-timer/domain/models"
	"decrypter-timer/usecases"
)

type DecrypterTimeUseCaseInterface interface {
	Encrypt(ctx context.Context, in *usecases.EncryptRequest) (*model.EncryptData, error)
	Decrypt(ctx context.Context, in *usecases.DecryptRequest) (*model.EncryptData, error)
}
