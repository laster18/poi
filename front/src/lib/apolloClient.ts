import { ApolloClient, HttpLink, InMemoryCache, split } from '@apollo/client'
import { WebSocketLink } from '@apollo/client/link/ws'
import { getMainDefinition } from '@apollo/client/utilities'

const getLink = () => {
  let terminatingLink
  if (!process.browser) {
    const httpLink = new HttpLink({
      uri: 'http://api:8080/query',
      credentials: 'include',
    })

    terminatingLink = httpLink
  } else {
    const httpLink = new HttpLink({
      uri: 'http://localhost:8080/query',
      credentials: 'include',
    })

    const wsLink = new WebSocketLink({
      uri: 'ws://localhost:8080/query',
      options: {
        reconnect: true,
      },
    })

    const splitLink = split(
      ({ query }) => {
        const definition = getMainDefinition(query)
        return (
          definition.kind === 'OperationDefinition' &&
          definition.operation === 'subscription'
        )
      },
      wsLink,
      httpLink,
    )

    terminatingLink = splitLink
  }

  return terminatingLink
}

const makeCache = () => {
  let cache: InMemoryCache
  if (process.browser) {
    cache = new InMemoryCache({
      typePolicies: {
        Query: {
          fields: {
            isLoggedIn: {
              read() {
                return true
              },
            },
          },
        },
        Room: {
          fields: {
            messages: {
              keyArgs: false,
              merge(existing, incoming) {
                if (!incoming) return existing
                if (!existing) return incoming

                const { nodes, ...rest } = incoming
                // We only need to merge the nodes array.
                // The rest of the fields (pagination) should always be overwritten by incoming
                const result = rest
                result.nodes = [...nodes, ...existing.nodes]
                return result
              },
            },
          },
        },
      },
    })
  } else {
    cache = new InMemoryCache()
  }

  return cache
}

export const apolloClient = new ApolloClient({
  link: getLink(),
  cache: makeCache(),
  credentials: 'inclued',
  connectToDevTools: true, // TODO: import environment
})
