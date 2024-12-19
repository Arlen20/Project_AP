# Kazakh History

An engaging and user-friendly web application with full frontend functionality and backend support for CRUD operations, premium account features, login, and registration. This project is built to provide a dynamic experience for admins and users, with seamless integration of a MongoDB database.

  ![Screenshot 2024-12-15 210446](https://github.com/user-attachments/assets/11a5dc43-2003-49f4-8eb5-a8b6d676da07)


## About The Project

The purpose of this project is to explain everything related to Kazakh History and the Kazakh Kaganat using all available methods. Users can view, filter, and engage with content, while admins have full control to manage the website's content through CRUD operations. Premium users also have the ability to edit articles, pending admin approval.

**Key Features:**

- **Admin Panel**: Admins can log in and manage website content using full CRUD operations.
- **Premium User Feature**: Users can request premium status to edit articles. Requests are handled and approved by admins.
- **Login & Registration**: Users can register and log in with credentials stored in a MongoDB database.
- **Content Filtering**: Users can filter displayed content, such as images, by year.

## Built With

This project utilizes modern web development tools and frameworks:

- **Frontend**: HTML, CSS, JavaScript, and a modern JavaScript framework (like React,Bootstrap,Postman,Jquary)
- **Backend**: Go (Golang) for server-side development
- **Database**: MongoDB for data storage and retrieval

## Getting Started

Follow these steps to set up the project locally.

### Prerequisites

Make sure you have the following installed:

- **Go** (latest version)
- **MongoDB** (installed and running locally or via a cloud service like MongoDB Atlas)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Arlen20/Project_AP.git
   ```
2. Navigate to the project directory:
   ```sh
   cd project_name
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Set up environment variables in a `.env` file:
   ```env
   MONGO_URI=your_mongodb_connection_string
   JWT_SECRET=your_jwt_secret
   PORT=3000
   ```
5. Start the development server:
   ```sh
   go run main.go
   ```

## Usage

### Admin Features

- **Login**: Admins can log in using their credentials.
- **CRUD Operations**: Admins can create, read, update, and delete website content.

### User Features

- **Registration & Login**: Users can register and log in using the form provided.
- **Premium Account Request**: Users can request a premium account, which allows them to edit articles upon admin approval.
- **Content Filtering**: Users can filter images or content by year.

## Contact

**Team Members:**

- **Baibossyn Arlen**
- **Nurbol Nurlybai**
- **Muslim Toleumukhanbet**

## Acknowledgments

- [MongoDB](https://www.mongodb.com/)
- [Go (Golang)](https://go.dev/)

This README serves as a guide to help you understand and set up the project. If you have any questions, feel free to reach out.
