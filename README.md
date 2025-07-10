# golang-nextjs-usernamagement

golang-nextjs-usernamagement

# From root folder

docker-compose up --build

# OR

# docker image

docker build -t user-service ./backend/user-service

sudo docker-compose up --build

## Structure

<pre>
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
</pre>
