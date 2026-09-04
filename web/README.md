# Todo App — Frontend

Plain HTML/CSS/JS frontend for your Go + MongoDB todo API.

## Pages

- `login.html` — log in
- `register.html` — create an account (auto-logs you in after)
- `dashboard.html` — list, add, edit, complete, delete tasks
- `index.html` — redirects to dashboard or login depending on whether you're logged in

## Setup

1. Make sure your Go backend is running on `http://localhost:8080` (`go run main.go`)
2. Open `index.html` directly in your browser (double-click it), or serve the folder
   with any static server, e.g.:
   ```bash
   npx serve .
   ```
3. Register a new account, then you'll land on the dashboard.

## Important: backend CORS

Your Go server needs to allow requests from the browser (CORS), since the
frontend and API run on different origins. See the CORS middleware snippet
given alongside this — add it to `main.go` and restart your server before
testing the frontend.

## Config

If your backend runs on a different port/host, change the `API_BASE` constant
at the top of `api.js`.

## How auth works here

- After login/register, the JWT token is saved in the browser's `localStorage`.
- Every API call in `api.js` automatically attaches it as `Authorization: Bearer <token>`.
- `dashboard.html` calls `requireAuth()` on load — if there's no token, it redirects
  straight to `login.html`.
- Logging out just clears the token from `localStorage`.
