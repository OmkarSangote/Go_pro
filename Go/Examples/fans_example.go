/*
	A More Practical Example:

Imagine you have two data sources (like two APIs or file readers), and you want to gather data from both of them. You can use Fan-Out to fetch data concurrently and Fan-In to merge the results into one stream for further processing.

// ..............Pseudocode Example...................
*/
package main

func readAPI(apiURL string, ch chan<- Data) {
	// Fetch data from an API
	ch <- fetchDataFromAPI(apiURL)
}

func fanIn(ch1, ch2 <-chan Data) <-chan Data {
	merged := make(chan Data)
	go func() {
		for {
			select {
			case v := <-ch1:
				merged <- v
			case v := <-ch2:
				merged <- v
			}
		}
	}()
	return merged
}

func main() {
	api1 := make(chan Data)
	api2 := make(chan Data)

	go readAPI("http://api1.com", api1)
	go readAPI("http://api2.com", api2)

	mergedChannel := fanIn(api1, api2)

	// Handle merged data stream
	for data := range mergedChannel {
		process(data)
	}
}
