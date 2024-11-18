import { isServer, QueryClient } from '@tanstack/react-query'

let browserQueryClient: QueryClient | undefined

function makeQueryClient() {
  return new QueryClient()
}

export function getQueryClient() {
  if (isServer) {
    return makeQueryClient()
  } else {
    if (!browserQueryClient) {
      browserQueryClient = makeQueryClient()
    }

    return browserQueryClient
  }
}
