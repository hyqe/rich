package phone

type Number struct {
	Country_code string
	Area_code    string
	Phone_number string
}

func (n Number) String() string {
	return "+" + n.Country_code + n.Area_code + n.Phone_number
}

//^\+[1-9]\d{1,14}$
