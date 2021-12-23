import { gql, GraphQLClient } from "graphql-request";
import { useGlobalAuthState } from "../Auth/UserAuth";

const graphlqlEndpoint = process.env.REACT_APP_GRAPHQL_ENDPOINT_PRIVATE

const UpdatePreferences = gql`
    mutation updatePreferences($input: AddPreferencesInput!){
      updatePreferences(input: $input)
    }
`;

export const useUpdatePreferences = () => { 
    const authState = useGlobalAuthState();
    const jwt = authState.authToken.get();

    const headers = {
      Authorization: "Bearer " + jwt,
    };
  
    const client = new GraphQLClient(graphlqlEndpoint, {
      headers,
    });
  
    return async (input) => {
      try {
        const res = await client.request(UpdatePreferences, input);
        return res?.updatePreferences;
      } catch (error) {
        return JSON.parse(JSON.stringify(error, undefined, 2)).response
      }
    };
  };