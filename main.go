package main

import (
	"fmt"
)

type The struct {
	Best string
}
type Are struct {
	The The
}
type We struct {
	Are Are
}

// ----------------
type Hello struct {
	World string
}

// -----------------
type Tech struct {
	Academy string
}
type Man struct {
	Tech Tech
}
type Data struct {
	Man []Man
}
type Object struct {
	Str [][][]Data
}

// -----------------
type Fruit struct {
	Is string
}
type Favourite struct {
	Fruit Fruit
}
type My struct {
	Favourite []Favourite
}

// -----------------------
type Number struct {
	First   []int
	Secound []int
}

func main() {
	We := We{
		Are: Are{
			The: The{
				Best: "KODA",
			},
		},
	}
	fmt.Println("1. We.Are.The.Best : ", We.Are.The.Best)
	//---------------------------------------------------------
	hello := Hello{
		World: "Hello World",
	}
	fmt.Println("2. hello.world : ", hello.World)
	// ---------------------------------------------------------
	obj := Object{
		Str: make([][][]Data, 4),
	}
	obj.Str[3] = make([][]Data, 2)
	obj.Str[3][1] = make([]Data, 3)

	obj.Str[3][1][2] = Data{
		Man: []Man{
			{
				Tech: Tech{
					Academy: "Tech Academy",
				},
			},
		},
	}
	fmt.Println("3. obj.str[3][1][2].man[0].tech.academy : ", obj.Str[3][1][2].Man[0].Tech.Academy)
	//----------------------------------------------------------------------------------------------
	my := []My{
		{
			Favourite: []Favourite{
				{},
				{},
				{},
				{
					Fruit: Fruit{
						Is: "Apple",
					},
				},
			},
		},
	}
	fmt.Println("4. my[0].Favourite[3].Fruit.Is : ", my[0].Favourite[3].Fruit.Is)
	//---------------------------------------------------------------------------
	num := Number{
		First:   []int{1, 12},
		Secound: []int{5, 6, 20},
	}
	fmt.Println("5. num.First[1] + num.Secound[2] : ", num.First[1]+num.Secound[2])
}
