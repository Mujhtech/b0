# Run the backend
run-backend:
	cd backend && go build -o b0 ./cmd && ./b0 serve

# Run the frontend
run-frontend:
	cd frontend && pnpm run dev