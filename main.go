package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func consultarServicio(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("-> Tarea %d terminada\n", id)
}

func handler(w http.ResponseWriter, r *http.Request) {
	inicio := time.Now()
	var wg sync.WaitGroup

	for i := 1; i < 3; i++ {
		wg.Add(1)
		go consultarServicio(i, &wg)
	}

	wg.Wait()

	duracion := time.Since(inicio).Round(time.Millisecond)
	fmt.Fprintf(w, "3 tareas completadas en: %v (Secuencial serian 3s)\n", duracion)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println(" Servidor esta en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
