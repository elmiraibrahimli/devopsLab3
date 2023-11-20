package handlers

import (
    "database/sql"
    "net/http"
    "strconv"
    "devopslab3/db"

    "github.com/gin-gonic/gin"
)

type Item struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

func GetAllItems(c *gin.Context) {
    var items []Item
    rows, err := db.Db.Query("SELECT id, name, description FROM items")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        items = append(items, item)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
    id := c.Param("id")
    var item Item
    err := db.Db.QueryRow("SELECT id, name, description FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.Description)

    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {
    var newItem Item
    if err := c.BindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.Db.Exec("INSERT INTO items (name, description) VALUES ($1, $2)", newItem.Name, newItem.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newItem)
}

func UpdateItem(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
        return
    }

    var updateItem Item
    if err := c.BindJSON(&updateItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err = db.Db.Exec("UPDATE items SET name = $1, description = $2 WHERE id = $3", updateItem.Name, updateItem.Description, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
}

func DeleteItem(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
        return
    }

    _, err = db.Db.Exec("DELETE FROM items WHERE id = $1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}




// package handlers

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"devopslab3/db"

// 	"github.com/gorilla/mux"
// )

// type Item struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	// Add other fields as needed
// }

// func createItem(w http.ResponseWriter, r *http.Request) {
// 	var newItem Item
// 	_ = json.NewDecoder(r.Body).Decode(&newItem)

// 	// Insert into PostgreSQL database
// 	_, err := db.Exec("INSERT INTO items(name) VALUES($1)", newItem.Name)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newItem)
// }

// func getAllItems(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("SELECT id, name FROM items")
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var items []Item
// 	for rows.Next() {
// 		var item Item
// 		if err := rows.Scan(&item.ID, &item.Name); err != nil {
// 			log.Println(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		items = append(items, item)
// 	}

// 	json.NewEncoder(w).Encode(items)
// }

// func updateItem(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	var updatedItem Item
// 	_ = json.NewDecoder(r.Body).Decode(&updatedItem)

// 	// Update item in PostgreSQL database
// 	_, err = db.Exec("UPDATE items SET name = $1 WHERE id = $2", updatedItem.Name, id)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(updatedItem)
// }

// func deleteItem(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Delete item from PostgreSQL database
// 	_, err = db.Exec("DELETE FROM items WHERE id = $1", id)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
