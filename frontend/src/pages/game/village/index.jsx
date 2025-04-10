import { useParams, Navigate } from "react-router-dom";

export default function VillagePage() {
  const { id } = useParams();

  if (!id) {
    return <Navigate to="/" />;
  }

  return (
    <div>
      <h1>Village {id}</h1>
    </div>
  );
}
