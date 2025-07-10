export async function getUsers() {
  const res = await fetch("http://localhost:8080/users");
  return res.json();
}
