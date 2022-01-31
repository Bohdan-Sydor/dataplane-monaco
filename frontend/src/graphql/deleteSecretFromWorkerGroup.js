import { gql, GraphQLClient } from 'graphql-request';
import { useGlobalAuthState } from '../Auth/UserAuth';

const graphlqlEndpoint = process.env.REACT_APP_GRAPHQL_ENDPOINT_PRIVATE;

const query = gql`
    mutation deleteSecretFromWorkerGroup($environmentName: String!, $WorkerGroup: String!, $Secret: String!) {
        deleteSecretFromWorkerGroup(environmentName: $environmentName, WorkerGroup: $WorkerGroup, Secret: $Secret)
    }
`;

export const useDeleteSecretFromWorkerGroup = () => {
    const authState = useGlobalAuthState();
    const jwt = authState.authToken.get();

    const headers = {
        Authorization: 'Bearer ' + jwt,
    };

    const client = new GraphQLClient(graphlqlEndpoint, {
        headers,
    });

    return async (input) => {
        try {
            const res = await client.request(query, input);
            return res?.deleteSecretFromWorkerGroup;
        } catch (error) {
            return JSON.parse(JSON.stringify(error, undefined, 2)).response;
        }
    };
};
