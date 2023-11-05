package calculations

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

func CalculateFactorial(number int64, logging bool) (int64, error) {
	if number < 0 {
		return 0, errors.New("number should be not negative")
	}

	if logging {
		factorialLog := log.WithFields(log.Fields{
			"number":  number,
			"logging": logging,
		})
		factorialLog.Info("Start calculations...")
		factorialLog.Infof("Calculate %d!", number)
		defer factorialLog.Info("Calculations complete!")
	}

	return factorial(number), nil
}

func factorial(number int64) int64 {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
