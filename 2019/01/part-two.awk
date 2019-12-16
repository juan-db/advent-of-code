#!/bin/awk -f
	{ 
		fuel = $1
		while ((fuel = int(fuel / 3) - 2) > 0) {
			total += fuel
		}
	}
END	{ print total }
