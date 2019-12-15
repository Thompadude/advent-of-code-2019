package main

func getFuelRequirements(masses []int) int {
	var total int
	for _, mass := range masses {
		total += getFuelRequirement(mass)
	}
	return total
}

func getFuelRequirement(mass int) int {
	return mass/3 - 2
}
