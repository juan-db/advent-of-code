#!/bin/awk -f
	{ 
		fuel = $1
		while ((fuel = int(fuel / 3) - 2) > 0) {
			print "fuel " fuel
			total += fuel
			print "total " total
		}
	}
END	{ print total }
