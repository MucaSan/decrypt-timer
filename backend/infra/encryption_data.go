package infrastructure

import (
	"context"
	model "decrypter-timer/domain/models"
	"decrypter-timer/usecases/ifaces"
)

func NewEncryptionProcedure() ifaces.EncryptionProcedureInterface {
	return &EncryptionProcedure{}
}

type EncryptionProcedure struct{}

func (ep *EncryptionProcedure) EncryptMessage(ctx context.Context, encryptData *model.EncryptData) (*model.EncryptData, error) {

}

func (ep *EncryptionProcedure) DecryptMessage(ctx context.Context, decryptData *model.EncryptData) (*model.EncryptData, error) {

}
