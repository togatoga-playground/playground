import { useEffect, useState } from "react";

function SayHello() {
  const [data, setData] = useState({ name: "" });
  useEffect(() => {
    fetch("api/hello")
      .then((res) => res.json())
      .then((profile) => setData(profile));
  }, []);

  return <div>Hello {data.name}</div>;
}

export default SayHello;
