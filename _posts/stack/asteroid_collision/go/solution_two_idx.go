package main

func asteroidCollisionTwoIdx(asteroids []int) []int {
	i, j := 0, 1

	if len(asteroids) == 2 {
		first := asteroids[i]
		second := asteroids[j]

		if first*second > 0 || first < 0 {
			return asteroids
		}

		if first == -1*second {
			return []int{}
		}

		if first > -1*second {
			return []int{first}
		}

		return []int{second}
	}

	for i >= 0 && j < len(asteroids) {
		first := asteroids[i]
		second := asteroids[j]

		if first*second > 0 || first < 0 {
			i++
			j++
			continue
		}

		if first > -1*second {
			asteroids = append(asteroids[:i+1], asteroids[j+1:]...)
			continue
		}

		if first < -1*second {
			asteroids = append(asteroids[:i], asteroids[j:]...)
			j--
			i--
			if i < 0 {
				i = 0
				j = 1
			}

			continue
		}

		if first == -1*second {
			asteroids = append(asteroids[:i], asteroids[j+1:]...)
			j--
			i--
			if i < 0 {
				i = 0
				j = 1
			}
			continue
		}
	}

	return asteroids
}
