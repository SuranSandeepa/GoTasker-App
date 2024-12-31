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
    printTasks(tasks)
    println()
    tasks = addTask(tasks, "Deploy Application")
    println("=== After Adding New Task ===")
    printTasks(tasks)
    fmt.Println("=== End of Task List ===")
}

func printTasks(tasks []string){
    fmt.Println("Here are your tasks:")
        // Iterate and print each task with index + 1
    for index, task := range tasks {
        fmt.Printf("Task %d: %s\n", index+1, task)
    }
}

func addTask(tasks []string, newTask string) []string {
    var updatedTasks = append(tasks, newTask)
    return updatedTasks
}

