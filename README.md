# 🚀 Píldora Formativa: Introducción a Go (Golang)

> **P1 CRC (250h) - Píldora Formativa | Factoría F5**  
> **Tema:** Lenguaje Go (Golang), Filosofía, Concurrencia Nativa y Comparativa técnica frente a Python.  
> **Duración estimada:** ~20-25 minutos (15 min teoría + 8 min live coding + 2 min conclusiones).

---

## 📌 1. Descripción del Tema

**Go (Golang)** es un lenguaje de programación de código abierto desarrollado por ingenieros de Google en 2007 (Robert Griesemer, Rob Pike y Ken Thompson) para resolver los desafíos de escalabilidad, compilación lenta y desarrollo masivo en sistemas de red modernos.

### 🌟 Pilares Fundamentales
1. **Simplicidad y Minimalismo:** Solo 25 palabras reservadas. Un único tipo de bucle (`for`). Cero jerarquías complejas de herencia.
2. **Rendimiento Nativo:** Compila directamente a código máquina en un único binario ejecutable estático con arranque instantáneo.
3. **Concurrencia de Primera Clase:** Implementa el modelo CSP (*Communicating Sequential Processes*) mediante **Goroutines** (hilos virtuales ligeros de ~2 KB de memoria) y **Canales (Channels)**.
4. **La Base de Cloud Native:** El estándar absoluto en infraestructura moderna: **Docker, Kubernetes, Terraform, Prometheus, Traefik y CockroachDB** están programados en Go.

---

## 🇬🇧 2. Go vs Python (English Comparison Section)

During the technical presentation, an in-depth continuous English comparison (~4 minutes) contrasts Go and Python across three dimensions:

| Aspect | 🐍 Python | 🐹 Go |
| :--- | :--- | :--- |
| **Execution Model** | Interpreted bytecode (CPython VM) | Statically compiled to native machine code |
| **Typing System** | Dynamic typing (runtime verification) | Strict static typing with compile-time checks (`:=`) |
| **Concurrency** | Limited by the GIL; `asyncio` colored functions | Native **Goroutines** (~2 KB stack) and typed **Channels** |
| **Throughput & Speed**| High CPU overhead; requires C extensions (NumPy) | 10x – 40x faster raw execution; sub-millisecond GC |
| **Deployment / DevOps** | Virtualenvs (`venv`), complex dependency trees, heavy Docker images (>300 MB) | Single standalone binary; tiny Docker images (`FROM scratch` < 15 MB) |
| **Best Used For** | AI/ML, Data Science, Scripting, Prototyping | Cloud Native, Microservices, High-throughput APIs, Networking |

---

## 🛠️ 3. Pasos para Inicializar y Ejecutar el Proyecto

Este proyecto demuestra un servidor HTTP nativo ultra-eficiente con **Goroutines y `sync.WaitGroup`** en tan solo **36 líneas de código**, utilizando **exclusivamente la librería estándar de Go** (cero dependencias externas).

### Prerrequisitos
- Tener instalado [Go (1.20 o superior)](https://go.dev/dl/).

### Inicialización desde CERO
1. **Comprobar la versión instalada:**
   ```bash
   go version
   ```
2. **Inicializar el módulo Go:**
   ```bash
   go mod init pildora-go
   ```
3. **Ejecutar el servidor en directo:**
   ```bash
   go run main.go
   ```
4. **Probar la respuesta en el navegador o terminal:**
   ```bash
   curl http://localhost:8080
   ```
   *Respuesta en pantalla:*
   ```text
   3 tareas completadas en: 1.001s (Secuencial serian 3s)
   ```
   *(Las 3 tareas de 1 segundo se ejecutan concurrentemente en micro-hilos del runtime, reduciendo la latencia un 66%).*

---

## 💡 4. Buenas Prácticas en Go (Best Practices)

- **Manejo explícito de errores:** Tratar los errores siempre inmediatamente después de que ocurran (`if err != nil`), no ignorarlos.
- **Formateo automático:** Usar siempre `gofmt` antes de hacer commit. No perder tiempo discutiendo estilos.
- **Uso de `defer`:** Utilizar `defer resp.Body.Close()` o `defer file.Close()` para garantizar la liberación de recursos tan pronto como se adquieren.
- **CSP (Concurrencia):** *"No te comuniques compartiendo memoria; comparte memoria comunicándote."* Priorizar canales y paso de mensajes sobre bloqueos manuales tipo Mutex siempre que sea posible.
- **Evitar goroutine leaks:** Asegurar siempre que toda goroutine tenga una condición de finalización clara o un contexto con cancelación (`context.Context`).

---

## 📚 5. Recursos Adicionales

- [A Tour of Go (Tutorial interactivo oficial)](https://go.dev/tour/)
- [Effective Go (Guía de estilo y patrones idiomáticos)](https://go.dev/doc/effective_go)
- [Go by Example (Ejemplos prácticos y concisos de cada concepto)](https://gobyexample.com/)
- [The Go Blog - Concurrency is not Parallelism (Rob Pike)](https://go.dev/blog/waza-talk)
- [Documentación oficial de Go](https://pkg.go.dev/)