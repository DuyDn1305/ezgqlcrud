import { gql, useQuery } from "@apollo/client";

const GET_CHARS = gql`
    query {
    todos {
      id
      text
      status
    }
  }
`;

function App() {
  const { loading, error, data } = useQuery(GET_CHARS);

  return (
    <div className="flex h-[300px] items-center justify-center">
      {loading && <p>Loading...</p>}
      {error && (
        <div>
          <p>Some thing went wrong....</p>
          <code>{error.toString()}</code>
        </div>
      )}
      {data && <code>{JSON.stringify(data)}</code>}
    </div>
  );
}

export default App;
