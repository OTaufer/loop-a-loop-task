package main

import (
    "fmt"
)

// Item represents an inventory item.
type Item struct {
    ID          int
    Name        string
    Description string
    Price       float64
    Quantity    int
}

func main() {
    var inventory []Item // Initialize an empty inventory

    for {
        fmt.Println("Main Menu:")
        fmt.Println("1. Add an item")
        fmt.Println("2. Remove an item")
        fmt.Println("3. List all items")
        fmt.Println("4. Quit")

        var choice int
        fmt.Print("Enter your choice: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            // Add an item
            newItem := addItem()
            inventory = append(inventory, newItem)
            fmt.Println("Item added to inventory.")
        case 2:
            // Remove an item
            removeItem(&inventory)
        case 3:
            // List all items
            listItems(inventory)
        case 4:
            // Quit
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func addItem() Item {
    var newItem Item

    fmt.Print("Enter ID: ")
    fmt.Scanln(&newItem.ID)
    fmt.Print("Enter Name: ")
    fmt.Scanln(&newItem.Name)
    fmt.Print("Enter Description: ")
    fmt.Scanln(&newItem.Description)
    fmt.Print("Enter Price: ")
    fmt.Scanln(&newItem.Price)
    fmt.Print("Enter Quantity: ")
    fmt.Scanln(&newItem.Quantity)

    return newItem
}

func removeItem(inventory *[]Item) {
    var idToRemove int
    fmt.Print("Enter ID of the item to remove: ")
    fmt.Scanln(&idToRemove)

    // TASK 2 Create a deletion of the given item using a loop.

    fmt.Println("Item not found in inventory.")
}

func listItems(inventory []Item) {
    if len(inventory) == 0 {
        fmt.Println("Inventory is empty.")
        return
    }

    fmt.Println("Inventory:")
    // TASK 1 List and print inventory using a loop
}