package pattern

import "fmt"

/*
Паттерн "фасад" это абстракция, аккумулирующая в себе высокоуровневый набор операций 
для работы с некоторой сложной подсистемой. 
Фасад агрегирует классы, реализующие функциональность этой подсистемы, но не скрывает их. 
Важно понимать, что клиент, при этом, не лишается более низкоуровневого доступа к классам подсистемы, 
если ему это, конечно, необходимо. Фасад упрощает выполнение некоторых операций с подсистемой, 
но не навязывает клиенту свое использование.
*/

//внешний тип wallet
type Wallet struct {
    accountId *int
    holderName *string
    balance *int
}

//фасад, дополняющий тип wallet
type WalletFacade struct {
	wallet Wallet
}


func (wf *WalletFacade) subtructMoneyFromWallet(accountID int, amount int) error {
    if (accountID != *wf.wallet.accountId) {
		return fmt.Errorf("wrong account id")
	}
	if (*wf.wallet.balance > amount) {
		return fmt.Errorf("insufficent funds")
	}
	*wf.wallet.balance = *wf.wallet.balance - amount
    return nil
}
