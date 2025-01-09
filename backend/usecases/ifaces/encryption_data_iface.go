package ifaces

import (
	"context"
	model "decrypter-timer/domain/models"
)

type EncryptionProcedureInterface interface {
	EncryptMessage(ctx context.Context, encryptData *model.EncryptData) (*model.EncryptData, error)
	DecryptMessage(ctx context.Context, decryptData *model.EncryptData) (*model.EncryptData, error)
}
