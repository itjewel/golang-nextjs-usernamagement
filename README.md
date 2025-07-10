# golang-nextjs-usernamagement

golang-nextjs-usernamagement

# From root folder

docker-compose up --build

# OR

sudo docker-compose up --build

## Structure

my-app/
├── backend/ # Golang microservices
│ ├── user-service/
│ │ ├── cmd/ # Main entry
│ │ ├── config/
│ │ ├── internal/
│ │ │ ├── handler/
│ │ │ ├── service/
│ │ │ ├── repository/
│ │ │ ├── model/
│ │ │ └── middleware/
│ │ ├── pkg/
│ │ └── main.go
│ └── auth-service/
│ └── ...
│
├── frontend/ # Next.js app
│ ├── pages/
│ ├── components/
│ ├── services/
│ └── ...
│
├── docker-compose.yml
└── README.md
