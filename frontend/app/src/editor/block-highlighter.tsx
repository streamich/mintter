import {useHighlightRef} from '@app/mouse-context'
import {Box} from '@components/box'
import {ReactNode} from 'react'

export function BlockHighLighter({children}: {children: ReactNode}) {
  let id = useHighlightRef()

  return (
    <Box
      css={
        id
          ? {
              [`& [data-highlight="${id}"]`]: {
                zIndex: 1,
              },
              [`& [data-highlight="${id}"]:before`]: {
                userSelect: 'none',
                content: '',
                position: 'absolute',
                top: '-0.5rem',
                right: 0,
                left: 0,
                width: '200vw',
                height: 'calc(100% + 1rem)',
                transform: 'translateX(-50%)',
                background: '$primary-component-bg-hover',
                zIndex: -1,
                pointerEvents: 'none',
              },
              [`& a[data-highlight="${id}"], & q[data-highlight="${id}"]`]: {
                background: '$primary-component-bg-hover',
              },
            }
          : {}
      }
    >
      {children}
    </Box>
  )
}