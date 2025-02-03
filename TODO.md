
URL Shortener Project

# Setup:
	Initialize a new Go project.
	Set up a basic web server using a framework like net/http or gin-gonic/gin.

# Database:
	Use a database like SQLite or PostgreSQL to store the original URLs and their shortened versions.
	Create a table with fields for the original URL, the shortened URL, and any metadata (like creation date).

# Shortening Logic:
	Implement a function to generate a unique short code for each URL. You can use a hash function or a base conversion method.
	Ensure that the generated short codes are unique and handle collisions if they occur.

# API Endpoints:
	[done - golang] Create an endpoint to accept a URL and return the shortened version.
	[done - golang] Create an endpoint to redirect a short URL to the original URL.
	OpenAPI

# User Interface:
	Build a simple web interface where users can input a URL and receive the shortened version.
	Display a list of previously shortened URLs (optional).

# Testing:
	Write unit tests for your shortening logic and API endpoints.
	Test the application thoroughly to ensure it handles edge cases and errors gracefully.


# Bugs golang
	- validate url

