This project is a basic web server written in Go that serves static files and handles basic HTTP requests, including form submissions.

Features:

Serves static HTML files (e.g., index.html).

Handles a GET request at the /huh endpoint.

Processes form submissions via POST requests at the /form endpoint.

Simple logging to track server startup and errors.

File Structure

.
|-- main.go           # Go code for server logic
|-- go.mod            # Go module file
|-- static
|   |-- index.html    # Static landing page
|   |-- form.html     # Simple form for user input

How It Works

Static Files: The server serves files from the static directory. The default file served at / is index.html.

Endpoints:

/huh: Responds with wassup bro for GET requests. Returns 404 for other methods.

/form: Handles POST requests and parses form data (fields: name and goon).

Setup and Run

Requirements

Go (v1.16 or higher recommended)

Steps to Run

Clone or download the project.

Navigate to the project directory.

Run the following command to start the server:

go run main.go

Open your browser and visit http://localhost:8080 to see the static site.

Navigate to http://localhost:8080/form to submit the form.

Visit http://localhost:8080/huh for a custom endpoint response.
