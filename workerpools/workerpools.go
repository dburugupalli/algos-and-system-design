package main 

import(
 "fmt", 
 "log"
)


type Site struct {
  URL string 
}

type Result struct {
  Status int
}

func crawl(wId int, jobs <- chan Site, results chan <- Result){
      for site := range jobs {
        log.Printf("Worker ID: %d\n", wId)
        resp, err := http.Get(site.URL)
        if err  != nil {
          log.Println(err.Error())
        }
        result <- Result{Status:resp.StatusCode}
      }
}

func main(){
  fmt.Println("worker-pools in Go")
  jobs := make(chan Site, 3)
  result := make(chan Result, 3)  
  for w :=1; w<=3; w++ {
      go crawl(w, jobs, results)
  }
  
  urls := []string{
   "https://tutorialedge.net"
   "https://tutorialedge.net/pricing/"
   "https://example.com"
   "https://google.com"
  }
  
  for _, url := range url{
    jobs <- Site{URL:url}
  }
  
  close(jobs)
}

