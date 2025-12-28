package data

type Location string

// func DistanceTo(origin location,  destination location) distance{
// 	// TODO claculatyion
// 	return 10
// }
// this method is associated with Location type and (origin Location) - receiver
func (origin Location) DistanceTo(destination Location) distance{
	// TODO calculation
	return 10
}

func locationTest(){
	nyc := Location("333.4, 34.")
	nyc.DistanceTo(Location("-23, -44"))

	print(nyc) 
}


type dis int // creating a new type
type json = map[string]string // it's just an alias, not changing the nature of type

type distance float64 //mile // creating a new type based on other type
type distancekm float64

// global function
// func ToKm(miles distance) distancekm {
// 	return distancekm(1.6 * miles)
// }

//method
func (miles distance) ToKm() distancekm {
	return distancekm(1.6 * miles)
}
func (km distancekm) ToMiles() distance {
	return distance(km / 1.6 )
}

func test(){
	d := distance(34.5)
	km := d.ToKm()
	km.ToMiles()
	print(d)
}