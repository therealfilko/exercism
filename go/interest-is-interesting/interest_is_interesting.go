package interest

// Definition der Zinssätze als Konstanten
const (
    InterestRateNegativeBalance float32 = 3.213
    InterestRateUnder1000       float32 = 0.5
    InterestRateUnder5000       float32 = 1.621
    InterestRateAbove5000       float32 = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    switch {
    case balance < 0:
        return InterestRateNegativeBalance
    case balance >= 0 && balance < 1000:
        return InterestRateUnder1000
    case balance >= 1000 && balance < 5000:
        return InterestRateUnder5000
    default:
        return InterestRateAbove5000
    }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    switch {
    case balance < 0:
        return balance * float64(InterestRateNegativeBalance) / 100
    case balance >= 0 && balance < 1000:
        return balance * float64(InterestRateUnder1000) / 100
    case balance >= 1000 && balance < 5000:
        return balance * float64(InterestRateUnder5000) / 100
    default:
        return balance * float64(InterestRateAbove5000) / 100
    }
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    switch {
    case balance < 0:
        return balance + (balance * float64(InterestRateNegativeBalance) / 100)
    case balance >= 0 && balance < 1000:
        return balance + (balance * float64(InterestRateUnder1000) / 100)
    case balance >= 1000 && balance < 5000:
        return balance + (balance * float64(InterestRateUnder5000) / 100)
    default:
        return balance + (balance * float64(InterestRateAbove5000) / 100)
    }
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    years := 0
    for balance < targetBalance {
        // Berechne den Zinssatz basierend auf dem aktuellen Kontostand
        var interestRate float64
        switch {
        case balance < 0:
            interestRate = float64(InterestRateNegativeBalance) / 100
        case balance < 1000:
            interestRate = float64(InterestRateUnder1000) / 100
        case balance < 5000:
            interestRate = float64(InterestRateUnder5000) / 100
        default:
            interestRate = float64(InterestRateAbove5000) / 100
        }
        
        // Aktualisiere den Kontostand mit den Zinsen für dieses Jahr
        balance += balance * interestRate
        
        // Zähle dieses Jahr
        years++
    }
    return years
}
