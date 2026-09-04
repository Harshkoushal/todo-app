// Change this if your backend runs on a different port
const API_BASE = "http://localhost:8080/api";

function getToken() {
  return localStorage.getItem("token");
}

function setToken(token) {
  localStorage.setItem("token", token);
}

function clearToken() {
  localStorage.removeItem("token");
  localStorage.removeItem("userEmail");
}

function requireAuth() {
  if (!getToken()) {
    window.location.href = "login.html";
  }
}

async function apiRequest(path, method = "GET", body = null) {
  const headers = { "Content-Type": "application/json" };
  const token = getToken();
  if (token) headers["Authorization"] = "Bearer " + token;

  const res = await fetch(API_BASE + path, {
    method,
    headers,
    body: body ? JSON.stringify(body) : null,
  });

  let data = null;
  try {
    data = await res.json();
  } catch (e) {
    data = null;
  }

  if (!res.ok) {
    const message = (data && (data.error || data.message)) || "Request failed";
    throw new Error(message);
  }

  return data;
}
