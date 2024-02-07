package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    foo := make(map[string]int)
    foo["quarter_of_a_dozen"] = 3
    foo["half_of_a_dozen"] = 6
    foo["dozen"] = 12
    foo["small_gross"] = 120
    foo["gross"] = 144
    foo["great_gross"] = 1728
    return foo
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bar := make(map[string]int)
    return bar
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    if amount, ok := units[unit]; ok {
        bill[item] += amount
        return true
    } else {
        return false
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    if _, ok := bill[item]; !ok {
        return false
    }

    if amount, ok := units[unit]; ok {
        if bill[item] - amount < 0 {
            return false
        } else if bill[item] - amount == 0 {
            delete(bill, item)
        } else {
            bill[item] -= amount
        }
        return true
    } else {
        return false
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    if _, ok := bill[item]; ok {
        return bill[item], true 
    }
    return 0, false
}
