
# Go API for Book and List Management

This project is a RESTful API built with Go, utilizing the Gin framework, GORM for ORM, and designed to manage books and lists. It includes CRUD operations for books and lists, allowing for the creation, retrieval, updating, and deletion of these entities.

## Features

- **Book Management**
  - Create, update, retrieve, and delete books
  - Search books by ID or name
  - Retrieve all books

- **List Management**
  - Create, update, delete lists
  - Add or remove books from lists
  - Retrieve all books in a list
  - Retrieve all lists for a user

## Prerequisites

- Go 1.16 or higher
- MySQL or any other supported database for GORM
- Gin Gonic framework
- GORM library

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/oliverperboni/GoTomekeeper.git
   cd GoApi
   ```

2. **Install Dependencies:**

   Ensure you have Go modules enabled and run:

   ```bash
   go mod tidy
   ```

3. **Set Up the Database:**

   Configure your database connection in the code. Update the database configuration in the project as needed.

4. **Run the Application:**

   ```bash
   go run main.go
   ```

   By default, the server will start on `localhost:8080`.

## API Endpoints

### Book Endpoints

- **Create Book**
  - `POST /book`
  - **Request Body:**

    ```json
    {
      "Name": "Book Name",
      "Author": "Author Name",
      "PublishedAt": "YYYY-MM-DD",
      "ISBN": "123-456-789"
    }
    ```

  - **Response:**

    ```json
    {
      "message": "your book was created :)"
    }
    ```

- **Get All Books**
  - `GET /book`
  - **Response:**

    ```json
    [
      {
        "Id": 1,
        "Name": "Book Name",
        "Author": "Author Name",
        "PublishedAt": "YYYY-MM-DD",
        "ISBN": "123-456-789"
      }
    ]
    ```

- **Get Book by ID**
  - `GET /book/id/:id`
  - **Response:**

    ```json
    {
      "Id": 1,
      "Name": "Book Name",
      "Author": "Author Name",
      "PublishedAt": "YYYY-MM-DD",
      "ISBN": "123-456-789"
    }
    ```

- **Get Book by Name**
  - `GET /book/name/:name`
  - **Response:**

    ```json
    {
      "Id": 1,
      "Name": "Book Name",
      "Author": "Author Name",
      "PublishedAt": "YYYY-MM-DD",
      "ISBN": "123-456-789"
    }
    ```

- **Update Book**
  - `PUT /book`
  - **Request Body:**

    ```json
    {
      "Id": 1,
      "Name": "Updated Book Name",
      "Author": "Updated Author Name",
      "PublishedAt": "YYYY-MM-DD",
      "ISBN": "123-456-789"
    }
    ```

  - **Response:**

    ```json
    {
      "message": "your book was updated :)"
    }
    ```

- **Delete Book**
  - `DELETE /book`
  - **Request Body:**

    ```json
    {
      "Id": 1
    }
    ```

  - **Response:**

    ```json
    {
      "message": "your book was deleted :)"
    }
    ```

### List Endpoints

- **Create List**
  - `POST /lists`
  - **Request Body:**

    ```json
    {
      "Name": "List Name",
      "UserID": 1
    }
    ```

  - **Response:**

    ```json
    {
      "message": "your list was created :)"
    }
    ```

- **Get All Lists for User**
  - `GET /list/user`
  - **Query Parameters:**
    - `userID`: User ID

  - **Response:**

    ```json
    [
      {
        "Id": 1,
        "Name": "List Name",
        "UserID": 1
      }
    ]
    ```

- **Get All Books in a List**
  - `GET /list/allBooks`
  - **Query Parameters:**
    - `userID`: User ID
    - `listID`: List ID

  - **Response:**

    ```json
    [
      {
        "Id": 1,
        "Name": "Book Name",
        "Author": "Author Name",
        "PublishedAt": "YYYY-MM-DD",
        "ISBN": "123-456-789"
      }
    ]
    ```

- **Get Book in List**
  - `GET /list/book`
  - **Query Parameters:**
    - `bookID`: Book ID
    - `listID`: List ID

  - **Response:**

    ```json
    {
      "Id": 1,
      "Name": "Book Name",
      "Author": "Author Name",
      "PublishedAt": "YYYY-MM-DD",
      "ISBN": "123-456-789"
    }
    ```

- **Add Book to List**
  - `POST /lists/AddBook`
  - **Query Parameters:**
    - `bookID`: Book ID
    - `listID`: List ID

  - **Response:**

    ```json
    {
      "msg": "book added :)"
    }
    ```

- **Remove Book from List**
  - `DELETE /list/book/`
  - **Query Parameters:**
    - `bookID`: Book ID
    - `listID`: List ID

  - **Response:**

    ```json
    {
      "msg": "book was deleted from the list :)"
    }
    ```

- **Update List**
  - `PUT /lists`
  - **Request Body:**

    ```json
    {
      "Id": 1,
      "Name": "Updated List Name",
      "UserID": 1
    }
    ```

  - **Response:**

    ```json
    {
      "message": "your list was updated :)"
    }
    ```

- **Delete List**
  - `DELETE /lists`
  - **Request Body:**

    ```json
    {
      "Id": 1
    }
    ```

  - **Response:**

    ```json
    {
      "msg": "list deleted :)"
    }
    ```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request. For major changes, open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


Happy coding!


