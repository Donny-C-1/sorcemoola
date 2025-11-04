<style> 
    * { font-family: "Comic Sans MS"; }
    .center {
        text-align: center;
    }
</style>

<div class="center">

# SorceMoola

SorceMoola is a crowdfunding web application that allows users to create, fund, and manage fundraising campaigns.

[View Demo](sorcemoola.vercel.com) |
[Report bug](<[google.com](https://github.com/Donny-C-1/sorcemoola?tab=contributing-ov-file#reporting-bugs)>) |
[Request Feature](<[google.com](https://github.com/Donny-C-1/sorcemoola?tab=contributing-ov-file#suggesting-enhancements)>)

![Static Badge](https://img.shields.io/badge/Project_Name-SorceMoola-%233CB371?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/status-development-dodgerblue?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/version-1.0-%233CB371?style=for-the-badge)

</div>

## Tech Stack

-   **Backend:** GO
-   **Frontend:** Svelte
-   **Database:** PostgreSQL (hosted on Neon DB)

## Features

-   User authentication and authorization
-   Campaign creation and management
-   Secure payment processing
-   Campaign discovery and filtering
-   User profile management
-   Real-time funding progress tracking

## Prerequisites

-   Go (1.25+)
-   Node.js (20+)
-   npm or yarn

## Getting Started

To check out the application, you can view the live demo right here or follow the instructions below to setup the application on your local device.

### Installation

1. Clone the repository:

```sh
git clone https://github.com/Donny-C-1/sorcemoola.git
```

2. Backend Setup

```sh
# Navigate to the server folder
cd sorcemoola/server

# Install dependencies
go mod tidy
```

- Configure environment variables in a `.env` file
```env
DATABASE_URL=<your_postgres_database_url>
PORT=<port>
```


- Run the backend
```go
go run main.go
```

3. Frontend Setup
```sh
# Navigate to frontend location
cd ../client

# Install dependencies
npm install

# Start development server
npm run dev
```

## License

This project is licensed under the Apache License v2. See the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome. Please review the [Contributing.md](CONTRIBUTING.md) file before attempting to contribute.

## 📱 Contact

Made with 💛💜 by Donny C. For any questions or suggestions please contact [Donny C](mailto:chikwemdonald@gmail.com).
