import { useEffect, useState } from "react";
import { getUsers } from "./services/api";

export default function Home() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    getUsers().then(setUsers);
  }, []);
  console.log({ users });
  return (
    <div>
      <h1>Users</h1>
      <ul>dsd</ul>
    </div>
  );
}
