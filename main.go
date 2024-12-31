package main

import (
	"fmt"
	"net/http"
)

// GoTasker: A Simple Task Management App
func main() {
    // Handle the root path "/" with the helloUser function
    http.HandleFunc("/", helloUser)
    
    // Handle the "/tasks" path with the showTasks function
    http.HandleFunc("/tasks", showTasks)

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

// helloUser handles the root path "/" and sends a welcome message
func helloUser(writer http.ResponseWriter, request *http.Request) {
    var greeting = "Hello, User! Welcome to GoTasker!"
    
    // Write the greeting to the response
    fmt.Fprintf(writer, greeting)
}

// showTasks handles the "/tasks" path and sends a list of tasks
func showTasks(writer http.ResponseWriter, request *http.Request) {
    // Simulate a task list
    tasks := []string{
        "Learn Go",
        "Write Code",
        "Test Application",
        "Deploy App",
    }

    // Send a heading for the tasks
    fmt.Fprintf(writer, "=== Your Tasks ===\n")

    // Iterate over the task list and print each task
    for i, task := range tasks {
        fmt.Fprintf(writer, "Task %d: %s\n", i+1, task)
    }
}


//showTasks function
func printTasks(tasks []string){
    fmt.Println("Here are your tasks:")
        // Iterate and print each task with index + 1
    for index, task := range tasks {
        fmt.Printf("Task %d: %s\n", index+1, task)
    }
}

//addTask function
func addTask(tasks []string, newTask string) []string {
    var updatedTasks = append(tasks, newTask)
    return updatedTasks
}



