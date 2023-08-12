# Shortener API in Golang

This is a simple API in Golang that accepts a URL in the query parameter and returns a shortened link that will redirect the user to the original URL.

![Diagram](https://user-images.githubusercontent.com/60992933/260232763-0ca250b0-da77-4d77-adb8-5b8acb7d1212.png)

## Prerequisites

To run the project, you need to have the following tools installed on your system:

- `make`: The build automation tool.
- `Docker`: To run the container with the Postgres database.

## Getting Started

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/your-username/shortener-api.git
   cd shortener-api
   ```

2. Install Go in your machine, and follow the instruction in the following url:

   ```bash
   https://go.dev
   ```

3. Install Make in your machine, and follow the instruction in the following url:

   ```bash
   https://www.gnu.org/software/make/
   ```

4. Initialize the Postgres database container:

   ```bash
   make postgresinit
   ```

5. Add environment variables below to your environment:

   ```bash
   # Your base URL
   BASE_URL=

   # Your url connection string
   DATABASE_URL=
   ```

6. Initialize the Postgres database container:

   ```bash
   make createdb
   ```

7. Run the database migrations:

   ```bash
   make migrateup
   ```

8. Finally, run the API locally:
   ```bash
   make run
   ```

## Making Requestes to the API

You can use curl or any other HTTP client to interact with the API.

### Shorten a URL

To shorten a URL, make a POST request with the url query parameter. The API will return a shortened link.

```bash
curl -X POST "http://localhost:8000/?url=https://www.example.com"
```

Sample Response:

```json
{
  "url": "http://localhost:8080/EawUrx"
}
```

### Redirect to the Original URL

To be redirected to the original URL, simply open the shortened URL in your browser

```bash
http://localhost:8000/EawUrx
```

This will redirect you to https://www.example.com.
