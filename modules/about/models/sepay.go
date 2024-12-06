package aboutmodels

// //https://docs.sepay.vn/lap-trinh-webhooks.html
// // HTTP server for the webhook
// http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
// 	var data WebhookData

// 	// Parse the incoming JSON
// 	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	// Determine transaction type and amounts
// 	if data.Code == "in" {
// 		data.AmountIn = data.Accumulated
// 	} else if data.Code == "out" {
// 		data.AmountOut = data.Accumulated
// 	}

// 	// Insert the data into the database
// 	_, err := db.NewInsert().Model(&data).Exec(ctx)
// 	if err != nil {
// 		http.Error(w, "Database insertion failed", http.StatusInternalServerError)
// 		log.Printf("Failed to insert record: %v", err)
// 		return
// 	}

// 	// Send success response
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"success": true,
// 	})
// })

// log.Println("Webhook server is running on port 8080...")
// if err := http.ListenAndServe(":8080", nil); err != nil {
// 	log.Fatalf("Failed to start server: %v", err)
// }
