package main

import "fmt"

type Char struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func main() {
	profile := MakeProfile("Shoto")
	fmt.Println("Name: ", profile.Name)
	fmt.Println("Health: ", profile.Health)
	fmt.Println("Power: ", profile.Power)
	fmt.Println("Exp: ", profile.Exp)
	fmt.Println("=== Heroes Power Up ===")
	powerUp := PowerUp(profile, 3)
	fmt.Println("Name: ", powerUp.Name)
	fmt.Println("Health: ", powerUp.Health)
	fmt.Println("Power: ", powerUp.Power)
	fmt.Println("Exp: ", powerUp.Exp)
}

func MakeProfile(name string) Char {
	var profile Char

	if name == "Midoriya" {
		profile.Name = "Midoriya"
		profile.Health = 500
		profile.Power = 250
		profile.Exp = 50
	}
	if name == "Bakugo" {
		profile.Name = "Bakugo"
		profile.Health = 600
		profile.Power = 500
		profile.Exp = 200
	}
	if name == "Shoto" {
		profile.Name = "Shoto"
		profile.Health = 600
		profile.Power = 700
		profile.Exp = 150
	}

	return profile
}

func PowerUp(p Char, i int) Char {
	p.Health += (p.Health * i)
	p.Power += (p.Power * i)
	p.Exp += (p.Exp * i)
	return p
}
