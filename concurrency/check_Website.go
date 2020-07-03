package concurrency



type WebChecker func(string) bool

type result struct{
	string
	bool
}
func CheckWebsite(wc WebChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel	<- result{u, wc(u)}
		}(url)
	}
	for i:=0;i<len(urls);i++{
		result := <- resultChannel
		res[result.string] = result.bool
	}
	return res
}