package main

import "math"
import "sort"

type Transaction struct {
    Lender string
    Borrowers []string
    Amount float64
}

type IndividualBal struct {
    Id string
    Balance float64
}
type Settlement struct {
    Sender string
    Receiver string
    Amount float64
}

func Settle(transactions []Transaction) []Settlement {
    usersBal := make(map[string]float64)
    for i := 0; i<len(transactions); i++ {
        currTrans := transactions[i]
        payment := currTrans.Amount
        borrowers := currTrans.Borrowers

        usersBal[currTrans.Lender] += (float64(len(borrowers))*payment)/(float64(len(borrowers))+1)
        for j:=0; j<len(borrowers); j++ {
            usersBal[borrowers[j]] -= payment/(float64(len(borrowers))+1)
        }
    }

    userBalArr := make([]IndividualBal, 0)
    for i, v := range usersBal {
        userBalArr = append(userBalArr, IndividualBal{Id: i, Balance: v})
    }

    sort.Slice(userBalArr, func(i, j int) bool {
      return userBalArr[i].Balance < userBalArr[j].Balance
    })

    settlements := make([]Settlement, 0)
    i, j := 0, len(userBalArr)-1
    for i<j {
        settlementAmount := math.Min(math.Abs(userBalArr[i].Balance), userBalArr[j].Balance)

        settlement := Settlement{
            Sender: userBalArr[i].Id,
            Receiver: userBalArr[j].Id,
            Amount: settlementAmount,
        }
        settlements = append(settlements, settlement)

        userBalArr[i].Balance += settlementAmount
        userBalArr[j].Balance -= settlementAmount
        if(userBalArr[i].Balance==0){
            i++
        }
        if(userBalArr[j].Balance==0){
            j--
        }
    }

    return settlements
}
