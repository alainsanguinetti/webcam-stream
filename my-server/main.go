package main

import "net/http"

func main() {
    http.ListenAndServe("/", http.FileServer(http.Dir(".")))
        
    http.HandleFunc( ":4242/mjpeg", serveMjpeg )  
      
}

func serveMjpeg( w http.ResponseWriter, r * http.Request ) {
    w.Header().Set( "Content-Type:", "video/x-motion-jpeg" )

    http.ServeFile(w, r, "./shot.jpeg" )
}
