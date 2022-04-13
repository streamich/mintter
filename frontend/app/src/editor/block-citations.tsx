import {Account, getAccount, Link, Publication, SidepanelItem} from '@app/client'
import {useBlockCitations} from '@app/editor/citations'
import {css} from '@app/stitches.config'
import {getBlock} from '@app/utils/get-block'
import {Avatar} from '@components/avatar'
import {Box} from '@components/box'
import {useSidepanel} from '@components/sidepanel'
import {Text} from '@components/text'
import {FlowContent} from '@mintter/mttast'
import {useMachine} from '@xstate/react'
import {assign, createMachine} from 'xstate'

export type BlockCitationsProps = {
  blockId: string
}

export function BlockCitations({blockId}: BlockCitationsProps) {
  const blockCitations = useBlockCitations(blockId)
  const [state, send] = useMachine(blockCitationsMachine)
  return blockCitations.length ? (
    <Box
      contentEditable={false}
      css={{
        padding: '$4',
        width: '$full',
        maxWidth: 240,
        marginTop: '-$5',
        '@bp2': {
          transform: 'translateX(100%)',
          position: 'absolute',
          right: 0,
          top: 4,
          overflow: 'hidden',
        },
      }}
    >
      <button className={citationsButton()} onClick={() => send('TOGGLE')}>
        <Box className={`citation-index ${citationCount()}`}>{blockCitations.length}</Box>
        <Text color="primary" fontWeight="bold" size="1">
          {state.matches('collapsed') ? 'Show' : 'Hide'} citations
        </Text>
      </button>
      {state.matches('expanded') ? (
        <Box
          as="ul"
          css={{
            margin: 0,
            padding: 0,
            display: 'flex',
            flexDirection: 'column',
            gap: '$3',
          }}
        >
          {blockCitations.map((citation) => (
            <BlockCitationItem key={citation.source?.blockId} citation={citation} />
          ))}
        </Box>
      ) : null}
    </Box>
  ) : null
}

var citationsButton = css({
  all: 'unset',
  padding: '$2',
  borderRadius: '$2',
  display: 'flex',
  alignItems: 'center',
  gap: '$3',
  '&:hover': {
    background: '$block-hover',
    '& .citation-index': {},
  },
})

var citationCount = css({
  width: 24,
  height: 24,
  borderRadius: '$2',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  background: '$block-hover',
  // background: 'green',
  fontFamily: '$default',
  fontSize: '$1',
  color: '$text-muted',
  textAlign: 'center',
  '& span': {
    background: 'red',
    display: 'block',
  },
})

var blockCitationsMachine = createMachine({
  initial: 'collapsed',
  context: {},
  states: {
    collapsed: {
      on: {
        EXPAND: 'expanded',
        TOGGLE: 'expanded',
      },
    },
    expanded: {
      on: {
        COLLAPSE: 'collapsed',
        TOGGLE: 'collapsed',
      },
    },
  },
})

type BlockCitationItemProps = {
  citation: Link
}

function BlockCitationItem({citation}: BlockCitationItemProps) {
  const sidepanelService = useSidepanel()

  const [state, send] = useMachine(
    //@ts-ignore
    () => blockCitationMachine,
    {
      services: {
        fetchCitation: () => (sendBack) => {
          ;(async () => {
            try {
              let data = await getBlock(citation.source)
              let author = await getAccount(data?.publication.document?.author || '')
              console.log('async result: ', {data, author})

              if (data) {
                sendBack({
                  type: 'CITATION.FETCH.SUCCESS',
                  publication: data.publication,
                  block: data.block,
                  author,
                })
              } else {
                sendBack({type: 'CITATION.FETCH.ERROR'})
              }
            } catch {
              sendBack({type: 'CITATION.FETCH.ERROR'})
            }
          })()
        },
      },
      actions: {
        assignAuthor: assign({
          author: (_, event) => event.author,
        }),
        assignBlock: assign({
          block: (_, event) => event.block,
        }),
        assignPublication: assign({
          publication: (_, event) => event.publication,
        }),
        assignError: assign({
          errorMessage: (context) => 'Error fetching',
        }),
        clearError: assign({
          errorMessage: (context) => '',
        }),
        openInSidepanel: (_, event) => {
          sidepanelService.send({
            type: 'SIDEPANEL.ADD',
            item: event.item,
          })
          sidepanelService.send('SIDEPANEL.OPEN')
        },
      },
    },
  )

  let title = state.context?.publication?.document?.title || 'Untitled Document'
  let authorAlias = state.context?.author?.profile?.alias || 'anonymous'

  console.log({state, title, authorAlias})

  return (
    <Box
      css={{
        display: 'flex',
        alignItems: 'center',
        borderRadius: '$3',
        '&:hover': {
          background: '$block-hover',
          cursor: 'pointer',
        },
      }}
      onClick={() => {
        send({
          type: 'OPEN.IN.SIDEPANEL',
          item: {
            type: 'block',
            url: `mtt://${citation.source?.documentId}/${citation.source?.version}/${citation.source?.blockId}`,
          },
        })
      }}
    >
      <Box
        css={{
          padding: '$3',
          // background: '$block-hover',
          borderRadius: '$2',
          display: 'flex',
          flexDirection: 'column',
          gap: '$2',
        }}
      >
        <Text size="1" fontWeight="bold" color="default">
          {title}
        </Text>
        <Box
          css={{
            display: 'flex',
            gap: '$2',
            alignItems: 'center',
          }}
        >
          <Avatar size="1" />
          <Text size="1" color="muted">
            {authorAlias}
          </Text>
        </Box>
      </Box>
    </Box>
  )
}

type BlockCitationEvent =
  | {type: 'CITATION.FETCH.SUCCESS'; publication: Publication; block: FlowContent; author: Account}
  | {type: 'CITATION.FETCH.ERROR'}
  | {type: 'RETRY'}
  | {type: 'OPEN.IN.SIDEPANEL'; item: SidepanelItem}

const blockCitationMachine = createMachine({
  tsTypes: {} as import('./block-citations.typegen').Typegen0,
  schema: {
    context: {} as {
      publication: Publication | null
      block: FlowContent | null
      author: Account | null
      errorMessage: string
    },
    events: {} as BlockCitationEvent,
  },
  initial: 'loading',
  context: {
    publication: null as Publication | null,
    block: null as FlowContent | null,
    author: null as Account | null,
    errorMessage: '',
  },
  states: {
    loading: {
      invoke: {
        id: 'fetchCitation',
        src: 'fetchCitation',
      },
      on: {
        'CITATION.FETCH.SUCCESS': {
          target: 'ready',
          actions: ['assignPublication', 'assignBlock', 'assignAuthor'],
        },
        'CITATION.FETCH.ERROR': {
          target: 'errored',
          actions: ['assignError'],
        },
      },
    },
    ready: {
      on: {
        'OPEN.IN.SIDEPANEL': {
          actions: ['openInSidepanel'],
        },
      },
    },
    errored: {
      on: {
        RETRY: {
          target: 'loading',
          actions: ['clearError'],
        },
      },
    },
  },
})