import {styled} from '@mintter/ui/stitches.config'

export const EmbedUI = styled('q', {
  borderRadius: '$1',
  transition: 'all ease-in-out 0.1s',
  backgroundColor: '$background-alt',
  '&:hover': {
    backgroundColor: '$secondary-softer',
    cursor: 'pointer',
    // color: '$text-contrast',
  },
  '&::before, &::after': {
    fontWeight: '$bold',
    fontSize: '1.4em',
    color: '$text-alt',
  },
  '&::before': {
    content: '[',
  },
  '&::after': {
    content: ']',
  },
})