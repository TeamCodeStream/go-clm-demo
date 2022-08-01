package main

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"io"
	"net/http"
	"time"
)

func mysql(w http.ResponseWriter, r *http.Request) {
	txn := newrelic.FromContext(r.Context())
	s := newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		// Product, Collection, and Operation are the most important
		// fields to populate because they are used in the breakdown
		// metrics.
		Product:    newrelic.DatastoreMySQL,
		Collection: "users",
		Operation:  "INSERT",

		ParameterizedQuery: "INSERT INTO users (name, age) VALUES ($1, $2)",
		QueryParameters: map[string]interface{}{
			"name": "Dracula",
			"age":  439,
		},
		Host:         "mysql-server-1",
		PortPathOrID: "3306",
		DatabaseName: "my_database",
	}
	defer s.End()

	time.Sleep(20 * time.Millisecond)
	io.WriteString(w, `performing fake query "INSERT * from users"`)
}

func roundtripper(w http.ResponseWriter, r *http.Request) {
	// NewRoundTripper allows you to instrument external calls without
	// calling StartExternalSegment by modifying the http.Client's Transport
	// field.  If the Transaction parameter is nil, the RoundTripper
	// returned will look for a Transaction in the request's context (using
	// FromContext). This is recommended because it allows you to reuse the
	// same client for multiple transactions.
	client := &http.Client{}
	client.Transport = newrelic.NewRoundTripper(client.Transport)

	request, _ := http.NewRequest("GET", "http://example.com", nil)
	// Since the transaction is already added to the inbound request's
	// context by WrapHandleFunc, we just need to copy the context from the
	// inbound request to the external request.
	request = request.WithContext(r.Context())
	// Alternatively, if you don't want to copy entire context, and instead
	// wanted just to add the transaction to the external request's context,
	// you could do that like this:
	//
	//	txn := newrelic.FromContext(r.Context())
	//	request = newrelic.RequestWithTransactionContext(request, txn)

	resp, err := client.Do(request)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
