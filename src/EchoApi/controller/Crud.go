package controller

// "database/sql"
// "echoapi/model"
// "encoding/json"
// "log"
// "net/http"
// "strconv"

// select all
// func GetUsers(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		rows, err := db.Query("SELECT username,password,email,role FROM omni.user")
// 		if err != nil {
// 			// log.Fatal(err)
// 			log.Println(err)
// 		}
// 		defer rows.Close()

// 		users := []model.User{}
// 		for rows.Next() {
// 			var u model.User
// 			if err := rows.Scan(&u.Username, &u.Password, &u.Email, &u.Role); err != nil {
// 				// log.Fatal(err)
// 				log.Println(err)
// 			}
// 			users = append(users, u)
// 		}
// 		if err := rows.Err(); err != nil {
// 			// log.Fatal(err)
// 			log.Println(err)
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(users)
// 	}
// }

// func GetUserByID(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		idStr := vars["id"]
// 		var dr model.DataResponse
// 		// Check if id is numeric
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			dr.Error = true
// 			dr.Message = "ID Tidak Sesuai !"

// 			DataRespon := []model.DataResponse{}
// 			DataRespon = append(DataRespon, dr)

// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusBadRequest)
// 			log.Println("Invalid id:", idStr)
// 			json.NewEncoder(w).Encode(DataRespon)
// 			return
// 		}

// 		var u model.User
// 		err = db.QueryRow("Select id,username,email From omni.user where id = $1", id).Scan(&u.Id, &u.Username, &u.Email)

// 		//jika tidak terdapat data
// 		if err == sql.ErrNoRows {
// 			dr.Error = true
// 			dr.Message = "No Data for id : " + vars["id"]

// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusNotFound)
// 			json.NewEncoder(w).Encode(dr)

// 			return
// 		} else if err != nil { //jika terdapat error pada db
// 			log.Println("Error Querying database")
// 			dr.Error = true
// 			dr.Message = "Database Internal Error"

// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusInternalServerError)
// 			json.NewEncoder(w).Encode(dr)

// 			return
// 		} else {
// 			dr.Error = false
// 			dr.Message = "OK"
// 			dr.Data = u

// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)
// 			json.NewEncoder(w).Encode(dr)
// 			return
// 		}
// 	}
// }
