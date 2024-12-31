package main

import "fmt"

// GoTasker: A Simple Task Management App
func main() {
    // Define tasks as variables
    var golang = "Learn Go"
    var coding = "Write Code"
    var testing = "Test Application"

    // Create a slice with tasks
    var tasks = []string{golang, coding, testing}

    // Print header for GoTasker app
    fmt.Println("=== GoTasker: Your Task Management App ===")
    fmt.Println("Here are your tasks:")

    // Iterate and print each task with index + 1
    for index, task := range tasks {
        fmt.Printf("Task %d: %s\n", index+1, task)
    }

    fmt.Println("=== End of Task List ===")
}

