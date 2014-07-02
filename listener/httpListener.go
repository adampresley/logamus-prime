package listener

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/adampresley/logamus-prime/message"
	"github.com/gorilla/mux"
)

func StartHttpListener(address string, messageChannels []chan message.Message) error {
	router := mux.NewRouter()

	/*
	 * Local method to pushlish a message to all passed in channels
	 */
	publish := func(newMessage message.Message) {
		for _, c := range messageChannels {
			c <- newMessage
		}
	}

	/*
	 * Route handlers
	 */
	router.HandleFunc("/error", func(writer http.ResponseWriter, req *http.Request) {
		log.Println("POST /error")
		newMessage, err, status := parseMessageJson(req, message.ERROR_MESSAGE)

		if err != nil {
			http.Error(writer, err.Error(), status)
			return
		}

		publish(*newMessage)
		fmt.Fprintf(writer, "New error queued for writing")
	}).Methods("POST").Headers("Content-Type", "application/json")

	router.HandleFunc("/info", func(writer http.ResponseWriter, req *http.Request) {
		log.Println("POST /info")
		newMessage, err, status := parseMessageJson(req, message.INFO_MESSAGE)

		if err != nil {
			http.Error(writer, err.Error(), status)
			return
		}

		publish(*newMessage)
		fmt.Fprintf(writer, "New info message queued for writing")
	}).Methods("POST").Headers("Content-Type", "application/json")

	router.HandleFunc("/warning", func(writer http.ResponseWriter, req *http.Request) {
		log.Println("POST /warning")
		newMessage, err, status := parseMessageJson(req, message.WARNING_MESSAGE)

		if err != nil {
			http.Error(writer, err.Error(), status)
			return
		}

		publish(*newMessage)
		fmt.Fprintf(writer, "New warning message queued for writing")
	}).Methods("POST").Headers("Content-Type", "application/json")

	/*
	 * Fire up the HTTP listener!!
	 */
	log.Println("Starting HTTP Listener @ " + address)
	return http.ListenAndServe(address, router)
}

/*
Parses a form with a message This method
expects there to be a set of form variables:

* Date - In the format of '2014-01-01'
* Time - In the format of '10:14:00'
* Message - At least 5 characters
* StackTrace - JSON string which is an array of structs. Each
  struct contains fileName, lineNumber, and functionName. Optional
* Tags - An array of tags to categorize the message. Optional
*/
func parseMessageJson(req *http.Request, messageType message.MessageType) (*message.Message, error, int) {
	var err error

	body, _ := ioutil.ReadAll(req.Body)

	/*
	 * Create our message
	 */
	newMessage := &message.Message{Type: messageType}

	err = json.Unmarshal(body, newMessage)
	if err != nil {
		return nil, errors.New("Unable to parse body: " + err.Error()), 400
	}

	/*
	 * Validate we have our required data
	 */
	if newMessage.Date == "" {
		return nil, errors.New("Date is expected"), 400
	}

	if newMessage.Time == "" {
		return nil, errors.New("Time is expected"), 400
	}

	if len(strings.TrimSpace(newMessage.Message)) < 5 {
		return nil, errors.New("Please provide a valid message of at least 5 characters"), 400
	}

	return newMessage, nil, 200
}
