import {
  Account,
  getAccount,
  getPublication,
  Link,
  Publication,
} from '@app/client'
import {blockNodeToSlate} from '@app/client/v2/block-to-slate'
import {queryKeys} from '@app/hooks'
import {ClientPublication} from '@app/publication-machine'
import {QueryClient} from '@tanstack/react-query'
import {assign, createMachine} from 'xstate'

type CreateDiscussionMachineProps = {
  client: QueryClient
  link: Link
}

type DiscussionMachineContext = {
  link: Link
  source: ClientPublication | null
  publication: Publication | null
  errorMessage: string
  author: Account | null
}

type DiscussionMachineServices = {
  fetchSource: {
    data: Publication
  }
  fetchAuthor: {
    data: Account
  }
}

/**
 *
 * - fetch current account to check if the user is the author
 *
 */
export function createDiscussionMachine({
  client,
  link,
}: CreateDiscussionMachineProps) {
  return createMachine(
    {
      id: `discussion-machine-${link.source?.documentId}-${link.source?.version}`,
      predictableActionArguments: true,
      tsTypes: {} as import('./discussion-machine.typegen').Typegen0,
      schema: {
        context: {} as DiscussionMachineContext,
        services: {} as DiscussionMachineServices,
      },
      context: {
        link,
        source: null,
        publication: null,
        author: null,
        errorMessage: '',
      },
      initial: 'fetching',
      states: {
        fetching: {
          invoke: [
            {
              src: 'fetchSource',
              id: 'fetchSource',
              onDone: {
                actions: ['assignSource', 'assignPublication'],
                target: 'idle',
              },
            },
          ],
        },
        idle: {
          invoke: {
            src: 'fetchAuthor',
            id: 'fetchAuthor',
            onDone: {
              actions: ['assignAuthor'],
            },
            onError: {
              actions: ['assignError'],
            },
          },
        },
      },
    },
    {
      services: {
        fetchSource: (context) =>
          client.fetchQuery({
            queryKey: [
              queryKeys.GET_PUBLICATION,
              context.link.source?.documentId,
              context.link.source?.version,
            ],
            queryFn: () =>
              getPublication(
                context.link.source?.documentId,
                context.link.source?.version,
              ),
            staleTime: Infinity,
          }),
        fetchAuthor: async (context) => {
          let documentAuthor = context.source?.document.author || ''
          let userAccount = await client.fetchQuery({
            queryKey: [queryKeys.GET_ACCOUNT, ''],
            queryFn: () => getAccount(''),
          })
          if (documentAuthor == userAccount.id) {
            return userAccount
          } else {
            let authorAccount = await client.fetchQuery({
              queryKey: [queryKeys.GET_ACCOUNT, documentAuthor],
              queryFn: () => getAccount(documentAuthor),
            })
            return authorAccount
          }
        },
      },
      actions: {
        assignPublication: assign({
          publication: (_, event) => event.data,
        }),
        assignSource: assign({
          source: (_, event) => {
            let pub = event.data
            if (pub.document?.children.length) {
              let content = blockNodeToSlate(pub.document.children, 'group')
              let publication: ClientPublication = Object.assign(pub, {
                document: {
                  ...pub.document,
                  content: [content],
                },
              })
              return publication
            } else {
              throw Error('Something went wrong on Discussion fetchSource')
            }
          },
        }),
        assignAuthor: assign({
          author: (_, event) => event.data,
        }),
        assignError: assign({
          errorMessage: (_, event) => JSON.stringify(event),
        }),
      },
    },
  )
}
