import {mainService as defaultMainService} from '@app/app-providers'
import {createPublicationMachine} from '@app/publication-machine'
import {error} from '@app/utils/logger'
import {Box} from '@components/box'
import {DiscussionItem} from '@components/discussion-item'
import {useActor} from '@xstate/react'
import {ActorRefFrom} from 'xstate'

export type DiscussionProps = {
  service: ActorRefFrom<ReturnType<typeof createPublicationMachine>>
  mainService?: typeof defaultMainService
}

export function Discussion({
  service,
  mainService = defaultMainService,
}: DiscussionProps) {
  const [state] = useActor(service)
  const [mainState] = useActor(mainService)

  if (state.matches('discussion.fetching')) {
    return <span>loading discussion...</span>
  }

  if (state.matches('discussion.errored')) {
    error('Discussion Error')
    return <span>Discussion ERROR</span>
  }

  if (state.matches('discussion.ready.visible')) {
    return (
      <Box
        css={{
          display: 'flex',
          flexDirection: 'column',
          gap: '$4',
          paddingHorizontal: '$4',
        }}
      >
        {state.context.links.map((link) => {
          let {source} = link
          let key = `link-${source?.documentId}-${source?.version}-${source?.blockId}`
          return <DiscussionItem key={key} link={link} />
        })}
      </Box>
    )
  }

  return null
}