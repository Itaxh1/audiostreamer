http.HandleFunc("/audio", func(w http.ResponseWriter, r *http.Request) {
  flusher, ok := w.(http.Flusher)
  if !ok {
   panic("expected http.ResponseWriter to be an http.Flusher")
  }  w.Header().Set("Connection", "Keep-Alive")
  w.Header().Set("Transfer-Encoding", "chunked")  for true {
   binary.Write(w, binary.BigEndian, &buffer)
   flusher.Flush() // Trigger "chunked" encoding
   return
  }
 })

