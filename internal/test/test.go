package test

/*
GWPM = totalChars / duration * 12
NWPM = (totalChars - mistakes) / duration * 12
Accuracy = (totalChars - mistakes) / totalChars * 100
*/

type Result struct {
	GWPM     float64
	NWPM     float64
	Accuracy float64
}

func Calculate(totalChars, mistakes, duration int) Result {
	if duration <= 0 || totalChars <= 0 {
		return Result{}
	}

	gross := float64(totalChars) / float64(duration) * 12
	net := float64(totalChars-mistakes) / float64(duration) * 12

	if net < 0 {
		net = 0
	}

	accuracy := float64(totalChars-mistakes) / float64(totalChars) * 100

	if accuracy < 0 {
		accuracy = 0
	}

	return Result{
		GWPM:     gross,
		NWPM:     net,
		Accuracy: accuracy,
	}
}
