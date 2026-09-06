package blackjack

func ParseCard(card string) int {
	switch card {
    case "ace":
    	return 11
    case "two":
	   	return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten", "jack", "queen", "king":
        return 10
    }

    return 0
}

func FirstTurn(card1, card2, dealerCard string) string {
    hand := ParseCard(card1) + ParseCard(card2)
    
	switch {
    case card1 == "ace" && card2 == "ace":
        return "P"
    case hand == 21 && ParseCard(dealerCard) < 10:
        return "W"
    case hand >= 17:
        return "S"
    case hand >= 12 && ParseCard(dealerCard) > 6:
        return "H"
    case hand >= 12:
        return "S"
    }

    return "H"
}
