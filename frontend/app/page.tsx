import { useEffect, useState } from "react";
import { getUsers } from "./services/api";

export default function Home() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    getUsers().then(setUsers);
  }, []);

  return (
    <div>
      <h1>Users</h1>
      <ul>
        {users.map((u: unknown) => (
          <li key={u.id}>{u.name}</li>
        ))}
      </ul>
    </div>
  );
}
