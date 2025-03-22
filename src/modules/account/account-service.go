package accountModule

import (
	errorHelpers "go-gin-test-job/src/common/error-helpers"
	"go-gin-test-job/src/database"
	"go-gin-test-job/src/database/entities"
	accountModuleDto "go-gin-test-job/src/modules/account/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getAccounts(status entities.AccountStatus, orderParams map[string]string, offset int, count int, searchParams map[string]string) ([]*entities.Account, int64) {
	return database.GetAccountsAndTotal(status, orderParams, offset, count, searchParams)
}

func createAccount(c *gin.Context, dto *accountModuleDto.PostCreateAccountRequestDto) (*entities.Account, error) {
	var account *entities.Account
	transactionError := database.DbConn.Transaction(func(tx *gorm.DB) error {
		if database.IsAddressExists(tx, dto.Address) {
			return errorHelpers.RespondConflictError(c, "Address already exists")
		}
		// newAccount, err := database.CreateAccount(tx, entities.CreateAccount(dto.Address, dto.Status))
		newAccount, err := database.CreateAccount(tx, &entities.Account{
			Address: dto.Address,
			Status:  dto.Status,
			Name:    dto.Name,
			Ranking: *dto.Ranking,
			Memo:    dto.Memo,
		})
		if err != nil {
			return err
		}
		account = newAccount
		return nil
	}, database.DefaultTxOptions)
	if transactionError != nil {
		return nil, transactionError
	}
	return account, nil
}
